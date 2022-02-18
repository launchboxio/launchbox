package api

type Application struct {
	c *Client
}

func (c *Client) Application() *Application {
	return &Application{c}
}

func (a *Application) List() {

}

func (a *Application) Create() {

}

func (a *Application) Update() {

}

func (a *Application) Delete() {

}

func (a *Application) Find() {

}