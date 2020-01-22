package message

import "encoding/json"

type Message interface {
	Header(key, val string) Message
	From(email string, name ...string) Message
	ReplyTo(email string) Message
	To(recipients ...Recipient) Message
	Subject(subject string) Message
	BodyHtml(html string) Message
	BodyPlainText(plainText string) Message
	Substitution(key string, val interface{}) Message
	Meta(key string, val interface{}) Message
	TrackLinks() Message
	TrackRead() Message
	Option(key string, val interface{}) Message
	UnsubscribeUrl(u string) Message
	Json() (s string, err error)
}

type body struct {
	Html      string `json:"html,omitempty"`
	PlainText string `json:"plaintext,omitempty"`
}

type message struct {
	Headers           map[string]string      `json:"headers,omitempty"`
	FromEmail         string                 `json:"from_email,omitempty"`
	FromName          string                 `json:"from_name,omitempty"`
	ReplyToEmail      string                 `json:"reply_to,omitempty"`
	Recipients        []Recipient            `json:"recipients,omitempty"`
	SubjectText       string                 `json:"subject,omitempty"`
	Body              *body                  `json:"body,omitempty"`
	Substitutions     map[string]interface{} `json:"global_substitutions,omitempty"`
	MetaData          map[string]interface{} `json:"metadata,omitempty"`
	TrackLinksEnabled int                    `json:"track_links,omitempty"`
	TrackReadEnabled  int                    `json:"track_read,omitempty"`
	Options           map[string]interface{} `json:"options,omitempty"`
}

func (m *message) Header(key, val string) Message {
	m.Headers[key] = val
	return m
}

func (m *message) From(email string, name ...string) Message {
	m.FromEmail = email
	if len(name) > 0 {
		m.FromName = name[0]
	}

	return m
}

func (m *message) ReplyTo(email string) Message {
	m.ReplyToEmail = email
	return m
}

func (m *message) To(recipients ...Recipient) Message {
	m.Recipients = append(m.Recipients, recipients...)
	return m
}

func (m *message) Subject(subject string) Message {
	m.SubjectText = subject
	return m
}

func (m *message) BodyHtml(html string) Message {
	if m.Body == nil {
		m.Body = &body{}
	}

	m.Body.Html = html

	return m
}

func (m *message) BodyPlainText(plainText string) Message {
	if m.Body == nil {
		m.Body = &body{}
	}

	m.Body.PlainText = plainText

	return m
}

func (m *message) Substitution(key string, val interface{}) Message {
	m.Substitutions[key] = val
	return m
}

func (m *message) Meta(key string, val interface{}) Message {
	m.MetaData[key] = val
	return m
}

func (m *message) TrackLinks() Message {
	m.TrackLinksEnabled = 1
	return m
}

func (m *message) TrackRead() Message {
	m.TrackReadEnabled = 1
	return m
}

func (m *message) Option(key string, val interface{}) Message {
	m.Options[key] = val
	return m
}

func (m *message) UnsubscribeUrl(u string) Message {
	return m.Option("unsubscribe_url", u)
}

func (m *message) Json() (s string, err error) {
	var b []byte
	if b, err = json.Marshal(m); err != nil {
		return
	}

	s = string(b)

	return
}

func NewMessage() Message {
	return &message{
		Headers:       map[string]string{},
		Substitutions: map[string]interface{}{},
		MetaData:      map[string]interface{}{},
		Options:       map[string]interface{}{},
	}
}
