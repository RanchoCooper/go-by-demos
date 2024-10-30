package main

import (
	"fmt"
)

type Server struct {
	Option
}

type Option struct {
	Port int
	Mode string
}

type ServerOption func(*Option)

func WithPort(port int) ServerOption {
	return func(option *Option) {
		option.Port = port
	}
}

func WithMode(mode string) ServerOption {
	return func(option *Option) {
		option.Mode = mode
	}
}

func NewServer(opts ...ServerOption) Server {
	o := &Option{} // handle default value when init Option
	for _, opt := range opts {
		opt(o)
	}

	return Server{Option: *o}
}

func main() {
	server := NewServer(WithPort(8080), WithMode("release"))

	fmt.Printf("server is running on port %d with %s mode", server.Port, server.Mode)
}
