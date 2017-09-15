package trackem

import (
	"encoding/base64"

	pb "git.kono.sh/bkono/trackem/protos"
	"github.com/micro/protobuf/proto"
)

func Decode(encoded string) (*pb.EmailTracker, error) {
	raw, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	et := new(pb.EmailTracker)
	if err := proto.Unmarshal(raw, et); err != nil {
		return nil, err
	}

	return et, nil
}
