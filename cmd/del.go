// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/RowlingWu/agenda/entity"
	"log"
	"os"
	"encoding/json"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete current user",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		users := [6]entity.User{
			{"rowling", "a", "@", "13"},
			{"bob", "b", "#", "4"},
			{"alice", "a", "e", "5"},
			{"faith", "f", "a", "1"},
			{"john", "j", "e", "e"},
			{"mary", "m", "$", "4"},
		}
		f, err := os.OpenFile("entity/userInfo.txt", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, u := range users {
			b, _ := json.Marshal(u)
			f.WriteString(string(b) + "\n")
		}
		f.Close()


		// TODO: Work your own magic here
		// read username from curUser.txt
		log.Println("read info of the current user...")
		name := entity.ReadCur()

		// clear curUser.txt
		log.Println("delete current user...")
		entity.ClearCurUsr()
		// seek userInfo in userInfo.txt and delete it
		entity.SeekUsr(name)
		log.Println("success in deleting current user")
	},
}

func init() {
	RootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
