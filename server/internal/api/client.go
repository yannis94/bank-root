package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yannis94/bank-root/internal/service"
)

func (server *ApiServer) handleCreateClient(w http.ResponseWriter, r *http.Request) error {
    createClientReq := &service.CreateClientRequest{}

    if err := json.NewDecoder(r.Body).Decode(&createClientReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "JSON invalid format" })
    }

    defer r.Body.Close()

    //password match and hash here...

    client := service.NewClient(createClientReq.FirstName, createClientReq.LastName, createClientReq.Email, createClientReq.Password)
    clientCreated, err := server.repo.CreateClient(client)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }

    return writeJSON(w, http.StatusCreated, clientCreated)
}

func (server *ApiServer) handleGetClientById(w http.ResponseWriter, r *http.Request) error {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])

    if err != nil {
        errMessage := fmt.Sprintf("%d is not a valid id.", vars["id"])
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

    client, err := server.repo.DeleteClient(deleteClientReq.Id)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }

    return writeJSON(w, http.StatusOK, client)
}
