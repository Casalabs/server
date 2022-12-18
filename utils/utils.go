package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadENV(key string) string {
	err := godotenv.Load(".env")
	HandleErr(err)
	return os.Getenv(key)
}
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// data -> byte
func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

// byte -> data
// FromBytes takes an interface and data and the will encode the data to the interface
func FromBytes(i interface{}, data []byte) error {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	err := encoder.Decode(i)
	return err
}

func Splitter(s string, sep string, i int) string {
	result := strings.Split(s, sep)
	if len(result)-1 < i {
		return ""
	}
	return result[i]
}

// json -> byte
func ToJSON(i interface{}) []byte {
	r, err := json.MarshalIndent(i, "", "")
	HandleErr(err)
	return r
}
