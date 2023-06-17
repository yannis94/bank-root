package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/yannis94/bank-root/internal/helper"
	"github.com/yannis94/bank-root/internal/service"
)


func (server *ApiServer) handleCreateClient(w http.ResponseWriter, r *http.Request) error {
    createClientReq := &service.CreateClientRequest{}

    if err := json.NewDecoder(r.Body).Decode(&createClientReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "JSON invalid format" })
    }

    defer r.Body.Close()

    testClient, err := server.repo.GetClientByEmail(createClientReq.Email)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }
    log.Println(testClient)

    if testClient.Email != "" {
        return writeJSON(w, http.StatusForbidden, ApiError{ Details: "Email already used." })
    }

    if createClientReq.Password != createClientReq.PasswordVerify {
        return writeJSON(w, http.StatusForbidden, ApiError{ Details: "Password not equals." })
    }

    hashPwd, err := helper.HashPassword(createClientReq.Password)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Internal server error, unable to hash the password" })
    }

    client := service.NewClient(createClientReq.FirstName, createClientReq.LastName, createClientReq.Email, hashPwd)
    clientCreated, err := server.repo.CreateClient(client)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }

    return writeJSON(w, http.StatusCreated, clientCreated)
}

func (server *ApiServer) handleClientSignIn(w http.ResponseWriter, r *http.Request) error {
    clientSiginReq := &service.ClientSignInRequest{}

    if err := json.NewDecoder(r.Body).Decode(&clientSiginReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "Invalid JSON request." })
    }

    client, err := server.repo.GetClientByEmail(clientSiginReq.Email)

    if err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "Database error." })
    }

    if client == nil {
        return writeJSON(w, http.StatusNotFound, ApiError{ Details: "This email address is not link to any account." })
    }

    if !helper.ConfirmPassword(clientSiginReq.Password, client.Password) {
        return writeJSON(w, http.StatusForbidden, ApiError{ Details: "Incorrect password." })
    }
    
    jti := uuid.New().String()

    token, err := server.auth.CreateJWT("client", jti)
    refresh_token, err := server.auth.CreateRefreshToken("client")

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Unable to generate an access token." })
    }

    session := service.NewSession(jti, refresh_token)

    err = server.repo.CreateSession(session)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Unable to generate an access token." })
    }

    cookie := &http.Cookie{
        Name: "access_token",
        Value: token,
        Path: "/",
        Expires: time.Now().Add(24 * time.Hour),
        HttpOnly: true,
    }

    http.SetCookie(w, cookie)

    return writeJSON(w, http.StatusAccepted, client)
}

func (server *ApiServer) handleGetClientById(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        errMessage := fmt.Sprintf("%s is not a valid id.", vars["id"])
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: errMessage })
    }

    clientFound, err := server.repo.GetClientById(id)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    } 

    if clientFound == nil {
        return writeJSON(w, http.StatusNotFound, ApiError{ Details: "Client not found." })
    }

    return writeJSON(w, http.StatusFound, clientFound)
}

func (server *ApiServer) handleDeleteClient(w http.ResponseWriter, r *http.Request) error {
    deleteClientReq := &service.DeleteClientRequest{}

    if err := json.NewDecoder(r.Body).Decode(&deleteClientReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "JSON invalid format." })
    }

    account_number, err := server.repo.GetAccountNumber(deleteClientReq.Id)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }

    if account_number == "" {
        return writeJSON(w, http.StatusNotFound, ApiError{ Details: "Client not found." })
    }

    err = server.repo.CreateClosedAccount(account_number)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }

    client, err := server.repo.DeleteClient(deleteClientReq.Id)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }

    return writeJSON(w, http.StatusOK, client)
}
