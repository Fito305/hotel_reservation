// The idiom is that we place context (ctx) as the first argument in our functions.
// We have access to ToObjectID because it's in the same db folder/ module.
package db

import (
    "context"
    "github.com/Fito305/hotel-reservation/types"
    

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)


const userColl = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper

    GetUserByEmail(context.Context, string) (*types.User, error)
    GetUserByID(context.Context, string) (*types.User, error)
    GetUsers(context.Context) ([]*types.User, error)
    InsertUser(context.Context, *types.User) (*types.User, error)
    DeleteUser(context.Context, string) error
    UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error // You can name the parameters which are named ctx, filter and update here.
}

// The implementation
type MongoUserStore struct {
    client *mongo.Client
    coll  *mongo.Collection
}

// Constructor. Named with New...
func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
    return &MongoUserStore{
        client: client,
        coll: client.Database(DBNAME).Collection(userColl),
    }
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	println("--- Dropping")
	return s.coll.Drop(ctx)
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {
	update := bson.M{"$set": params}
    _, err := s.coll.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    return nil
}
    
func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    // TODO: Maybe its a good idea to handle if we did not delete any user.
    // maybe log it or something?
    _, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
    if err != nil {
        return err
    }
    return nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
    res, err := s.coll.InsertOne(ctx, user)
    if err != nil {
        return nil, err
    }
    user.ID = res.InsertedID.(primitive.ObjectID) // The user passed via the parameter for this func has no ID. So we assign res ID and cast it to a primitive Object.
    // We had to change User struct ID: string to ID: primitive.ObjectID type 
    return user, nil // What is happening is that we are going to manipulate the pointer and then retrun the new pointer 
}


func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
    cur, err := s.coll.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    var users []*types.User
    if err := cur.All(ctx, &users); err != nil {
        return nil, err
    }
    return users, err
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
    // validate the correctness of the ID
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    
    var user types.User // Instantiate a new user
    if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil { 
        // bson.M's id has an underscore becuase thats what we have in the users dir file. In decode() we need to specify the user as a reference/pointer.
        // the bson.M{"_id": `id`} won't work because it's a string. I needs to be converted into a mongo object.(which is like a string but not compatible with Go strings).
        // The ToObjectID() in db/db.go does the convertion. Changed the implementstion to be the oid variable above.
        return nil, err
    }
    return &user, nil // We pass the user by reference &user
}

func (s *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
    var user types.User // Instantiate a new user
    if err := s.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil { 
        return nil, err
    }
    return &user, nil // We pass the user by reference &user
}
