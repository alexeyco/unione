package message_test

import (
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/alexeyco/unione/message"
	"github.com/alexeyco/unione/utils"
)

func TestMessage_Header(t *testing.T) {
	msg := message.NewMessage().
		Header("foo", "bar")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"headers":{"foo":"bar"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_From(t *testing.T) {
	msg := message.NewMessage().
		From("foo@bar.example", "John Doe")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"from_name":"John Doe","from_email":"foo@bar.example"}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_ReplyTo(t *testing.T) {
	msg := message.NewMessage().
		ReplyTo("foo@bar.example")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"reply_to":"foo@bar.example"}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_To(t *testing.T) {
	msg := message.NewMessage().
		To(message.NewRecipient("foo@bar.example"))

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"recipients":[{"email":"foo@bar.example"}]}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Subject(t *testing.T) {
	msg := message.NewMessage().
		Subject("Lorem ipsum dolor")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"subject":"Lorem ipsum dolor"}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_BodyHtml(t *testing.T) {
	msg := message.NewMessage().
		BodyHtml("Lorem ipsum dolor")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"body":{"html":"Lorem ipsum dolor"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_BodyPlainText(t *testing.T) {
	msg := message.NewMessage().
		BodyPlainText("Lorem ipsum dolor")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"body":{"plaintext":"Lorem ipsum dolor"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Attach(t *testing.T) {
	msg := message.NewMessage()
	if err := msg.Attach(attachmentFileName); err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	content := attachmentContent(t)

	expectedJson := `{"attachments":[{"type":"image/jpeg","name":"attachment_test.jpg","content":"` + content + `"}]}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_InlineAttach(t *testing.T) {
	msg := message.NewMessage()
	if err := msg.InlineAttach(attachmentFileName, "foobar"); err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	content := attachmentContent(t)

	expectedJson := `{"attachments":[{"type":"image/jpeg","name":"foobar","content":"` + content + `"}]}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Substitution(t *testing.T) {
	msg := message.NewMessage().
		Substitution("foo", "bar")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"global_substitutions":{"foo":"bar"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Meta(t *testing.T) {
	msg := message.NewMessage().
		Meta("foo", "bar")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"metadata":{"foo":"bar"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_TrackLinks(t *testing.T) {
	msg := message.NewMessage().
		TrackLinks()

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"track_links":1}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_TrackRead(t *testing.T) {
	msg := message.NewMessage().
		TrackRead()

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"track_read":1}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_Option(t *testing.T) {
	msg := message.NewMessage().
		Option("foo", "bar")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"options":{"foo":"bar"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func TestMessage_UnsubscribeUrl(t *testing.T) {
	msg := message.NewMessage().
		UnsubscribeUrl("https://foo.bar/baz")

	givenJson, err := utils.ToJson(msg)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	expectedJson := `{"options":{"unsubscribe_url":"https://foo.bar/baz"}}`
	utils.JsonIsEqual(t, expectedJson, givenJson)
}

func attachmentContent(t *testing.T) string {
	var b []byte
	var err error
	if b, err = ioutil.ReadFile(attachmentFileName); err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err)
	}

	return base64.StdEncoding.EncodeToString(b)
}
