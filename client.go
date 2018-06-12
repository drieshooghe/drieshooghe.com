package main

type client interface {
	tokens() map[string]string
	getInfo()
}
