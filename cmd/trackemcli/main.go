package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"

	pb "git.kono.sh/bkono/trackem/protos"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/ksuid"
)

var (
	to      = flag.String("to", "", "comma separated list of recipients for the email")
	from    = flag.String("from", "", "sender email address")
	subject = flag.String("subject", "", "subject line of the email")
	baseURL = flag.String("baseurl", "https://t.kono.sh/evt.gif?q=", "baseURL of the tracking server <https://t.kono.sh/evt.gif?q=>")
)

func main() {
	fmt.Println("Building tracking url...")
	flag.Parse()
	t := &pb.EmailTracker{
		Id:      ksuid.New().String(),
		To:      *to,
		From:    *from,
		Subject: *subject,
	}

	b, err := proto.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	enc := base64.URLEncoding.EncodeToString(b)
	turl := *baseURL + enc
	fmt.Printf("%s", turl)
}
