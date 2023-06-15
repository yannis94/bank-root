package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yannis94/bank-root/internal/repository"
	"github.com/yannis94/bank-root/internal/service"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
    Details string
}

type ApiServer struct {
    port string
    repo repository.Storage
}

func NewApiServer(port string, repo repository.Storage) *ApiServer {
    return &ApiServer{
        port: port,
        repo: repo,
    }
}

func (server *ApiServer) Start() {
    router := mux.NewRouter()

    router.HandleFunc("/account", httpHandleFuncTransform(server.handleAccount))
    router.HandleFunc("/account/{id}", httpHandleFuncTransform(server.handleGetAccount))

    log.Println("Server listening on port", server.port)

    if err := http.ListenAndServe(server.port, router); err != nil {
        log.Fatalf("Server listening on port %s.\nError:%v", server.port, err)
    }

}

func (server *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
    switch r.Method {
        case "GET":
            return server.handleGetAccount(w, r)
        case "POST":
            return server.handleCreateAccount(w, r)
        case "DELETE":
            return server.handleDeleteAccount(w, r)
        default:
            return errors.New("Method not allowed")
    }
}

func (server *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
    account := service.NewAccount("Yannis", "Bgci")
    return writeJSON(w, http.StatusAccepted, account)
}

func (server *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
    createAccountReq := &service.CreateAccountRequest{}

    if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: err.Error() })
    }

    account := service.NewAccount(createAccountReq.FirstName, createAccountReq.LastName)

    if err := server.repo.CreateAccount(account); err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: err.Error() })
    }

    return writeJSON(w, http.StatusCreated, account)
}

func (server *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
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

