package shared

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func EncodeStruct(item interface{}) (buff bytes.Buffer) {

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(item)
	if err != nil {
		log.Fatal("encode:", err)
	}
	return
}

func DecodeStruct[T interface{}](itemByte []byte) T {

	var item T

	buff := bytes.Buffer{}
	buff.Write(itemByte)

	dec := gob.NewDecoder(&buff)

	err := dec.Decode(&item)

	if err != nil {
		log.Fatal("decode:", err)
	}

	return item
}

func HandleError(err error, successMessage string) {
	if err == nil {
		fmt.Println(successMessage)
		return
	}

	log.Fatal(fmt.Errorf("error: %v", err))
}
