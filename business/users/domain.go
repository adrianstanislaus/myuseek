package users

import (
	"context"
	"time"
)

type Domain struct {
	Id                int
	FirstName         string
	LastName          string
	Username          string
	Email             string
	Password          string
	Bio               string
	Profile_pic       string
	Subscription_type string
	Token             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Usecase interface {
	Register(ctx context.Context, firstname, lastname, username, email, password, bio, profile_pic, subscription_type string) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, firstname, lastname, username, email, password, bio, profile_pic, subscription_type string) (Domain, error)
}