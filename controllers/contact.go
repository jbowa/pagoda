package controllers

import (
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type Contact struct {
	Controller
}

func (a *Contact) Get(c echo.Context) error {
	p := NewPage(c)
	p.Layout = "main"
	p.Name = "contact"
	p.Data = "This is the contact page"
	return a.RenderPage(c, p)
}

func (a *Contact) Post(c echo.Context) error {
	msg.Set(c, msg.Success, "Thank you for contacting us!")
	msg.Set(c, msg.Info, "We will respond to you shortly.")
	return a.Redirect(c, "home")
}