package operations

import (
	"strings"
	"time"
)

type operation struct {
	OperationType string    `json:"type"`
	A             float64   `json:"a"`
	B             float64   `json:"b"`
	Result        float64   `json:"result"`
	CompletedAt   time.Time `json:"completed_at"`
}

var recentOperations = []operation{}

func appendOperation(operationType string, a float64, b float64, result float64) (operation, operation) {

	op := operation{
		OperationType: operationType,
		A:             a,
		B:             b,
		Result:        result,
		CompletedAt:   time.Now(),
	}

	recentOperations = append(recentOperations, op)

	var lastop operation

	if len(recentOperations) >= 2 {
		lastop = recentOperations[len(recentOperations)-2]
	}

	return op, lastop
}

func OperationsByType(operationType string) []operation {
	ops := []operation{}

	for _, op := range recentOperations {
		if strings.Compare(op.OperationType, operationType) == 0 {
			ops = append(ops, op)
		}
	}

	return ops
}

func Operations() []operation {

	return recentOperations
}
