package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"context"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/lazada/grpc-ui/reflection"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

type Req struct {
	Type string           `json:"type"`
	Data *json.RawMessage `json:"data"`
}

type InvokeReq struct {
	IsStream    bool   `json:"is_stream"`
	Addr        string `json:"addr"`
	PackageName string `json:"package_name"`
	ServiceName string `json:"service_name"`
	MethodName  string `json:"method_name"`
	GRPCMethod  string `json:"grpc_method"`
	GRPCArgs    string `json:"grpc_args"`
}

type InvokeResp struct {
	Type        string      `json:"type"`
	PackageName string      `json:"package_name"`
	ServiceName string      `json:"service_name"`
	MethodName  string      `json:"method_name"`
	Error       string      `json:"error"`
	Data        interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		addr := r.FormValue("addr")

		_, _, err := net.SplitHostPort(addr)
		if err != nil {
			log.Printf("Invalid `addr` param")
			http.Error(w, "Invalid `addr` param", http.StatusBadRequest)
			return
		}
		services, err := reflection.GetInfo(r.Context(), addr)

		if err != nil {
			log.Printf("Can't get grpc info: %v", err)
			http.Error(w, fmt.Sprintf("Can't get grpc info: %v", err), http.StatusInternalServerError)
			return
		}

		resp := struct {
			Services []reflection.Service `json:"services"`
		}{
			Services: services,
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Can't encode json: %v", err)
			return
		}
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Can't upgrade connection: %v", err)
			return
		}
		defer conn.Close()

		for {
			msg := Req{}
			if err := conn.ReadJSON(&msg); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("ReadJSON error: %v", err)
				}
				return
			}
			switch msg.Type {
			case "invoke":
				if err := handleInvoke(r.Context(), conn, msg.Data); err != nil {
					log.Printf("Can't handle invoke method %v", err)
					sendError(conn, err)
				}
			default:
				log.Printf("Invalid message type: %v", msg.Type)
				sendError(conn, errors.New("Invalid message type"))
				return
			}
		}
	})

	http.ListenAndServe(":3000", nil)
}

func sendError(conn *websocket.Conn, err error) error {
	resp := struct {
		Type string `json:"type"`
		Data string `json:"data"`
	}{
		Type: "ERROR",
		Data: err.Error(),
	}
	if err := conn.WriteJSON(&resp); err != nil {
		return err
	}

	return nil
}

func handleInvoke(ctx context.Context, conn *websocket.Conn, bytes *json.RawMessage) error {
	log.Printf("handleInvokeMethod")

	if bytes == nil {
		return errors.New("Empty `data` field")
	}

	req := InvokeReq{}
	if err := json.Unmarshal(*bytes, &req); err != nil {
		return err
	}

	if req.IsStream {
		go func() {
			err := reflection.InvokeStream(ctx, req.Addr, req.GRPCMethod, []byte(req.GRPCArgs), func(msg string) {
				log.Printf("Stream message: %v", msg)
				if err := conn.WriteJSON(InvokeResp{
					Type:        "stream_message",
					PackageName: req.PackageName,
					ServiceName: req.ServiceName,
					MethodName:  req.MethodName,
					Data:        msg,
				}); err != nil {
					log.Printf("Can't marshal message: %v", err)
				}
			})
			if err != nil {
				if err := conn.WriteJSON(InvokeResp{
					Type:        "invoke_resp",
					PackageName: req.PackageName,
					ServiceName: req.ServiceName,
					MethodName:  req.MethodName,
					Error:       err.Error(),
				}); err != nil {
					log.Printf("Can't marshal message: %v", err)
				}
			}
		}()

		return nil
	}

	invokeRes, err := reflection.Invoke(ctx, req.Addr, req.GRPCMethod, []byte(req.GRPCArgs))
	if err != nil {
		err := conn.WriteJSON(InvokeResp{
			Type:        "invoke_resp",
			PackageName: req.PackageName,
			ServiceName: req.ServiceName,
			MethodName:  req.MethodName,
			Error:       err.Error(),
		})
		if err != nil {
			return err
		}
		return nil
	}

	if err := conn.WriteJSON(InvokeResp{
		Type:        "invoke_resp",
		PackageName: req.PackageName,
		ServiceName: req.ServiceName,
		MethodName:  req.MethodName,
		Data:        invokeRes,
	}); err != nil {
		return err
	}

	return nil
}
