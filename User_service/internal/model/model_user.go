package model

type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Role     string `bson:"role"`
}
