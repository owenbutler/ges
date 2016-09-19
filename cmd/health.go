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
	"os"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"gopkg.in/olivere/elastic.v1"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "cluster health",
	Long:  `Cluster health and general health statistics`,
	Run:   RunHealth,
}

func init() {
	RootCmd.AddCommand(healthCmd)
}

func RunHealth(cmd *cobra.Command, args []string) {
	health(elasticClient)
}

func health(esClient *elastic.Client) {
	health, err := esClient.ClusterHealth().Do()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching health : %s\n", err)
		return
	}

	table := uitable.New()
	if PrintHeaders {
		table.AddRow("cluster", "status", "nodes", "data", "pri", "shards", "relo", "init", "unassign")
	}
	table.AddRow(
		health.ClusterName,
		health.Status,
		health.NumberOfNodes,
		health.NumberOfDataNodes,
		health.ActivePrimaryShards,
		health.ActiveShards,
		health.RelocatingShards,
		health.InitializedShards,
		health.UnassignedShards)

	fmt.Println(table)
}
