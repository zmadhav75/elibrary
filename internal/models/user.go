package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Subscribed   bool      `json:"subscribed"`
	SubscribedAt time.Time `json:"subscribed_at"`
	LastNotified time.Time `json:"last_notified"`
}

type SubscriptionRequest struct {
	Email string `json:"email"`
}
