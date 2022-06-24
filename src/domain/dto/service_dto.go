package dto

import "time"

type ServiceList struct {
	Namespace  string
	Name       string
	Type       string
	Labels     map[string]string
	Selector   map[string]string
	CreateTime time.Time
}
