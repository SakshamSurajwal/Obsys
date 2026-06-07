package service

import (
	"time"
)

type RequestObj struct {
	Body      chan interface{}
	TypeOfReq string
	Service   string
	Name      string
	From      time.Time
	To        time.Time
	Limit     string
	Offset    string
}

func NewRequest(Body chan interface{}, TypeOfReq string, Service string, Name string, From time.Time, To time.Time, Limit string, Offset string) *RequestObj {
	return &RequestObj{
		Body:      Body,
		Service:   Service,
		TypeOfReq: TypeOfReq,
		Name:      Name,
		From:      From,
		To:        To,
		Limit:     Limit,
		Offset:    Offset,
	}
}
