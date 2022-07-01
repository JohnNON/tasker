package mongostore

import (
	"context"
	"time"

	"github.com/JohnNON/tasker/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewStore(ctx context.Context, collection *mongo.Collection) *Store {
	return &Store{
		ctx:        ctx,
		collection: collection,
	}
}

func (st Store) CreateTask(task *model.Task) error {
	if st.checkTask(task.Type) {
		return nil
	}

	_, err := st.collection.InsertOne(st.ctx, task)
	return err
}

func (st Store) CompleteTask(taskType string) error {
	filter := bson.D{
		primitive.E{Key: "type", Value: taskType},
		primitive.E{Key: "state", Value: 0},
	}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "state", Value: 1},
		primitive.E{Key: "finished_at", Value: time.Now().UTC()},
	}}}

	t := &model.Task{}
	return st.collection.FindOneAndUpdate(st.ctx, filter, update).Decode(t)
}

func (st Store) checkTask(taskType string) bool {
	filter := bson.D{
		primitive.E{Key: "type", Value: taskType},
		primitive.E{Key: "state", Value: 0},
	}

	t := &model.Task{}
	return st.collection.FindOne(st.ctx, filter).Decode(t) != mongo.ErrNoDocuments
}
