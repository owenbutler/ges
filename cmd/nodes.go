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

func nodes(url string) {

	state := getClusterState(url)

	table := uitable.New()
	if PrintHeaders {
		table.AddRow("id", "address", "master?", "name")
	}
	for nodeID, node := range state.Nodes {
		master := ""
		if nodeID == state.MasterNode {
			master = "*"
		}
		table.AddRow(nodeID, node.TransportAddress, master, node.Name)
	}
	fmt.Println(table)
}

func RunNodes(cmd *cobra.Command, args []string) {
	nodes(ElasticURL)
}

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Node stats",
	Long:  `Node stats, including master node identity`,
	Run:   RunNodes,
}

func init() {
	RootCmd.AddCommand(nodesCmd)
}
