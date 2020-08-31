package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Server set up in 2 lines. How easy is that?!
	registerHandlers()
	http.ListenAndServe(":8080", nil)
}

func registerHandlers() {
	http.HandleFunc("/forum", testHandler)
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	var b contract
	defer req.Body.Close()
	err := json.Unmarshal(getJSONBody(req), &b)
	if err != nil {
		log.Println(err)
		var errs []error
		errs = append(errs, err)
		sendFailResponse(w, b, errs)
	}
	fails := meetsContract(&b)
	if len(fails) > 0 {
		sendFailResponse(w, b, fails)
	}
	sendSuccessResponse(w, b)
}

func getJSONBody(req *http.Request) []byte {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}

// TODO: Remove this and use the real struct, this is gross and you should feel ashamed.
type contract struct {
	Source          string            `json:"source"`
	FeedbackContent []contractContent `json:"feedback-content"`
	Size            int               `json:"size"`
}

// TODO: Remove this and use the real struct, this is gross and you should feel ashamed.
type contractContent struct {
	Type          string `json:"type"`
	Author        string `json:"author"`
	ForumUsername string `json:"forumUsername"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	ExternalLink  string `json:"external-link"`
}

// Playing with request validation, will have to refine rules later.
// TODO: Split validation into `controller` package
func meetsContract(c *contract) []error {
	var nonos []error
	if c.Source == "" {
		no := errors.New("ERROR: Field `Source` cannot be blank")
		nonos = append(nonos, no)
	}
	if c.Size <= 0 {
		no := errors.New("ERROR: Must receive payload with greater than 0 items")
		nonos = append(nonos, no)
	}
	for _, e := range c.FeedbackContent {
		if e.Type == "" {
			no := errors.New("ERROR: Field `Type` cannot be blank")
			nonos = append(nonos, no)
			continue
		}
		if e.Author == "" {
			no := errors.New("ERROR: Field `Author` cannot be blank")
			nonos = append(nonos, no)
			continue
		}
		if e.ForumUsername == "" {
			no := errors.New("ERROR: Field `ForumUsername` cannot be blank")
			nonos = append(nonos, no)
			continue
		}
		if e.Title == "" {
			no := errors.New("ERROR: Field `Title` cannot be blank")
			nonos = append(nonos, no)
			continue
		}
		if e.Body == "" {
			no := errors.New("ERROR: Field `Body` cannot be blank")
			nonos = append(nonos, no)
			continue
		}
		if e.ExternalLink == "" {
			no := errors.New("Error: Field `ExternalLink` cannot be blank")
			nonos = append(nonos, no)
			continue
		}
	}
	if len(nonos) > 0 {
		return nonos
	}
	return nil
}

type failedResponse struct {
	Errors              []error
	OriginalRequestBody contract
}

func sendFailResponse(w http.ResponseWriter, b contract, errs []error) {
	var fr failedResponse
	fr.Errors = errs
	fr.OriginalRequestBody = b
	f, err := json.Marshal(fr)
	if err != nil {
		log.Println("ERROR: Cannot init response. Reason: ", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(f)
}

func sendSuccessResponse(w http.ResponseWriter, b contract) {

}
