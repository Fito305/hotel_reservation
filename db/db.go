// dg.go is for general Database things
package db

// import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	DBNAME = "hotel-reservation"
	TestDBNAME = "hotel-reservation-test"
	DBURI  = "mongodb://localhost:27017" // endpoint of our database
)


type Store struct {
	User UserStore
	Hotel HotelStore
	Room RoomStore
	Booking BookingStore
}

// Don't need this function

// func ToObjectID(id string) (primitive.ObjectID) {
//     oid, err := primitive.ObjectIDFromHex(id) // oid is objectID
//     if err != nil {
//         panic(err)
//     }
//     return oid
//
// }
