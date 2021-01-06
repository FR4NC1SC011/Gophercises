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
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show tasks",
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

			err = db.View(func(tx *bolt.Tx) error {
				bucket := tx.Bucket(task_list)
				if bucket == nil {
					return fmt.Errorf("Bucket %q not found!", task_list)
				}
				c := bucket.Cursor()

				for k, v := c.First(); k != nil; k, v = c.Next() {
					key := binary.BigEndian.Uint64(k)
					fmt.Printf("%v. %s\n", key, v)
				}
				return nil
			})
		*/
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vacation? 🏖")
			return
		}
		fmt.Println("You have the following tasks: ")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
