package unione_test

import (
	"io/ioutil"
	"net/http"
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
	var givenJson string
	client := utils.NewHttpClient(func(req *http.Request) (res *http.Response, err error) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		givenJson = string(b)

		return
	})

	msg := message.NewMessage().
		From("foo@bar.example", "John Doe").
		To(message.NewRecipient("recipient@example.com")).
		Subject("Lorem ipsum").
		BodyHtml("Novus ordo seclorum")

	err := unione.New("foo", "bar").
		Client(client).
		Send(msg)

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"username":"foo","api_key":"bar","message":{"from_name":"John Doe","from_email":"foo@bar.example","recipients":[{"email":"recipient@example.com"}],"subject":"Lorem ipsum","body":{"html":"Novus ordo seclorum"}}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}
