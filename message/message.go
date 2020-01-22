package message

// Message email message interface.
type Message interface {
	// Header sets custom email header.
	Header(key, val string) Message

	// From sets sender's contact.
	From(email string, name ...string) Message

	// ReplyTo sets email Reply-To header.
	ReplyTo(email string) Message

	// To sets email recipients.
	To(recipients ...Recipient) Message

	// Subject sets email subject.
	Subject(subject string) Message

	// BodyHtml sets email html body part.
	BodyHtml(html string) Message

	// BodyPlainText sets email plain text body part.
	BodyPlainText(plainText string) Message

	// Substitution passes the substitution to user.
	Substitution(key string, val interface{}) Message

	// Meta sets metadata value. Used for webhooks.
	//
	// See: https://one.unisender.com/en/docs/page/Webhook.
	Meta(key string, val interface{}) Message

	// TrackLinks enables email links tracking.
	TrackLinks() Message

	// TrackRead enables email read tracking.
	TrackRead() Message

	// Option sets custom email option.
	Option(key string, val interface{}) Message

	// UnsubscribeUrl sets custom unsubscribe link.
	//
	// See: https://one.unisender.com/en/docs/page/send#unsub
	UnsubscribeUrl(u string) Message
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

// NewMessage returns new message object.
func NewMessage() Message {
	return &message{
		Headers:       map[string]string{},
		Substitutions: map[string]interface{}{},
		MetaData:      map[string]interface{}{},
		Options:       map[string]interface{}{},
	}
}
