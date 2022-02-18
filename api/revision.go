package api

type Revision struct {
	c *Client
}

func (c *Client) Revision() *Revision {
	return &Revision{c}
}

func (a *Revision) List() {

}

func (a *Revision) Create() {

}

func (a *Revision) Update() {

}

func (a *Revision) Delete() {

}

func (a *Revision) Find() {

}
