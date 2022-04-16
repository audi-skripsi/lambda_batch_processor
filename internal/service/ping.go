package service

import "time"

func (s *service) Ping() (pingResponse string, timestamp int64) {
	return "pong", time.Now().Unix()
}
