package message_test

import (
	"testing"

	"github.com/alexeyco/unione/message"
	"github.com/alexeyco/unione/utils"
)

func TestRecipient_Name(t *testing.T) {
	recipient := message.NewRecipient("foo@bar.example").
		Name("John Doe")

	givenJson, err := utils.ToJson(recipient)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"email":"foo@bar.example","substitutions":{"to_name":"John Doe"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestRecipient_Substitution(t *testing.T) {
	recipient := message.NewRecipient("foo@bar.example").
		Substitution("foo", "bar")

	givenJson, err := utils.ToJson(recipient)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"email":"foo@bar.example","substitutions":{"foo":"bar"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestRecipient_Meta(t *testing.T) {
	recipient := message.NewRecipient("foo@bar.example").
		Meta("foo", "bar")

	givenJson, err := utils.ToJson(recipient)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"email":"foo@bar.example","metadata":{"foo":"bar"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestNewRecipient(t *testing.T) {
	recipient := message.NewRecipient("foo@bar.example")

	givenJson, err := utils.ToJson(recipient)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"email":"foo@bar.example"}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}
