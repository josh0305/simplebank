package token

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different type of error returned by Verify token function
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

var Issuer = "SimpleBank"

type Payload struct {
	UID      uuid.UUID `json:"id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		tokenID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    Issuer,
		},
	}

	return payload, nil
}

func (payload *Payload) Validate() error {
	if payload.ExpiresAt.Before(time.Now()) {
		return ErrExpiredToken
	}

	return nil
}
