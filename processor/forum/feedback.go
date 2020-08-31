// Package forum is where the Transformation logic should exist for any Forum related posts.
package forum

import (
	"encoding/json"
	"feedback-collector/myob"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ProcessFeedback will implement the FeedbackProcessor interface
func ProcessFeedback(req *http.Request) {
	rb, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	var f myob.Feedback
	err = json.Unmarshal(rb, &f)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Received: ", f.Source)
}
