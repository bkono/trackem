package main

import (
	"flag"
	"log"

	"git.kono.sh/bkono/trackem"
	pb "git.kono.sh/bkono/trackem/protos"
)

var (
	to      = flag.String("to", "", "comma separated list of recipients for the email")
	from    = flag.String("from", "", "sender email address")
	subject = flag.String("subject", "", "subject line of the email")
	baseURL = flag.String("baseurl", "https://t.kono.sh/em.gif?q=", "baseURL of the tracking server <https://t.kono.sh/evt.gif?q=>")
)

func main() {
	log.Println("Building tracking url...")
	flag.Parse()
	t := &pb.EmailTracker{
		ToAddrs:  *to,
		FromAddr: *from,
		Subject:  *subject,
	}

	enc := trackem.Encode(t)
	turl := *baseURL + enc
	log.Printf("%s\n", turl)
}
