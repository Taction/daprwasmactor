//go:build tinygo.wasm

package imports

//go:wasm-module dapr_actor
//go:export get_actor_state
func GetActorState(ptr uintptr, size uint32, resPtr uintptr) (resSize uint32)

//go:wasm-module dapr_actor
//go:export set_actor_state
func TransactionalStateOperation(ptr uintptr, size uint32) int32

//go:wasm-module dapr_actor
//go:export call_actor
func CallActor(ptr uintptr, size uint32, resPtr uintptr) (resSize uint32)

//go:wasm-module dapr_actor
//go:export actor_request
func ActorRequest(reqPtr uintptr)

//go:wasm-module dapr_actor
//go:export actor_response
func ActorResponse(ptr uintptr, len uint32) //nolint

//go:wasm-module dapr_actor
//go:export actor_error
func ActorError(ptr uintptr, len uint32)
