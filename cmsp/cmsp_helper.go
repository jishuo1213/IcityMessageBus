package cmsp

//#cgo CPPFLAGS:-I./include
//#cgo LDFLAGS:-L/home/fan/go/src/IcityMessageBus/cmsp/lib -lIcityCmspLibrary
//#include <icity_cmsp.h>
//#include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
	"log"
)

var queueMap map[string]*C.char

func init() {
	queueMap = make(map[string]*C.char)
}

func PutMessageIntoQueue(topic string, msg []byte) error {
	goQueue, err := getQueue(topic)
	if err != nil {
		return err
	}
	//if queue, found := queueMap[topic]; found {
	//	goQueue = queue
	//} else {
	//	cTopic := C.CString(topic)
	//	defer C.free(unsafe.Pointer(cTopic))
	//
	//	goQueue = C.openQueue(cTopic)
	//
	//	if goQueue != nil {
	//		queueMap[topic] = goQueue
	//	} else {
	//		return errors.New("open queue failed")
	//	}
	//}

	cMsg := (*C.char)(unsafe.Pointer(&msg[0]))
	ret := C.putOneMessageToQueue(goQueue, cMsg, C.uint(len(msg)))

	log.Print("put msg into queue ret =", ret)
	if int(ret) == 0 {
		return nil
	} else {
		return errors.New("put message failed")
	}
}

func getQueue(topic string) (*C.char, error) {
	var goQueue *C.char
	if queue, found := queueMap[topic]; found {
		goQueue = queue
	} else {
		cTopic := C.CString(topic)
		goQueue = C.openQueue(cTopic)
		defer C.free(unsafe.Pointer(cTopic))

		if goQueue != nil {
			queueMap[topic] = goQueue
		} else {
			return nil, errors.New("open queue failed")
		}
	}
	return goQueue, nil
}

func GetQueue(topic string) error {
	queue, err := getQueue(topic)
	log.Print(queue)
	return err
}

func ReadOneMessageFromQueue(topic string) ([]byte, error) {
	goQueue, err := getQueue(topic)
	if err != nil {
		return nil, err
	}

	var msgLength int;
	cMessage := C.getOneMessageFromQueue(goQueue, unsafe.Pointer(&msgLength))
	if cMessage != nil {
		goMsg := C.GoBytes(unsafe.Pointer(cMessage), C.int(msgLength))
		defer C.free(unsafe.Pointer(cMessage))
		return goMsg, nil
	} else {
		return nil, errors.New("get message failed")
	}
}
