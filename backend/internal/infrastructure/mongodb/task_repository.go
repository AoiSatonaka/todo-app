package mongodb

import (
	"context"
	"os"

	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct{}

var _ task.ITaskRepository = &TaskRepository{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func getTaskCollection(c *mongo.Client) *mongo.Collection {
	return c.
		Database(os.Getenv("DB_NAME")).
		Collection(os.Getenv("TASK_COLLECTION_NAME"))
}

func (tr *TaskRepository) List(uid string) (*[]task.Task, error) {
	cli, err := connectClient()
	if err != nil {
		return nil, err
	}
	defer closeClient(cli)

	coll := getTaskCollection(cli)
	cur, err := coll.Find(context.TODO(), bson.D{{Key: "uid", Value: uid}})
	if err != nil {
		return nil, err
	}

	var results []taskSchema
	if err := cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	var tasks []task.Task
	for _, result := range results {
		cur.Decode(&result)
		tasks = append(tasks, result.Task)
	}

	return &tasks, nil
}

func (tr *TaskRepository) FindById(i, uid string) (*task.Task, error) {
	cli, err := connectClient()
	if err != nil {
		return nil, err
	}
	defer closeClient(cli)

	coll := getTaskCollection(cli)

	var result taskSchema
	err = coll.FindOne(
		context.TODO(),
		bson.D{{Key: "uid", Value: uid}, {Key: "_id", Value: i}},
		options.FindOne()).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result.Task, nil
}

func (tr *TaskRepository) Create(t *task.Task, uid string) error {
	cli, err := connectClient()
	if err != nil {
		return err
	}
	defer closeClient(cli)

	// task to dbSchema
	ts, err := newTaskSchema(*t, uid)
	if err != nil {
		return err
	}

	coll := getTaskCollection(cli)
	if _, err = coll.InsertOne(context.TODO(), *ts); err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) Update(t *task.Task, uid string) error {
	cli, err := connectClient()
	if err != nil {
		return err
	}
	defer closeClient(cli)

	coll := getTaskCollection(cli)

	// task authorize
	if _, err := tr.FindById(t.Id, uid); err != nil {
		return err
	}

	if _, err = coll.UpdateByID(context.TODO(), t.Id, t); err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) Delete(i, uid string) error {
	cli, err := connectClient()
	if err != nil {
		return err
	}
	defer closeClient(cli)

	coll := getTaskCollection(cli)

	// task authorize
	if _, err := tr.FindById(i, uid); err != nil {
		return err
	}

	if _, err = coll.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: i}}); err != nil {
		return err
	}

	return nil
}
