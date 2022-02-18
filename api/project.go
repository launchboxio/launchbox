package api

type Project struct {
	c *Client
}

func (c *Client) Project() *Project {
	return &Project{c}
}

func (a *Project) List() {

}

func (a *Project) Create() {

}

func (a *Project) Update() {

}

func (a *Project) Delete() {

}

func (a *Project) Find() {

}
