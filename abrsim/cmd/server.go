// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	// "fmt"
	"github.com/spf13/cobra"
	"github.com/mami-project/trafic/abrsim/abr"
)

var serverIp string
var serverPort int
var serverSingle bool

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start an abr server",
	Long: `Start an ABR server.
It will basically sit there and wait for the client to request bunches of data
over a TCP connection`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("server called")
		// fmt.Printf("ip =   %s\n", serverIp)
		// fmt.Printf("port = %d\n", serverPort)
		abr.Server(serverIp, serverPort, serverSingle)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.PersistentFlags().StringVarP(&serverIp, "ip", "I", "127.0.0.1", "IP address of the abrsim server")
	serverCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 8081, "TCP port of the abrsim server")
	serverCmd.PersistentFlags().BoolVarP(&serverSingle,"one-off", "1", false, "Just accept one connection and quit (default is run until killed)")
}