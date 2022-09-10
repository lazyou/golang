package controllers

import (
	"app"
	"fmt"
)

type SiteController struct {
	app.Controller
}

func (p SiteController) Index() {
	fmt.Fprint(p.Response, p.Request.RequestURI)
}
