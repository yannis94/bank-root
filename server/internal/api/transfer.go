package api

import (
	"encoding/json"
	"net/http"

	"github.com/yannis94/bank-root/internal/service"
)

func (server *ApiServer) handleTransferDemand(w http.ResponseWriter, r *http.Request) error {
    transferReq := &service.TransferDemandRequest{}

    if err := json.NewDecoder(r.Body).Decode(&transferReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "Bad request." })
    }
    defer r.Body.Close()

    transferDemand := service.NewTransferDemand(transferReq.Amount, transferReq.FromAccount, transferReq.ToAccount, transferReq.Message)
    server.repo.CreateTransferDemand(transferDemand)

    return writeJSON(w, http.StatusOK, map[string]string{"message": "Transfer success."})
}

func (server *ApiServer) handleUpdateTransferDemand(w http.ResponseWriter, r *http.Request) error {
    updateTransferDemandReq := &service.TransferDemandUpdateRequest{}

    if err := json.NewDecoder(r.Body).Decode(updateTransferDemandReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "Invalid request" })
    }

    transferDemand, err := server.repo.GetTransferDemandById(updateTransferDemandReq.TransferDemandId)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }

    transferDemand.Accepted = updateTransferDemandReq.Acceped

    if !transferDemand.Accepted {
        transferDemand.Closed = true
    }
    
    err = server.repo.UpdateTransferDemand(transferDemand)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error." })
    }

    return nil
}

func (server *ApiServer) handleGetAcceptedTransferDemands(w http.ResponseWriter, r *http.Request) error {
    demands, err := server.repo.GetAcceptedTransferDemands()

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error, unable to get accepted transfer demands." })
    }

    return writeJSON(w, http.StatusOK, demands)
}

func (server *ApiServer) handlerCreateTransfer(w http.ResponseWriter, r *http.Request) error {
    transferReq := &service.TransferRequest{}

    if err := json.NewDecoder(r.Body).Decode(&transferReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "Invalid JSON format." })
    }

    transfer := service.NewTransfer(transferReq.DemandId, transferReq.Done)

    err := server.repo.CreateTransfer(transfer)

    if err != nil {
        return writeJSON(w, http.StatusInternalServerError, ApiError{ Details: "Database error, unable to find transfer demand." })
    }


    return writeJSON(w, http.StatusOK, map[string]string{"message": "Transfer save in database" })
}
