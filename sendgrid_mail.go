// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package sendgrid

import (
	"encoding/json"
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

type requestBody struct {
	Personalizations []personalization `json:"personalizations"`
	From             emailAdd          `json:"from"`
	Body             []content         `json:"content"`
}

type content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type emailAdd struct {
	Email string `json:"email"`
}

type personalization struct {
	To      []emailAdd `json:"to"`
	Subject string     `json:"subject"`
}

func getPersonlizations(subject string, to []string) []personalization {
	receipients := make([]emailAdd, len(to))
	for index, email := range to {
		receipients[index] = emailAdd{email}
	}
	personalizationInstance := personalization{receipients, subject}
	personalizations := []personalization{personalizationInstance}
	return personalizations
}

func Mail(from string, subject string, to []string, body string) (*rest.Response, error) {
	sender := emailAdd{from}
	contentText := content{"text/plain", body}
	contentBody := []content{contentText}
	personalizations := getPersonlizations(subject, to)
	reqBody := requestBody{personalizations, sender, contentBody}
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body, _ = json.Marshal(reqBody)
	response, err := sendgrid.API(request)
	return response, err
}
