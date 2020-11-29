package controller

type Controller struct {
	id int
}

func NewController() *Controller {
	var id int = 0
	id++
	return &Controller{id: id}
}

func (c Controller) ShowController() int {
	return c.id
}
