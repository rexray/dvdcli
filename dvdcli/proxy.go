package main

import "github.com/docker/docker/pkg/plugins"

type client interface {
	Call(string, interface{}, interface{}) error
}

type volumeDriverProxy struct {
	client
}

var volPlugin *plugins.Plugin
var volProxy volumeDriverProxy
