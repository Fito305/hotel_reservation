package main

import(
    "context"
    "log"
    "flag"
    
    "github.com/Fito305/hotel-reservation/db"
    // "github.com/Fito305/hotel-reservation/types" // not used
    "github.com/Fito305/hotel-reservation/api"
	"github.com/Fito305/hotel-reservation/api/middleware"


    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/bson" // not used
)

var config = fiber.Config{
    //Override default error handler
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        return c.JSON(map[string]string{"error": err.Error()})
    },
}

func main() {
    listenAddr := flag.String("listenAddr", ":4000", "THe listen address of the API server")
    flag.Parse()
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
    if err != nil {
        log.Fatal(err)
    }

    // handler initialization
	var (
   	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
	store = &db.Store{
		Hotel: hotelStore,
		Room: roomStore,
		User: userStore,
	}
    userHandler = api.NewUserHandler(userStore)
	hotelHandler = api.NewHotelHandler(store)
	authHandler = api.NewAuthHandler(userStore)
	roomHandler = api.NewRoomHandler(store)
	app = fiber.New(config)
	auth = app.Group("/api")
    apiv1 = app.Group("/api/v1", middleware.JWTAuthentication(userStore)) // *** We added the decorator so we have to call the function, pass in the userStore and it'll return the handler we need.
	)

	// auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// user handlers
    apiv1.Put("user/:id", userHandler.HandlePutUser)
    apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)  // You need to have the colon : before the id.
    apiv1.Post("/user", userHandler.HandlePostUser)
    apiv1.Get("/user", userHandler.HandleGetUsers)
    apiv1.Get("/user/:id", userHandler.HandleGetUser) //Bug** Fixed** I was missing the / in /user:id -> /user/:id
    	
	// hotel handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	apiv1.Post("/room/:id/book", roomHandler.HandleBookRoom)

	app.Listen(*listenAddr)
}


