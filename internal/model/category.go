package model

import "github.com/nilsyadv/common-service-evnt/pkg/model"

type Category struct {
	model.Trandb
	CategoryName string `json:"category_name"`
}
