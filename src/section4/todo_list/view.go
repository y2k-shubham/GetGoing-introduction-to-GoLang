package main

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}
