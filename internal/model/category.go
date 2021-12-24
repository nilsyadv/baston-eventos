package model

import "github.com/Nilesh-Coherent/common-service-evnt/pkg/model"

type Category struct {
	model.Trandb
	CategoryName string `json:"category_name"`
}
