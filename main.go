package main

import (
	"strings"

	"github.com/taction/daprwasmactor/sdk"
)

func main() {
	sdk.RegisterActorHandler(&Actor{})
}

type Actor struct{}

// Handle handles actor invoke
func (a *Actor) Handle(req *sdk.InvokeActorRequest) (data *sdk.InvokeActorResponse) {
	if req.Method == "get" {
		res, err := sdk.GetActorState(&sdk.GetActorStateRequest{
			ActorType: req.ActorType,
			ActorID:   req.ActorID,
			KeyName:   string(req.Data),
		})
		if err != nil {
			return &sdk.InvokeActorResponse{
				StatusCode: 500,
				Data:       []byte(err.Error()),
			}
		}
		return &sdk.InvokeActorResponse{
			StatusCode: 200,
			Data:       res.Data,
		}
	} else {
		datas := strings.Split(string(req.Data), "=")
		if len(datas) != 2 {
			return &sdk.InvokeActorResponse{
				StatusCode: 500,
				Data:       []byte("invalid request data"),
			}
		}
		success := sdk.TransactionalStateOperation(&sdk.TransactionalRequest{
			ActorType: req.ActorType,
			ActorID:   req.ActorID,
			Operations: []sdk.TransactionalOperation{
				{
					Operation: "upsert",
					Request: sdk.OperationRequest{
						Key:   datas[0],
						Value: datas[1],
					},
				},
			},
		})
		if !success {
			return &sdk.InvokeActorResponse{
				StatusCode: 500,
				Data:       []byte("transaction failed"),
			}
		}
		return &sdk.InvokeActorResponse{
			StatusCode: 200,
			Data:       []byte("success"),
		}
	}
}
