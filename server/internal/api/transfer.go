package api

import (
	"encoding/json"
	"net/http"

	"github.com/yannis94/bank-root/internal/service"
)

func (server *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
    transferReq := &service.TransferRequest{}

    if err := json.NewDecoder(r.Body).Decode(&transferReq); err != nil {
        return writeJSON(w, http.StatusBadRequest, ApiError{ Details: "Bad request." })
    }
    defer r.Body.Close()

    return writeJSON(w, http.StatusOK, map[string]string{"message": "Transfer success."})
}
