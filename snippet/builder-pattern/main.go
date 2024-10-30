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

type ServerBuilder struct {
	option Option
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (sb *ServerBuilder) WithPort(port int) *ServerBuilder {
	sb.option.Port = port
	return sb
}

func (sb *ServerBuilder) WithMode(mode string) *ServerBuilder {
	sb.option.Mode = mode
	return sb
}

func (sb *ServerBuilder) Build() *Server {
	return &Server{Option: sb.option}
}

func main() {
	builder := NewServerBuilder()

	server := builder.WithPort(8080).WithMode("release").Build()

	fmt.Printf("server is running on port %d with %s mode", server.Port, server.Mode)
}
