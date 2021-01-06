/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/binary"

	"task/db"

	"github.com/spf13/cobra"
)

var task_list = []byte("Task List")
var t Task

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			home_file, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}
			task_file := home_file + "/task.db"

			arg := args[0:]
			task := strings.Join(arg, " ")
			db, err := bolt.Open(task_file, 0644, nil)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			value := []byte(task)
			err = db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists(task_list)
				if err != nil {
					return err
				}
				key, _ := bucket.NextSequence()
				t.ID = int(key)
				err = bucket.Put(itob(t.ID), value)
				if err != nil {
					return err
				}
				return nil
			})

			fmt.Println(task + " added")
		*/
		db.CreateTask(args)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

type Task struct {
	ID int
}
