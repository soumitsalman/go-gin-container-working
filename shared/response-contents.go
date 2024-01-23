package shared

import "time"

func GetPong() map[string]string {
	return map[string]string{"message": "PONG IT"}
}

type MediaContent struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func GetContents() []MediaContent {
	current_time := time.Now().Format("2006-01-02 15:04:05")
	return []MediaContent{
		{ID: 1, Name: "name1", Version: current_time},
		{ID: 2, Name: "name2", Version: current_time},
	}
}
