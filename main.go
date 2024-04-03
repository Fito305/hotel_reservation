package main

import(
    "log"
    "flag"

    "github.com/Fito305/hotel-reservation/api"
    "github.com/gofiber/fiber/v2"
)

func main() {
    listenAddr := flag.String("listenAddr", ":4000", "THe listen address of the API server")
    flag.Parse()
    app := fiber.New()
    apiv1 := app.Group("/api/v1")

    apiv1.Get("/user", api.HandleGetUsers)
    apiv1.Get("/user:id", api.HandleGetUser)
    log.Fatal(app.Listen(*listenAddr))
}


