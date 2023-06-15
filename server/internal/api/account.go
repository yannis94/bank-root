package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yannis94/bank-root/internal/service"
)

func (server *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
    if r.Method == "DELETE" {
        return server.handleDeleteAccount(w, r)
    }

    vars := mux.Vars(r)

    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        errMessage := fmt.Sprintf("%s is not a vaild id.", vars["id"])
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: errMessage })
    }
    
    account, err := server.repo.GetAccountById(id)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error."})
    } 

    if account == nil {
        errMessage := fmt.Sprintf("Account id %d not found.", id)
        return writeJSON(w, http.StatusNotFound, ApiError{ Details: errMessage})
    } 

    return writeJSON(w, http.StatusAccepted, account)

}

func (server *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
    createAccountReq := &service.CreateAccountRequest{}

    if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: err.Error() })
    }
    defer r.Body.Close()

    account := service.NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
    
    accountCreated, err := server.repo.CreateAccount(account)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: err.Error() })
    }

    return writeJSON(w, http.StatusCreated, accountCreated)
}

func (server *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        errMessage := fmt.Sprintf("%s is not a vaild id.", vars["id"])
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: errMessage })
    }

    err = server.repo.DeleteAccount(id)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error."})
    }
    
    return writeJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Account id %d get deleted.", id)})
}
