package main

import (
    "fmt"
)

/**
 * @author Rancho
 * @date 2022/1/11
 */

// type Container struct {
//     SendEmail       EmailSender
//     CustomerWelcome *CustomerWelcome
// }
//
// func (c *Container) GetSendEmail() EmailSender {
//     if c.SendEmail == nil {
//         c.SendEmail = CreateSendEmail()
//     }
//
//     return c.SendEmail
// }
//
// func (c *Container) GetCustomerWelcome() *CustomerWelcome {
//     if c.CustomerWelcome == nil {
//         c.CustomerWelcome = CreateCustomerWelcome()
//     }
//
//     return c.CustomerWelcome
// }

type SendEmail struct {
    From string
}

func CreateSendEmail() *SendEmail {
    return &SendEmail{
        From: "ranchocooper@gmail.com",
    }
}

func (sender *SendEmail) Send(to, subject, body string) error {
    fmt.Println("I am the original sender")
    return nil
}

type EmailSender interface {
    Send(to, subject, body string) error
}

type CustomerWelcome struct {
    Emailer EmailSender
}

func CreateCustomerWelcome() *CustomerWelcome {
    return &CustomerWelcome{
        Emailer: CreateSendEmail(),
    }
}

func (w *CustomerWelcome) Welcome(name, email string) error {
    subject := "Welcome"
    body := fmt.Sprintf("hi, %s!", name)

    return w.Emailer.Send(email, subject, body)
}
