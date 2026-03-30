package lib

import (
	"fmt"
	"net/http"

	"strumati.cloud/bucket/xml"
)

func LS(w http.ResponseWriter) {
	header := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"

	payload := xml.Xml("ListAllMyBucketsResult", nil, []string{

		xml.Xml("Buckets", nil, []string{
			"",
		}),
		xml.Xml("Owner", nil, []string{
			xml.Xml("ID", nil, []string{
				"123-456-789",
			}),
			xml.Xml("DisplayName", nil, []string{
				"example",
			}),
		}),
	})

	fmt.Println(payload)

	fmt.Fprintf(w, "%s", header+payload)
}
