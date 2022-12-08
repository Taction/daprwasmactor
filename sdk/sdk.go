package sdk

import (
	msgpack "github.com/wapc/tinygo-msgpack"

	"github.com/taction/daprwasmactor/sdk/imports"
	"github.com/taction/daprwasmactor/sdk/mem"
)

var actor Actor

type Actor interface {
	Handle(req *InvokeActorRequest) (data *InvokeActorResponse)
}

func RegisterActorHandler(a Actor) {
	actor = a
}

//export __actor_call
func ActorCalled(size uint32) bool {
	reqbytes := make([]byte, size)
	imports.ActorRequest(mem.BytesToPointer(reqbytes))
	d := msgpack.NewDecoder(reqbytes)
	req, err := DecodeInvokeActorRequest(&d)
	if err != nil {
		return fail(err.Error())
	}
	res := actor.Handle(&req)
	resbytes, err := res.Encode()
	if err != nil {
		return fail(err.Error())
	}
	imports.ActorResponse(mem.BytesToPointer(resbytes), uint32(len(resbytes)))
	return true
}

func fail(errorMessage string) bool {
	imports.ActorError(mem.StringToPointer(errorMessage), uint32(len(errorMessage)))
	return false
}

func GetActorState(req *GetActorStateRequest) (data *GetActorStateResponse, err error) {
	reqBytes, err := req.Encode()
	if err != nil {
		return nil, err
	}
	resLen := imports.GetActorState(mem.BytesToPointer(reqBytes), uint32(len(reqBytes)), mem.DefaultPtr)

	resBytes := mem.GetBytes(resLen)
	enc := msgpack.NewDecoder(resBytes)
	d, err := DecodeGetActorStateResponse(&enc)
	return &d, err
}

func TransactionalStateOperation(req *TransactionalRequest) bool {
	reqBytes, err := req.Encode()
	if err != nil {
		return false
	}
	success := imports.TransactionalStateOperation(mem.BytesToPointer(reqBytes), uint32(len(reqBytes)))
	return success == 1
}

func CallActor(req *InvokeActorRequest) (data *InvokeActorResponse, err string) {
	return
}

//type GetActorStateRequest struct {
//	ActorType string
//	ActorID   string
//	KeyName   string
//}
//
//type GetActorStateResponse struct {
//	Data []byte
//}
//
//type InvokeActorRequest struct {
//	ActorType   string
//	ActorID     string
//	Method      string
//	Data        []byte
//	ContentType string
//}
//
//type InvokeActorResponse struct {
//	Data        []byte
//	ContentType string
//	StatusCode  int32
//}
//
//func DecodeInvokeActorRequest(d *msgpack.Decoder) (i InvokeActorRequest, e error) {
//	return
//}
