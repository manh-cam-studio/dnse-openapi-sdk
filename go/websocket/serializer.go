package websocket

import (
	"encoding/json"
	msgpack "github.com/vmihailenco/msgpack/v5"
)

// Serializer định nghĩa chuẩn mã hóa cho các frame socket
type Serializer interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

type JsonSerializer struct{}

func (s *JsonSerializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (s *JsonSerializer) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (s *JsonSerializer) Name() string {
	return "json"
}

type MsgpackSerializer struct{}

func (s *MsgpackSerializer) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (s *MsgpackSerializer) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

func (s *MsgpackSerializer) Name() string {
	return "msgpack"
}

// GetSerializer là hàm factory đơn giản
func GetSerializer(encoding string) Serializer {
	if encoding == "msgpack" {
		return &MsgpackSerializer{}
	}
	return &JsonSerializer{}
}
