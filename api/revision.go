package api

type Revisions struct {
	c *Client
}

func (c *Client) Revisions() *Revisions {
	return &Revisions{c}
}

func (a *Revisions) List() {

}

func (a *Revisions) Create() {

}

func (a *Revisions) Update() {

}

func (a *Revisions) Delete() {

}

func (a *Revisions) Find() {

}
