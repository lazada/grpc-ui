package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"context"

	"github.com/gorilla/websocket"
	"github.com/lazada/grpc-ui/reflection"
	"github.com/lazada/grpc-ui/proto"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

type InvokeReq struct {
	Addr       string `json:"addr"`
	ServiceName string `json:"service_name"`
	PackageName string `json:"package_name"`
	MethodName string `json:"method_name"`
	GRPCArgs   map[string]interface{} `json:"grpc_args"`
}

type InvokeResp struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

type InvokeStreamReq struct {
	Addr       string `json:"addr"`
	GRPCMethod string `json:"grpc_method"`
	GRPCArgs   string `json:"grpc_args"`
}

type InvokeStreamResp struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
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
		info, err := reflection.GetInfo(r.Context(), addr)

		if err != nil {
			log.Printf("Can't get grpc info: %v", err)
			http.Error(w, fmt.Sprintf("Can't get grpc info: %v", err), http.StatusInternalServerError)
			return
		}


		if err := json.NewEncoder(w).Encode(info); err != nil {
			log.Printf("Can't encode json: %v", err)
			return
		}
	})

	http.HandleFunc("/invoke", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("stream") == "" {
			handleUnary(w, r)
			return
		}
		handleStream(w, r)
	})

	http.ListenAndServe(":3000", nil)
}

func handleUnary(w http.ResponseWriter, r *http.Request) {
	req := InvokeReq{}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("json.Unmarshal error: %v", err), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	invokeRes, err := proto.Invoke(r.Context(), req.Addr, req.PackageName, req.ServiceName, req.MethodName, req.GRPCArgs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		enc.Encode(&InvokeResp{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}
	enc.Encode(&InvokeResp{
		Status: "ok",
		Data:   invokeRes,
	})
}

func handleStream(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Can't upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	log.Print("WebSocket connected")
	defer log.Print("WebSocket disconnected")

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	_ = ctx
	for {
		req := InvokeStreamReq{}
		if err := conn.ReadJSON(&req); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Printf("ReadJSON error: %v", err)
			}
			return
		}

		//go func() {
		//	err = reflection.InvokeStream(ctx, req.Addr, req.GRPCMethod, []byte(req.GRPCArgs), func(msg string) {
		//		conn.WriteJSON(&InvokeStreamResp{
		//			Status: "ok",
		//			Data:   msg,
		//		})
		//
		//	})
		//	if err != nil {
		//		conn.WriteJSON(&InvokeStreamResp{
		//			Status: "error",
		//			Error:  err.Error(),
		//		})
		//	}
		//}()
	}

}
