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
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task complete",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			home_file, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}
			task_file := home_file + "/task.db"
			db, err := bolt.Open(task_file, 0644, nil)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			err = db.Update(func(tx *bolt.Tx) error {
				bucket := tx.Bucket(task_list)
				if bucket == nil {
					return fmt.Errorf("Bucket %q not found!", task_list)
				}

				c := bucket.Cursor()

				for k, _ := c.First(); k != nil; k, _ = c.Next() {
					key := binary.BigEndian.Uint64(k)
					del, err := strconv.ParseUint(args[0], 10, 64)
					if err != nil {
						panic(err)
					}
					if del == key {
						fmt.Println("Lista encontrada")
						c.Delete()
					}
				}
				return nil
			})
		*/

		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.TaskDone(task.ID)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
