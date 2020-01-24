package message_test

import (
	"path/filepath"
	"testing"

	"github.com/alexeyco/unione/message"
)

const fileName = "attachment_test.jpg"

func TestNewAttachment(t *testing.T) {
	attachment, err := message.NewAttachment(fileName)
	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err)
	}

	expectedName := filepath.Base(fileName)
	if expectedName != attachment.Name {
		t.Fatalf(`Attachment filename should be "%s", "%s" given`, expectedName, attachment.Name)
	}

	custimFileName := "customFileName"
	attachment, err = message.NewAttachment(fileName, custimFileName)
	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err)
	}

	if custimFileName != attachment.Name {
		t.Fatalf(`Attachment filename should be "%s", "%s" given`, expectedName, attachment.Name)
	}
}

func TestNewAttachment_Error(t *testing.T) {
	attachment, err := message.NewAttachment("wrong/path/to/file.jpg")
	if err == nil {
		t.Fatal(`Error should not be nil`)
	}

	if attachment != nil {
		t.Fatal(`Attachment should be nil`)
	}
}
