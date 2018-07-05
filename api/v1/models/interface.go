package models

// APIElement describes the behaviour of an element of the API
// Elements that implement this interface allow for easier pulling of their data
type APIElement interface {
	JSON() ([]byte, error)
}
