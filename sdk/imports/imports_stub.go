//go:build !tinygo.wasm

package imports

func GetActorState(ptr uintptr, size uint32, resPtr uintptr) (resSize uint32) {
	return 0
}

func TransactionalStateOperation(ptr uintptr, size uint32) int32 {
	return 1
}

func CallActor(ptr uintptr, size uint32, resPtr uintptr) (resSize uint32) {
	return 0
}

func ActorRequest(reqPtr uintptr) {}

func ActorResponse(ptr uintptr, len uint32) {}

func ActorError(ptr uintptr, len uint32) {}
