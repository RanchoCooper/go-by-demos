package service

import (
	"context"
	"io"
	"log"

	"github.com/google/wire"
)

// Sender interface for sending email
type Sender interface {
	Send()
}

// MailConfig configuration of mail
type MailConfig struct {

}

// MailSender implement the Sender interface
type MailSender struct {

}

func (e *MailSender) Send() {
	log.Println("send email.")
}

// NewMailSender initialize MailSender with config (provider)
func NewMailSender(mc *MailConfig) *MailSender {
	return &MailSender{}
}

// declare that NewMailSender will return
// puzzle
var MailSet = wire.NewSet(NewMailSender, wire.Bind(new(Sender), new(*MailSender)))

type Options struct {
	// the set of recommended greetings
	Messages []string
	// the location to send greetings. goes to stdout if nil
	Writer io.Writer
}

type Greeter struct {
	msg []string
	w io.Writer
}

func NewGreeter(ctx context.Context, opts *Options) (*Greeter, error) {
	return &Greeter{
		msg: opts.Messages,
		w:   opts.Writer,
	}, nil
}

// puzzle
var GreeterSet = wire.NewSet(NewGreeter, wire.Struct(new(Options), "*"))

