package arcpaste

import (
	"context"
	"github.com/arcoz0308/arcoz0308.tech/utils/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasteStruct struct {
	ID       primitive.ObjectID `bson:"id"`
	Key      string             `bson:"key"`
	Raw      string             `bson:"raw"`
	Language string             `bson:"language"`
	Expire   int                `bson:"expire"`
	Password string             `bson:"password"`
}

func Data(key string) (PasteStruct, error) {
	var r PasteStruct
	err := database.Paste.FindOne(context.TODO(), bson.D{{"key", key}}).Decode(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}
