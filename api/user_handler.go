package api

import (

    "github.com/Fito305/hotel-reservation/types"
    "github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
    u := types.User{
        FirstName: "Felipe",
        LastName: "Acosta",
    }
    return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
    return c.JSON("Felipe")
}
