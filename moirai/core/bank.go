package core

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type BankAPI struct {
    url string
}

func NewBankAPI() *BankAPI {
    apiUri := os.Getenv("BANK_API_URI")
    authUrl := apiUri + "/moirai-auth"
    
    header := make(map[string]string)
    header["Content-Type"] = "application/json"

    var authData bytes.Buffer
    cred := map[string]string{"login": "moirai", "password": os.Getenv("MOIRAI_SECRET")}

    json.NewEncoder(&authData).Encode(cred)

    res, err := httpReq("POST", authUrl, header, &authData)

    if err != nil || res.StatusCode != 200 {
        log.Fatalf("Unable to connect to bank api, error :%s", err.Error())
    }

    return &BankAPI{ url: apiUri}
}

func httpReq(method string, endpoint string, headers map[string]string, body *bytes.Buffer) (*http.Response, error) {
    client := &http.Client{
        Transport: &http.Transport{
            DisableKeepAlives: true,
        },
        Timeout: 2 * time.Second,
    }

    req, _ := http.NewRequest(method, endpoint, body)

    for key, value := range headers {
        req.Header.Add(key, value)
    }

    return client.Do(req)
}

func (api *BankAPI) GetTransferDemand() ([]TransferDemand, error) {
    reqUrl := api.url + "/transfer"
    res, err := httpReq("GET", reqUrl, nil, nil)

    if err != nil {
        log.Println(err)
        return nil, err
    }

    var transferDemands []TransferDemand

    json.NewDecoder(res.Body).Decode(&transferDemands)
    defer res.Body.Close()

    return transferDemands, nil
}

func (api *BankAPI) GetAccount(accountNumber string) (*Account, error) {
    reqUrl := api.url + "/account/" + accountNumber

    res, err := httpReq("GET", reqUrl, nil, nil)

    if err != nil {
        log.Println(err)
        return nil, err
    }

    var account Account

    json.NewDecoder(res.Body).Decode(&account)
    defer res.Body.Close()

    return &account, nil
}

func (api *BankAPI) SendTransferValidation(transfer TransferDemand) error {
    reqUrl := api.url + "/transfer"
    header := make(map[string]string)
    header["Content-Type"] = "application/json"

    var data bytes.Buffer
    err := json.NewEncoder(&data).Encode(transfer)

    if err != nil {
        log.Println(err)
        return err
    }

    _, err = httpReq("PUT", reqUrl, header, &data)

    return err
}

func (api *BankAPI) UpdateAccount(client_account Account) error {
    reqUrl := api.url + "/account"
    header := make(map[string]string)
    header["Content-Type"] = "application/json"

    var data bytes.Buffer
    err := json.NewEncoder(&data).Encode(client_account)

    if err != nil {
        log.Println(err)
        return err
    }

    _, err = httpReq("PUT", reqUrl, header, &data)

    return err
}
