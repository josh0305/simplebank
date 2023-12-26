package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duraiton time.Duration) (string, *Payload, error)

	// VerifyToken verify if token is valid or not
	VerifyToken(token string) (*Payload, error)
}
