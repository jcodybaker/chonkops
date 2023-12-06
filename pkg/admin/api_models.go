package admin

type APIStatus struct {
	Relays map[string]APIRelay
}

type APIRelay struct {
	Enabled bool
}
