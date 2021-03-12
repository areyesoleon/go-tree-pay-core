package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Res struct {
	Ok      bool        `json:"ok"`
	Res     interface{} `json:"res,omitempty"`
	Message string      `json:"message,omitempty"`
}

// TimeData ...
type TimeData struct {
	Create time.Time `json:"create"`
	Update time.Time `json:"update"`
}

//User ...
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Name     string             `json:"name"`
	LastName string             `json:"lastName"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Admin    bool               `json:"admin"`
	TimeData
}
