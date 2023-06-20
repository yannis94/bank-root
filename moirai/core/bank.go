package core

type bankAPI struct {
    url string
    token string
}

func (api *bankAPI) GetTransferDemand() ([]*TransferDemand, error) {
    return nil, nil
}

func (api *bankAPI) TransferValidation() error {
    return nil
}
