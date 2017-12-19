package utils

import (
	"hash"
	"crypto/sha256"
	"encoding/json"
)

var h hash.Hash

func init() {
	h = sha256.New()
}

func EncodeObject(object interface{}) ([]byte, error) {
	//buffer := bytes.NewBuffer(make([]byte, 0, 100))
	//encoder := gob.NewEncoder(buffer)
	//err := encoder.Encode(object)
	//if err != nil {
	//	return nil, err
	//}
	return json.Marshal(object)

}

func DecodeObject(object interface{}, data []byte) (error) {
	//buffer := bytes.NewBuffer(data)
	//decoder := gob.NewDecoder(buffer)
	//err := decoder.Decode(object)
	return json.Unmarshal(data, object)
}

func DigestMessage(msg []byte) (string, error) {
	h.Reset()
	digest := h.Sum(msg)
	return string(digest), nil
}
