package message_test

import (
	"encoding/json"
	"testing"

	"github.com/alexeyco/unione/message"
)

func TestRecipient_Name(t *testing.T) {
	recipient := message.NewRecipient("foo@bar.example").
		Name("John Doe")

	givenJson, err := toJson(recipient)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"email":"foo@bar.example","substitutions":{"to_name":"John Doe"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestRecipient_Substitution(t *testing.T) {
	recipient := message.NewRecipient("foo@bar.example").
		Substitution("foo", "bar")

	givenJson, err := toJson(recipient)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"email":"foo@bar.example","substitutions":{"foo":"bar"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestRecipient_Meta(t *testing.T) {
	recipient := message.NewRecipient("foo@bar.example").
		Meta("foo", "bar")

	givenJson, err := toJson(recipient)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"email":"foo@bar.example","metadata":{"foo":"bar"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestNewRecipient(t *testing.T) {
	recipient := message.NewRecipient("foo@bar.example")

	givenJson, err := toJson(recipient)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"email":"foo@bar.example"}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func toJson(v interface{}) (s string, err error) {
	var b []byte
	if b, err = json.Marshal(v); err != nil {
		return
	}

	s = string(b)

	return
}
