package models

// ApiElement describes the behaviour of an element of the API
// Elements that implement this interface allow for easier pulling of their data
type ApiElement interface {
	JSON() ([]byte, error)
}
