package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     string
	Name   string
	Author string
}
