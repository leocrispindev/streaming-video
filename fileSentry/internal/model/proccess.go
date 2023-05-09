package model

type Proccess struct {
	Session    string     `json:"session"`
	VideoName  string     `json:"videoName"`
	Connection Connection `json:"connection,omitempty"`
	Id         int64      `json:"id,omitempty"`
	TopicName  string     `json:"topicName"`
	Segments   int
}

type Connection struct {
}
