package trackem

import (
	"log"
	"net/http"
)

// GIF is a transparent 1x1 tracking pixel
var GIF = []byte{
	71, 73, 70, 56, 57, 97, 1, 0, 1, 0, 128, 0, 0, 0, 0, 0,
	255, 255, 255, 33, 249, 4, 1, 0, 0, 0, 0, 44, 0, 0, 0, 0,
	1, 0, 1, 0, 0, 2, 1, 68, 0, 59,
}

func GetEMReq(w http.ResponseWriter, r *http.Request) {
	log.Println("received em.gif request")
	q := r.URL.Query()["q"]
	if len(q) == 0 {
		log.Println("q param is missing")
		w.WriteHeader(422)
		w.Write([]byte("query param missing"))
		return
	}

	log.Printf("processing q = %v", q)
	writeGIF(w)
	return
}

func writeGIF(w http.ResponseWriter) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Content-Type", "image/gif")
	_, err := w.Write(GIF)
	return err
}
