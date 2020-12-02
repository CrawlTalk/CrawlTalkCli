package main

type ErrorsT struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Time   int    `json:"time,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type DataT struct {
	User []UserT `json:"user,omitempty"`
	Flow []FlowT `json:"flow,omitempty"`
}

type FlowT struct {
	ID    int    `json:"id,omitempty"`
	Time  int    `json:"time,omitempty"`
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Info  string `json:"info,omitempty"`
}

type UserT struct {
	Password string `json:"password,omitempty"`
	Login    string `json:"login,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	AuthId   string `json:"auth_id,omitempty"`
	UUID     int64  `json:"uuid,omitempty"`
}

type MoreliaT struct {
	Type   string   `json:"type"`
	Data   DataT    `json:"data,omitempty"`
	Errors *ErrorsT `json:"errors,omitempty"`
}
