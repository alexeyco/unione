package message

type Recipient interface {
	Name(name string) Recipient
	Substitution(key string, val interface{}) Recipient
	Meta(key string, val interface{}) Recipient
}

type recipient struct {
	Email         string                 `json:"email"`
	Substitutions map[string]interface{} `json:"substitutions,omitempty"`
	MetaData      map[string]interface{} `json:"metadata,omitempty"`
}

func (r *recipient) Name(name string) Recipient {
	return r.Substitution("to_name", name)
}

func (r *recipient) Substitution(key string, val interface{}) Recipient {
	r.Substitutions[key] = val
	return r
}

func (r *recipient) Meta(key string, val interface{}) Recipient {
	r.MetaData[key] = val
	return r
}

func NewRecipient(email string) Recipient {
	return &recipient{
		Email:         email,
		Substitutions: map[string]interface{}{},
		MetaData:      map[string]interface{}{},
	}
}
