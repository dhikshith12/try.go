package main

import "time"


type authentication struct {
	username string
	password string
}

func (a authentication) autherizationString() string {
	return "Basic " + a.username + ":" + a.password + " " + time.Now().String()
}