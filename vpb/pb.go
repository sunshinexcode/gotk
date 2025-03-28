package vpb

import "google.golang.org/protobuf/proto"

type (
	Message = proto.Message
)

func Marshal(m Message) ([]byte, error) {
	return proto.Marshal(m)
}

func Unmarshal(b []byte, m Message) error {
	return proto.Unmarshal(b, m)
}
