package unione_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/alexeyco/unione"
	"github.com/alexeyco/unione/message"
	"github.com/alexeyco/unione/utils"
)

func TestClient_LanguageEn(t *testing.T) {

}

func TestClient_LanguageRu(t *testing.T) {

}

func TestClient_Send(t *testing.T) {
	responseMap := map[string]interface{}{
		"status": "foo",
		"job_id": "bar",
		"emails": []string{
			"success@foo.test",
			"success@bar.test",
		},
		"failed_emails": map[string]string{
			"failed@foo.test": "foo",
			"failed@bar.test": "bar",
		},
	}

	expectedResponseJson, _ := utils.ToJson(responseMap)

	var givenReguestJson string
	client := utils.NewTestHttpClient(func(req *http.Request) (res *http.Response, err error) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		givenReguestJson = string(b)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(expectedResponseJson)),
		}

		return
	})

	msg := message.NewMessage().
		From("foo@bar.example", "John Doe").
		To(message.NewRecipient("recipient@example.com")).
		Subject("Lorem ipsum").
		BodyHtml("Novus ordo seclorum")

	success, failed, err := unione.New("foo", "bar").
		Client(client).
		Send(msg)

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"username":"foo","api_key":"bar","message":{"from_name":"John Doe","from_email":"foo@bar.example","recipients":[{"email":"recipient@example.com"}],"subject":"Lorem ipsum","body":{"html":"Novus ordo seclorum"}}}`
	utils.JsonIsEqual(t, expectedJson, givenReguestJson)

	if !reflect.DeepEqual(success, responseMap["emails"]) {
		t.Fatal(`Success emails should be equal`)
	}

	failedEmails := map[string]error{}
	for email, err := range responseMap["failed_emails"].(map[string]string) {
		failedEmails[email] = errors.New(err)
	}

	if !reflect.DeepEqual(failed, failedEmails) {
		t.Fatal(`Failed emails should be equal`)
	}
}

func ExampleClient_Send() {
	recipient := message.NewRecipient("recipient@example.com").
		Name("John Doe")

	msg := message.NewMessage().
		From("site@example.com", "My site").
		To(recipient).
		Subject("Awesome news, buddy").
		BodyPlainText("Return to my site and enjoy")

	client := unione.New("username", "api-key")

	_, _, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}
}
