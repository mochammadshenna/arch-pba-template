package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
func NewEncoder(writer io.Writer) *jsoniter.Encoder {
	return json.NewEncoder(writer)
}
func NewDecoder(reader io.Reader) *jsoniter.Decoder {
	return json.NewDecoder(reader)
}
