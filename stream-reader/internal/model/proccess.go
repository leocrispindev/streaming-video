package model

type Proccess struct {
	Session    string     `json:"session"`
	VideoName  string     `json:"videoName"`
	Connection Connection `json:"connection,omitempty"`
	Id         int64
}

type Connection struct {
}
