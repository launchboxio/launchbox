package api

type Build struct {
	c *Client
}

func (c *Client) Build() *Build {
	return &Build{c}
}

func (b *Build) Create() {

}

func (b *Build) Cancel() {

}

func (b *Build) Get() {
	
}