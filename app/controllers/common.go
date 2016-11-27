package controllers

import "time"

// ResultJson is a json struct used in controller response
type ResultJson struct {
	Success bool
	Msg     string
	Data    interface{}
}

//PostData model.
type PostData struct {
	Title    string
	Content  string
	Date     time.Time
	Category int
	Tag      string
	Keywords string
	passwd   string
	Type     int
}
