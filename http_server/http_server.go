package http_server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lazada/grpc-ui/proto"
	"github.com/lazada/grpc-ui/reflection"
)


func New(addr, staticDir, targetAddr string) *HTTPServer {
	mux := http.NewServeMux()

	s := &HTTPServer{
		addr: addr,
		targetAddr: targetAddr,
		mux: mux,
	}

	mux.HandleFunc("/api/info", s.infoHandler)
	mux.HandleFunc("/api/invoke", s.invokeHandler)
	mux.HandleFunc("/", s.indexHandler)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(staticDir))))

	return s
}


type HTTPServer struct{
	addr string
	targetAddr string
	mux *http.ServeMux
}

type InvokeReq struct {
	ServiceName string `json:"service_name"`
	PackageName string `json:"package_name"`
	MethodName string `json:"method_name"`
	GRPCArgs   []proto.FieldValue `json:"grpc_args"`
}

type InvokeResp struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

type InvokeStreamReq struct {
	GRPCMethod string `json:"grpc_method"`
	GRPCArgs   string `json:"grpc_args"`
}

type InvokeStreamResp struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func (h *HTTPServer) infoHandler(w http.ResponseWriter, r *http.Request) {
	info, err := reflection.GetInfo(r.Context(), h.targetAddr)

	if err != nil {
		log.Printf("Can't get grpc info: %v", err)
		http.Error(w, fmt.Sprintf("Can't get grpc info: %v", err), http.StatusInternalServerError)
		return
	}


	if err := json.NewEncoder(w).Encode(info); err != nil {
		log.Printf("Can't encode json: %v", err)
		return
	}
}

func (h *HTTPServer) invokeHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("stream") == "" {
		h.handleUnary(w, r)
		return
	}
	h.handleStream(w, r)
}

func (h *HTTPServer) handleUnary(w http.ResponseWriter, r *http.Request) {
	req := InvokeReq{}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("json.Unmarshal error: %v", err), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	invokeRes, err := proto.Invoke(r.Context(), h.targetAddr, req.PackageName, req.ServiceName, req.MethodName, req.GRPCArgs)
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

func (h *HTTPServer) handleStream(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Can't upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	log.Print("WebSocket connected")
	defer log.Print("WebSocket disconnected")
	//
	//ctx, cancel := context.WithCancel(r.Context())
	//defer cancel()
	//
	//for {
	//	req := InvokeStreamReq{}
	//	if err := conn.ReadJSON(&req); err != nil {
	//		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
	//			log.Printf("ReadJSON error: %v", err)
	//		}
	//		return
	//	}
	//}

}


func (h *HTTPServer) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func (h *HTTPServer) Start() error {
	return http.ListenAndServe(h.addr, h.mux)
}
