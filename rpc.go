package plugin

type rpc struct {
	srv *Plugin
}

// Payload for the message rpc call
type Payload struct {
	Message string `json:"message"`
}

// Message is an exposed RPC endpoint `plugin.Message`
func (s *rpc) Message(input Payload, output *Payload) error {
	*output = input

	return nil
}
