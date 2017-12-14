package cmsp

//#cgo CFLAGS:-I./include
//#cgo LDFLAGS:-L./lib -lcmsp -lfreeMq
/**
#include <mqCli.h>
 */

import "C"

func putMessageIntoQueue(topic string,msg []byte) int {
	cTopic := C.CString(topic)
	cQueue := C.CString("")
	C.openCmsp(cTopic,cQueue)
}
