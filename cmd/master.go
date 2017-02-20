// Copyright © 2016 Owen Butler <owen.butler@gmail.com>
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
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "List master node",
	Long:  `List master node of the cluster`,
	Run:   RunMaster,
}

func RunMaster(cmd *cobra.Command, args []string) {
	master(ElasticURL)
}

func master(url string) {

	state := getClusterState(url)

	table := uitable.New()
	if PrintHeaders {
		table.AddRow("nodeid", "address", "name")
	}
	masterNode := state.Nodes[state.MasterNode]
	table.AddRow(
		state.MasterNode,
		masterNode.TransportAddress,
		masterNode.Name)

	fmt.Println(table)
}

func init() {
	RootCmd.AddCommand(masterCmd)
}
