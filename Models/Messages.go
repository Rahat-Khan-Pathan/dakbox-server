package Models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Messages struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MessageTo string             `json:"messageto"`
	Message   string             `json:"message"`
	DateTime  primitive.DateTime `json:"datetime"`
}
type MessagesSearch struct {
	Name       string `json:"name"`
	NameIsUsed bool   `json:"nameisused"`
}

func (obj Messages) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.MessageTo, validation.Required),
		validation.Field(&obj.Message, validation.Required),
	)
}
func (obj MessagesSearch) GetBranchSearchBSONObj() bson.M {
	self := bson.M{}

	if obj.NameIsUsed {
		regexPattern := fmt.Sprintf(".*%s.*", obj.Name)
		self["messageto"] = bson.D{{"$regex", primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}

	return self
}
