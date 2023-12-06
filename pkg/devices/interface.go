package devices

import "context"

type Status struct {
	Enabled bool
}

type Device interface {
	GetName() string
	GetID() string
	GetStatus() Status
}

type Discoverer interface {
	// Discover
	Discover(ctx context.Context) ([]Device, error)
}
