package dto

type EventLog struct {
	UID       string      `json:"uid"`
	Level     string      `json:"level"`
	AppName   string      `json:"app_name"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}
