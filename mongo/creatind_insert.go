package mongoDb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatingInsert(serviceName string, data []byte, m *mongo.Collection) error {
	filter := bson.D{
		{"service", bson.M{"$eq": serviceName}},
	}
	cursor, err := m.Find(context.Background(), filter)
	if err != nil {
		return err
	}
	tmp := make(map[string]any)
	err = cursor.Decode(&tmp)
	if err != nil {
		return err
	}
	return nil
}
