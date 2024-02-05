package proxywasm

//#include <stdlib.h>
//void free(void *ptr);
import "C"
import (
	"reflect"
	"unsafe"
)

func cleanup_off_heap_bytes(data *[]byte) {
	C.free(unsafe.Pointer(&(*data)[0]))
	//nolint
	*((*reflect.SliceHeader)(unsafe.Pointer(data))) = reflect.SliceHeader{
		Data: uintptr(0),
		Len:  0,
		Cap:  0,
	}
}
