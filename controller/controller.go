package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/studio-b12/gowebdav"
)

var (
	root     = ""
	user     = ""
	password = ""
	webdav   = 
)

func File(c *fiber.Ctx) error {

	directory := c.Query("directory")

	var webdavFilePath = directory

	bytes, _ := webdav.Read(webdavFilePath)

	return c.Send(bytes)
}
