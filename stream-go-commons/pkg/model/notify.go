package model_commons

type NotifyProccess struct {
	ProccessID   string
	VideoName    string
	Action       int
	TopicName    string
	VideoId      int
	VideoOptions VideoOptions
}

type VideoOptions struct {
	Width    int     // Width of frames.
	Height   int     // Height of frames.
	Depth    int     // Depth of frames.
	Bitrate  int     // Bitrate for video encoding.
	Frames   int     // Total number of frames.
	Duration float64 // Duration of video in seconds.
	Fps      float64 // Frames per second.
	Codec    string
}
