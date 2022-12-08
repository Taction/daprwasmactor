package mem

import (
	"reflect"
	"unsafe"
)

var (
	Buf        = make([]byte, DefaultLen)
	DefaultPtr = uintptr(unsafe.Pointer(&Buf[0]))
	DefaultLen = uint32(4096)
)

func StringToPtr(s string) (uintptr, uint32) {
	if s == "" {
		return DefaultPtr, 0
	}
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return unsafePtr, uint32(len(buf))
}

func GetString(size uint32) (result string) {
	return string(Buf[:size]) // string will copy the buffer.
}

func GetBytes(size uint32) (result []byte) {
	result = make([]byte, size)
	copy(result, Buf)
	return
}

//go:inline
func BytesToPointer(s []byte) uintptr {
	return (*(*reflect.SliceHeader)(unsafe.Pointer(&s))).Data
}

//go:inline
func StringToPointer(s string) uintptr {
	return (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data
}
