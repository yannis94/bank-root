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

    router.HandleFunc("/client/signup", customHandler(server.handleCreateClient)).Methods("POST")
    router.HandleFunc("/client/signin", customHandler(server.handleClientSignIn)).Methods("POST")
    router.HandleFunc("/client", customHandler(server.handleDeleteClient)).Methods("DELETE")
    router.HandleFunc("/client/{id}", jwtClientAuth(server.auth, customHandler(server.handleGetClientById))).Methods("GET")
    router.HandleFunc("/client/transfer-demand", jwtClientAuth(server.auth, customHandler(server.handleTransferDemand))).Methods("GET")

    router.HandleFunc("/account", jwtClientAuth(server.auth, customHandler(server.handleCreateAccount))).Methods("POST")
    router.HandleFunc("/account/{account_number}", customHandler(server.handleGetAccountByNumber)).Methods("GET")
    router.HandleFunc("/transfer", customHandler(server.handleTransferDemand)).Methods("POST")

    router.HandleFunc("/transfer", moiraiAuth(customHandler(server.handleGetAcceptedTransferDemands))).Methods("GET")
    router.HandleFunc("/transfer", moiraiAuth(customHandler(server.handleUpdateTransferDemand))).Methods("PUT")

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

func customHandler(f apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            writeJSON(w, http.StatusBadRequest, ApiError{ Details: "Endpoint does not exist." })
        }
    }
}

func jwtClientAuth(auth *auth.AuthService, f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("access_token")
        token := cookie.Value

        if err != nil {
            writeJSON(w, http.StatusForbidden, ApiError{ Details: "No access token found." })
        } else if _, err = auth.IsTokenValid(token, "client"); err != nil {
            if err.Error() == "Token is invalid." {
                newToken, err := auth.RefreshToken(token)
                log.Println(err)
                if err != nil {
                    writeJSON(w, http.StatusForbidden, ApiError{ Details: "Access token invalid." })
                } else {
                    cookie.Value = newToken
                    http.SetCookie(w, cookie)
                    f(w, r)
                }
            } else {
                writeJSON(w, http.StatusForbidden, ApiError{ Details: "Access token invalid." })
            }
        } else {
            f(w, r)
        }
    }
}

func moiraiAuth(f http.HandlerFunc) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        // check incoming request's host (should be moirai)
        f(w, r)
    }
}
