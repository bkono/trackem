package trackem

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	table = "email_trackers"
	dyndb = dynamodb.New(session.New(aws.NewConfig()))
)

// GIF is a transparent 1x1 tracking pixel
var GIF = []byte{
	71, 73, 70, 56, 57, 97, 1, 0, 1, 0, 128, 0, 0, 0, 0, 0,
	255, 255, 255, 33, 249, 4, 1, 0, 0, 0, 0, 44, 0, 0, 0, 0,
	1, 0, 1, 0, 0, 2, 1, 68, 0, 59,
}

// GetEMReq is an http.HandleFunc for recording an email tracking pixel request. Responds with a 1x1 transparent gif, or rejects as 422 if no query param provided.
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
	et, err := Decode(q[0])
	if err != nil {
		log.Printf("failed parsing q, err = %v", err)
		writeGIF(w)
		return
	}

	av, err := dynamodbattribute.MarshalMap(et)
	if err != nil {
		log.Printf("failed converting et to dynamo attribute map, err = %v", err)
		writeGIF(w)
		return
	}

	upd := make(map[string]*dynamodb.AttributeValue)
	updexp := "SET "
	i := 0
	for k, v := range av {
		// can't include id in the update attributes
		i++
		if k == "id" {
			continue
		}

		upd[":"+k] = v
		updexp = updexp + k + " = :" + k
		if i < len(upd) {
			updexp = updexp + ", "
		}
	}

	upd[":gets"] = &dynamodb.AttributeValue{
		N: aws.String("1"),
	}
	updexp = updexp + " ADD gets :gets"
	log.Printf("upd = %v", upd)
	log.Printf("updexp = %v", updexp)

	// names := map[string]*string{
	// "#id":      aws.String("id"),
	// "#from":    aws.String("from"),
	// "#to":      aws.String("to"),
	// "#subject": aws.String("subject"),
	// "#views":   aws.String("views"),
	// }

	key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: &et.Id,
		},
	}
	// vals := map[string]*dynamodb.AttributeValue{
	// 	":views": {
	// 		N: aws.String("1"),
	// 	},
	// 	":from": {
	// 		S: aws.String(et.From),
	// 	},
	// 	":to": {
	// 		S: aws.String(et.To),
	// 	},
	// 	":subject": {
	// 		S: aws.String(et.Subject),
	// 	},
	// }

	_, err = dyndb.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key:       key,
		ExpressionAttributeValues: upd,
		UpdateExpression:          aws.String(updexp),
	})
	if err != nil {
		log.Printf("failed writing av to dynamo, err = %v", err)
	}

	writeGIF(w)
	return
}

func writeGIF(w http.ResponseWriter) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Content-Type", "image/gif")
	_, err := w.Write(GIF)
	return err
}
