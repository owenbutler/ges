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

func heap(esUrl string) {

	nodesState := getClusterNodeState(esUrl)

	table := uitable.New()
	if PrintHeaders {
		table.AddRow("id", "old gen", "max", "ratio", "name")
	}
	for nodeID, node := range nodesState.Nodes {
		oldGenPool := node.Jvm.Mem.Pools["CMS Old Gen"]
		ratio := (float64(oldGenPool.UsedInBytes) / float64(oldGenPool.MaxInBytes)) * 100
		table.AddRow(
			nodeID,
			oldGenPool.Used,
			oldGenPool.Max,
			fmt.Sprintf("%.2f%%", ratio),
			node.Name,
		)

	}
	fmt.Println(table)
}

func RunHeap(cmd *cobra.Command, args []string) {
	heap(ElasticURL)
}

// heapCmd represents the heap command
var heapCmd = &cobra.Command{
	Use:   "heap",
	Short: "Heap statistics for nodes",
	Long:  `Show JVM heap statistics for nodes.`,
	Run:   RunHeap,
}

func init() {
	RootCmd.AddCommand(heapCmd)
}
