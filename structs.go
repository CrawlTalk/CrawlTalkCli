package main

type ErrorsT struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Time   int    `json:"time,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type DataT struct {
	Time    int        `json:"time,omitempty"`
	User    []UserT    `json:"user,omitempty"`
	Flow    []FlowT    `json:"flow,omitempty"`
	Message []MessageT `json:"message,omitempty"`
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

type MessageT struct {
	ID           int    `json:"id,omitempty"`
	Text         string `json:"text,omitempty"`
	FromUserUUID int    `json:"from_user_uuid,omitempty"`
	Time         int    `json:"time,omitempty"`
	FromFlowID   int    `json:"from_flow_id,omitempty"`
	FilePicture  string `json:"file_picture,omitempty"`
	FileVideo    string `json:"file_video,omitempty"`
	FileAudio    string `json:"file_audio,omitempty"`
	FileDocument string `json:"file_document,omitempty"`
	Emoji        string `json:"emoji,omitempty"`
	EditedTime   int    `json:"edited_time,omitempty"`
	EditedStatus bool   `json:"edited_status,omitempty"`
}
