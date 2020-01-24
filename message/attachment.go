package message

import (
	"encoding/base64"
	"io/ioutil"
	"mime"
	"path/filepath"
)

const defaultMimeType = "application/octet-stream"

// Attachment object.
type Attachment struct {
	// Type content type.
	Type string `json:"type"`
	// Name attachment name.
	Name string `json:"name"`
	// Content base64-encoded file content.
	Content string `json:"content"`
}

// NewAttachment returns Attachment by file name.
func NewAttachment(fileName string, n ...string) (a *Attachment, err error) {
	ext := filepath.Ext(fileName)

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = defaultMimeType
	}

	var name string
	if len(n) > 0 {
		name = n[0]
	} else {
		name = filepath.Base(fileName)
	}

	var b []byte
	if b, err = ioutil.ReadFile(fileName); err != nil {
		return
	}

	a = &Attachment{
		Type:    mimeType,
		Name:    name,
		Content: base64.StdEncoding.EncodeToString(b),
	}

	return
}
