package unione

// Response
type Response struct {
	// Status specifies if the request was successful.
	Status string `json:"status"`

	// JobID unique identifier of the executed sending.
	JobID string `json:"job_id,omitempty"`

	// Emails array, contains the email addresses where the sending was successful.
	Emails []string `json:"emails,omitempty"`

	// FailedEmails object, contains emails to which the sending for some reasons has not been carried out.
	// The object is filled in format: “address” : “state”
	FailedEmails map[string]string `json:"failed_emails,omitempty"`

	// Code API error code.
	// See: https://one.unisender.com/en/docs/page/Error_Codes
	Code int `json:"code,omitempty"`

	// Message API error message.
	Message string `json:"message,omitempty"`
}
