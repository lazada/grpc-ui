package http_server

import (
	"io/ioutil"
	"net/http"

	proto "github.com/golang/protobuf/proto"
	"github.com/lazada/grpc-ui/reflection"
	"google.golang.org/grpc/status"
)

func New(addr string) *HTTPServer {
	mux := http.NewServeMux()

	s := &HTTPServer{
		addr: addr,
		mux:  mux,
	}

	mux.HandleFunc("/api/invoke", s.invokeHandler)
	mux.HandleFunc("/api/reflection", s.reflectionHandler)
	mux.Handle("/", NewHTTPHandler())

	return s
}

type HTTPServer struct {
	addr string
	mux  *http.ServeMux
}

func httpError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func respondWithProto(w http.ResponseWriter, msg proto.Message) {
	marshalled, err := proto.Marshal(msg)

	if err != nil {
		httpError(w, http.StatusInternalServerError)
		return
	}

	w.Write(marshalled)
}

func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		httpError(w, http.StatusMethodNotAllowed)
		return true
	}

	return false
}

func (h *HTTPServer) invokeHandler(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r, http.MethodPost) {
		return
	}

	defer r.Body.Close()

	target := r.FormValue("host")

	if target == "" {
		httpError(w, http.StatusBadRequest)
		return
	}

	buff, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	req := &InvokeRequest{}

	err = proto.Unmarshal(buff, req)

	if err != nil {
		httpError(w, http.StatusBadRequest)
		return
	}

	result, err := reflection.Invoke(r.Context(), target, req.Method, req.Payload)

	resp := &InvokeResponse{}

	if err != nil {
		s, _ := status.FromError(err)

		if s != nil {
			resp.Response = &InvokeResponse_Error{
				Error: &Error{
					Message: s.Message(),
				},
			}
		} else {
			{
				resp.Response = &InvokeResponse_Error{
					Error: &Error{
						Message: err.Error(),
					},
				}
			}
		}
	} else {
		resp.Response = &InvokeResponse_Payload{
			Payload: result,
		}
	}

	respondWithProto(w, resp)
}

func (h *HTTPServer) reflectionHandler(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r, http.MethodGet) {
		return
	}

	target := r.FormValue("host")

	if target == "" {
		httpError(w, http.StatusBadRequest)
		return
	}

	data, err := reflection.GetReflection(r.Context(), target)

	resp := &ReflectionResponse{}

	if err != nil {
		resp.Response = &ReflectionResponse_Error{
			Error: &Error{
				Message: err.Error(),
			},
		}
	} else {
		resp.Response = &ReflectionResponse_Reflection{
			Reflection: &Reflection{
				Service:        data.Services,
				FileDescriptor: data.FileDescriptors,
			},
		}
	}

	respondWithProto(w, resp)
}

func (h *HTTPServer) Start() error {
	return http.ListenAndServe(h.addr, h.mux)
}
