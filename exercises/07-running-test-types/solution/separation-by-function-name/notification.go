package main

import "fmt"

type Notifier interface {
	SendMessage(string) error
}

type EmailNotifier struct {
}

type SMSNotifier struct {
}

func (email EmailNotifier) SendMessage(message string) error {
	fmt.Println("Sending email purchase notification")
	return nil
}

func (sms SMSNotifier) SendMessage(message string) error {
	fmt.Println("Sending SMS purchase notification")
	return nil
}