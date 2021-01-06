package db

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

type Task struct {
	ID    int
	Value string
}

var task_list = []byte("Task List")

var home_file, _ = os.UserHomeDir()

var task_file string = home_file + "/task.db"

func CreateTask(args []string) (int, error) {
	arg := args[0:]
	task := strings.Join(arg, " ")
	db, err := bolt.Open(task_file, 0644, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	value := []byte(task)
	var id int
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(task_list)
		if err != nil {
			return err
		}
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)
		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	fmt.Println(task + " added")
	return id, nil
}

func TaskDone(key int) error {
	db, err := bolt.Open(task_file, 0644, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(task_list)
		return b.Delete(itob(key))
	})

}

func AllTasks() ([]Task, error) {
	var tasks []Task

	db, err := bolt.Open(task_file, 0644, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(task_list)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				ID:    btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
