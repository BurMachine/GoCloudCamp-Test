package mongoDb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataStruct struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ServiceName string             `bson:"name"`
	Config      []interface{}      `bson:"configs"`
}

func CreatingInsert(serviceName string, data []byte, m *mongo.Collection) error {
	filter := bson.D{
		{"service", bson.M{"$eq": serviceName}},
	}
	cursor := m.FindOne(context.Background(), filter)

	tmp := make(map[string]any)
	err := cursor.Decode(&tmp)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return err
	}

	resSlice := make([]DataStruct, 0)

	if len(tmp) == 0 {
		insertMap := make(map[string]any)
		err = json.Unmarshal(data, &insertMap)
		if err != nil {
			return fmt.Errorf("Unmarshall error: %v", err)
		}
		firstVersion, ok := insertMap["data"].([]interface{})
		if !ok {
			return errors.New("First slice getting error")
		}
		newData := DataStruct{
			ID:          primitive.NewObjectID(),
			ServiceName: serviceName,
			Config:      firstVersion,
		}
		insertSlice := make([]interface{}, 0)
		insertSlice = append(insertSlice, newData)

		res, err := m.InsertMany(context.TODO(), insertSlice)
		if err != nil {
			return errors.New("InsertOne() error")
		}
		fmt.Printf("inserted ids: %v\n", res.InsertedIDs)
	} else {
		sliceData, ok := tmp["data"].([]interface{})
		if !ok {
			return err
		}
		resSlice = append(resSlice, DataStruct{
			ID:          primitive.NewObjectID(),
			ServiceName: serviceName,
			Config:      sliceData,
		})
	}
	return nil
}
