package controllers

import "blog/app/models"

var tagModel *models.Tag

// ResultJson is a json struct used in controller response
type ResultJson struct {
	Success bool
	Msg     string
	Data    interface{}
}
