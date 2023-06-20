package core

type TransferDemand struct {
    Id int
    Closed  bool
    FromAccount string
    ToAccount string
    Amount int
    Accepted bool
}

type Transfer struct {
    DemandId int
    Done bool
}
