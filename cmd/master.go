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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

type Node struct {
	Name             string `json:"name"`
	TransportAddress string `json:"transport_address"`
}

type MasterState struct {
	MasterNode string          `json:"master_node"`
	Nodes      map[string]Node `json:"nodes"`
}

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "list master node",
	Long:  `list master node of the cluster`,
	Run:   RunMaster,
}

func RunMaster(cmd *cobra.Command, args []string) {
	master(ElasticURL)
}

func master(url string) {
	resp, err := http.Get(fmt.Sprintf("%s/_cluster/state", url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading master : %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading master : %s\n", err)
		os.Exit(1)
	}

	state := MasterState{}
	json.Unmarshal(respBytes, &state)

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
