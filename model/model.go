package model



var	Employees []Employee


type Employee struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	MailId string `json:"mailid,omitempty"`
}