// Copyright © 2017 Owen Butler <owen.butler@gmail.com>
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

// shardsCmd represents the shards command
var shardsCmd = &cobra.Command{
	Use:   "shards",
	Short: "List shards in the cluster",
	Long:  `List all shards in the cluster, including size and movements`,
	Run:   RunShards,
}

func RunShards(cmd *cobra.Command, args []string) {
	shards(ElasticURL)
}

func shards(elasticURL string) {

	shardsState := getIndicesState(elasticURL)

	table := uitable.New()
	if PrintHeaders {
		table.AddRow("index", "shard", "pri/rep", "state", "docs", "size", "node", "relocating")
	}

	for indexName, index := range shardsState.Indices {

		for shardId, shards := range index.Shards {
			for _, shard := range shards {

				primaryOrReplica := "r"
				if shard.Routing.Primary {
					primaryOrReplica = "p"
				}
				table.AddRow(indexName,
					shardId,
					primaryOrReplica,
					shard.Routing.State,
					shard.Docs.NumDocs,
					shard.Size.Size,
					shard.Routing.Node,
					shard.Routing.RelocatingNode,
				)
			}
		}
	}

	fmt.Println(table)
}

func init() {
	RootCmd.AddCommand(shardsCmd)
}
