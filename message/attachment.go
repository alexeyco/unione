package message

import (
	"encoding/base64"
	"io/ioutil"
	"mime"
	"path/filepath"
)

const defaultMimeType = "application/octet-stream"

type attachment struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func newAttachment(fileName string, n ...string) (a *attachment, err error) {
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

	a = &attachment{
		Type:    mimeType,
		Name:    name,
		Content: base64.StdEncoding.EncodeToString(b),
	}

	return
}
