package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yannis94/bank-root/internal/auth"
	"github.com/yannis94/bank-root/internal/repository"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
    Details string `json:"error"`
}

type ApiServer struct {
    port string
    repo repository.Storage
    auth *auth.AuthService
}

func NewApiServer(port string, repo repository.Storage, auth *auth.AuthService) *ApiServer {
    return &ApiServer{
        port: port,
        repo: repo,
        auth: auth,
    }
}

func (server *ApiServer) Start() {
    router := mux.NewRouter()

    router.HandleFunc("/tkn", httpHandleFuncTransform(server.getClientToken)).Methods("GET")

    router.HandleFunc("/client", httpHandleFuncTransform(server.handleCreateClient)).Methods("POST")
    router.HandleFunc("/client", httpHandleFuncTransform(server.handleDeleteClient)).Methods("DELETE")
    router.HandleFunc("/client/{id}", jwtClientAuth(httpHandleFuncTransform(server.handleGetClientById))).Methods("GET")
    router.HandleFunc("/account", httpHandleFuncTransform(server.handleCreateAccount)).Methods("POST")
    router.HandleFunc("/account/{id}", httpHandleFuncTransform(server.handleGetAccount)).Methods("GET")
    router.HandleFunc("/demand", httpHandleFuncTransform(server.handleTransferDemand)).Methods("POST")
    router.HandleFunc("/demand", httpHandleFuncTransform(server.handleGetAcceptedTransferDemands)).Methods("GET")
    router.HandleFunc("/demand", httpHandleFuncTransform(server.handleUpdateTransferDemand)).Methods("PUT")
    router.HandleFunc("/transfer", httpHandleFuncTransform(server.handlerCreateTransfer)).Methods("POST")

    log.Println("Server listening on port", server.port)

    if err := http.ListenAndServe(server.port, router); err != nil {
        log.Fatalf("Server listening on port %s.\nError:%v", server.port, err)
    }

}

func writeJSON(w http.ResponseWriter, status int, content any) error {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(content)
}

func httpHandleFuncTransform(f apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            writeJSON(w, http.StatusBadRequest, ApiError{ Details: err.Error() })
        }
    }
}

func jwtClientAuth(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("JWT client authentication")
        f(w, r)
    }
}
