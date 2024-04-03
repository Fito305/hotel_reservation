package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"` // These are called Go / JSON tags. Will convert the left to the right. omitempty will omit the id if it's empty. Or use _ to omit.
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
