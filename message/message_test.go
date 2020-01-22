package message_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alexeyco/unione/message"
)

func TestMessage_Header(t *testing.T) {
	msg := message.NewMessage().
		Header("foo", "bar")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"headers":{"foo":"bar"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_From(t *testing.T) {
	msg := message.NewMessage().
		From("foo@bar.example", "John Doe")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"from_name":"John Doe","from_email":"foo@bar.example"}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_ReplyTo(t *testing.T) {
	msg := message.NewMessage().
		ReplyTo("foo@bar.example")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"reply_to":"foo@bar.example"}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_To(t *testing.T) {
	msg := message.NewMessage().
		To(message.NewRecipient("foo@bar.example"))

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"recipients":[{"email":"foo@bar.example"}]}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Subject(t *testing.T) {
	msg := message.NewMessage().
		Subject("Lorem ipsum dolor")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"subject":"Lorem ipsum dolor"}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_BodyHtml(t *testing.T) {
	msg := message.NewMessage().
		BodyHtml("Lorem ipsum dolor")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"body":{"html":"Lorem ipsum dolor"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_BodyPlainText(t *testing.T) {
	msg := message.NewMessage().
		BodyPlainText("Lorem ipsum dolor")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"body":{"plaintext":"Lorem ipsum dolor"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Substitution(t *testing.T) {
	msg := message.NewMessage().
		Substitution("foo", "bar")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"global_substitutions":{"foo":"bar"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Meta(t *testing.T) {
	msg := message.NewMessage().
		Meta("foo", "bar")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"metadata":{"foo":"bar"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_TrackLinks(t *testing.T) {
	msg := message.NewMessage().
		TrackLinks()

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"track_links":1}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_TrackRead(t *testing.T) {
	msg := message.NewMessage().
		TrackRead()

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"track_read":1}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Option(t *testing.T) {
	msg := message.NewMessage().
		Option("foo", "bar")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"options":{"foo":"bar"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_UnsubscribeUrl(t *testing.T) {
	msg := message.NewMessage().
		UnsubscribeUrl("https://foo.bar/baz")

	givenJson, err := msg.Json()
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"options":{"unsubscribe_url":"https://foo.bar/baz"}}`
	jsonIsEqual(t, expectedJson, givenJson)
}

func jsonToMap(s string) (m map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(s), &m)
	return
}

func jsonIsEqual(t *testing.T, expectedJson, givenJson string) {
	var expectedMap map[string]interface{}
	var err error
	if expectedMap, err = jsonToMap(expectedJson); err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
		return
	}

	var givenMap map[string]interface{}
	if givenMap, err = jsonToMap(givenJson); err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
		return
	}

	if !reflect.DeepEqual(expectedMap, givenMap) {
		t.Errorf(`JSON should be %s, %s given`, expectedJson, givenJson)
	}
}
