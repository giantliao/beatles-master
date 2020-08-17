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
	"github.com/giantliao/beatles-master/app/cmdclient"
	"github.com/giantliao/beatles-master/app/cmdcommon"
	"github.com/kprc/libeth/account"
	"log"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a miner with miner's beatles address",
	Long:  `remove a miner with miner's beatles address'`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := cmdcommon.IsProcessStarted(); err != nil {
			log.Println(err)
			return
		}

		if len(args) != 1 {
			log.Println("please input miner's beatles address")
			return
		}

		addr := account.BeatleAddress(args[0])
		if !addr.IsValid() {
			log.Println("miner's beatles address not correct")
			return
		}

		var param []string
		param = append(param, addr.String())

		cmdclient.StringOpCmdSend("", cmdcommon.CMD_MINER_REMOVE, param)
	},
}

func init() {
	minerCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
