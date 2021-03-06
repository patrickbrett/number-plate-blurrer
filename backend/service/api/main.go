package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

// ContentType contains the Content-Type header sent on all responses.
const ContentType = "application/json; charset=utf8"

// MessageResponse models a simple message responses.
type MessageResponse struct {
	Message string `json:"message"`
}

// WelcomeMessageResponse is the response returned by the /message endpoint.
var WelcomeMessageResponse = MessageResponse{"Test message!"}

// MessageHandler is a http.HandlerFunc for the /message endpoint.
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(WelcomeMessageResponse)
}

// ImagesHandler is a http.HandlerFunc for the /images endpoint.
func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		objects := listImages()
		json.NewEncoder(w).Encode(objects)
	} else if r.Method == http.MethodPut {
		signed := putImage()
		json.NewEncoder(w).Encode(signed)
	}
}

// LabelsHandler is a http.HandlerFunc for the /labels endpoint.
func LabelsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		objects := getLabels(r)
		json.NewEncoder(w).Encode(objects)
	}
}

// BlurHandler is a http.HandlerFunc for the /blur endpoint.
func BlurHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		result := blurImage(r)
		json.NewEncoder(w).Encode(result)
	}
}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	http.Handle("/message", h(MessageHandler))
	http.Handle("/images", h(ImagesHandler))
	http.Handle("/labels", h(LabelsHandler))
	http.Handle("/blur", h(BlurHandler))
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		next.ServeHTTP(w, r)
	})
}

func main() {
	RegisterRoutes()
	log.Fatal(gateway.ListenAndServe(":3000", nil))
}
