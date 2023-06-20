package queue

import (
	"testing"

	"github.com/yannis94/bank-root/core"
)

func TestEnqueue(t *testing.T) {
    testQ := BuildQueue()

    fakeTransferDemand := &core.TransferDemand{
        Id: 1,
        FromAccount: "uuid-00000001",
        ToAccount: "uuid-00000002",
        Amount: 1500,
        Accepted: false,
        Closed: false,
    }

    if testQ.Length > 0 {
        t.Logf("Queue's length should be zero and it is: %d", testQ.Length)
    }

    testQ.Enqueue(fakeTransferDemand)

    if testQ.Length != 1 {
        t.Logf("Queue's length should be one and it is: %d", testQ.Length)
    }

    testVal := testQ.Peek()

    if testVal.FromAccount != "uuid-00000001" {
        t.Logf("The from account test faild. Should be uuid-00000001 and it is %s", testVal.FromAccount)
    }
}

func TestDequeue(t *testing.T) {
    testQ := BuildQueue()


    fakeTransferDemand_1 := &core.TransferDemand{
        Id: 1,
        FromAccount: "uuid-00000001",
        ToAccount: "uuid-00000002",
        Amount: 1500,
        Accepted: false,
        Closed: false,
    }

    fakeTransferDemand_2 := &core.TransferDemand{
        Id: 1,
        FromAccount: "uuid-00000001",
        ToAccount: "uuid-00000002",
        Amount: 1500,
        Accepted: false,
        Closed: false,
    }

    val, err := testQ.Dequeue()

    if err == nil {
        t.Logf("Error should'nt be nil, you can't dequeue anything for now. Returned value : %v", val)
    }

    testQ.Enqueue(fakeTransferDemand_1)
    testQ.Enqueue(fakeTransferDemand_2)

    dequeue_1, err := testQ.Dequeue()

    if err != nil {
        t.Logf("Error should be nil but it is :%s", err.Error())
    }

    if dequeue_1.FromAccount != fakeTransferDemand_1.FromAccount {
        t.Log("Wrong value returned.")
    }

    dequeue_2, err := testQ.Dequeue()
    if err != nil {
        t.Logf("Error should be nil but it is :%s", err.Error())
    }

    if dequeue_2.FromAccount != fakeTransferDemand_2.FromAccount {
        t.Log("Wrong value returned. It should be the first value second_value inserted.")
    }

    val2, err := testQ.Dequeue()

    if err == nil {
        t.Logf("Error should'nt be nil, you can't dequeue anything for now. Returned value : %v", val2)
    }
}
