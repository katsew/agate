package resp

import "time"

// System-wide states like server time, version
type SystemInformation struct {
	Version int `json:"version"`
	Time 	int64 `json:"time"`
}

func NewSystemInfo() SystemInformation {
	return SystemInformation{
		Version: 0,
		Time: time.Now().Unix(),
	}
}