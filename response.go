package unione

type Response struct {
	Status       string   `json:"status"`
	JobID        string   `json:"job_id,omitempty"`
	Emails       []string `json:"emails,omitempty"`
	FailedEmails []string `json:"failed_emails,omitempty"`
	Code         int      `json:"code,omitempty"`
	Message      string   `json:"message,omitempty"`
}
