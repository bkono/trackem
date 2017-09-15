package trackem

import (
	"encoding/base64"
	"log"

	pb "git.kono.sh/bkono/trackem/protos"
	"github.com/micro/protobuf/proto"
	"github.com/segmentio/ksuid"
)

func Encode(et *pb.EmailTracker) string {
	if len(et.Id) == 0 {
		et.Id = ksuid.New().String()
	}

	b, err := proto.Marshal(et)
	if err != nil {
		log.Fatal(err)
	}

	return base64.URLEncoding.EncodeToString(b)
}
