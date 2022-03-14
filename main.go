package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lergo -lbcrypt

#include "ergo/ergo.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	addrStr := C.CString("9hdxkYakTHWXR992umPcvh8bAEGG9Sdoi7uW8TKXk1enXCDFBVJ")
	defer C.free(unsafe.Pointer(addrStr))

	var addrPtr C.AddressPtr
	defer C.free(unsafe.Pointer(&addrPtr))
	C.ergo_lib_address_from_base58(addrStr, &addrPtr)

	var outAddrStr *C.char
	defer C.free(unsafe.Pointer(&outAddrStr))
	C.ergo_lib_address_to_base58(addrPtr, 0, &outAddrStr)

	result := C.GoString(outAddrStr)
	fmt.Printf("Address roundtrip: %s", result)
}
