package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

/**
 * @author Rancho
 * @date 2022/1/11
 */

type FackEmailSender struct {
    mock.Mock
}

func (mock *FackEmailSender) Send(to, subject, body string) error {
    _ = mock.Called(to, subject, body)
    // return args.Error(0)
    return nil
}

func TestCustomerWelcome_Welcome(t *testing.T) {
    emailer := &FackEmailSender{}
    emailer.On("Send", "cooper@gmail.com", "Welcome", "hi, cooper!").Return(nil)

    // welcomer := CreateCustomerWelcome()
    // welcomer.Emailer = emailer
    // container := &Container{}
    // container.SendEmail = emailer
    // welcomer := container.GetCustomerWelcome()
    DefaultContainer.SendEmail = emailer

    err := DefaultContainer.CustomerWelcome.Welcome("cooper", "cooper@gmail.com")
    assert.NoError(t, err)

    emailer.AssertExpectations(t)
}
