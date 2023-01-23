package config

type EmailBody struct {
	To    string `json:"to"`
	Body  string `json:"Body"`
	Title string `json:"Title"`
}

type SmsBody struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
