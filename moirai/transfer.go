package main

type transferDemand struct {
    Id int
    Closed  bool
    FromAccount string
    ToAccount string
    Amount string
    Accepted bool
}

type transfer struct {
    DemandId int
    Done bool
}
