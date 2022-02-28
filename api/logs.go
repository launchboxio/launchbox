package api

type Logs struct {
	c *Client
}

func (c *Client) Logs() *Logs {
	return &Logs{c}
}

func (l *Logs) Stream() error {
	return nil
}
