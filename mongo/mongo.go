package mongoDb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoInitialConn(mongoUrl string) (*mongo.Client, error) {
	cOpts := options.Client().ApplyURI(mongoUrl)               // Получение опций из урла подключения
	mClient, err := mongo.Connect(context.Background(), cOpts) // Подключение к бд
	if err != nil {
		return nil, fmt.Errorf("[MONGO]: %v", err)
	}

	defer func() {
		if err := mClient.Disconnect(context.Background()); err != nil { // Закрытие соединения
			fmt.Printf("[MONGO]: %v", err)
		}
	}()
	return mClient, nil
}
