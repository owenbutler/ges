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

// indicesCmd represents the indices command
var indicesCmd = &cobra.Command{
	Use:   "indices",
	Short: "Show elasticsearch indices",
	Long:  "Show all elasticsearch indexes with some size info",
	Run:   RunIndices,
}

func RunIndices(cmd *cobra.Command, args []string) {
	indices(ElasticURL)
}

func indices(elasticURL string) {

	indicesState := getIndicesState(elasticURL)

	table := uitable.New()
	if PrintHeaders {
		table.AddRow("status", "name", "pri", "rep", "size", "docs")
	}

	for indexName, index := range indicesState.Indices {
		status := "green"
		size := index.IndexSize.Size
		docs := index.Docs.NumDocs

		primaries := 0
		replicas := 0

		for _, shards := range index.Shards {
			for _, shard := range shards {
				if shard.Routing.Primary {
					primaries++
				}
				if !shard.Routing.Primary {
					replicas++
				}
			}
		}

		table.AddRow(status, indexName, primaries, replicas, size, docs)
	}

	fmt.Println(table)
}

func init() {
	RootCmd.AddCommand(indicesCmd)
}
