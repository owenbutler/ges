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
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable shard allocation",
	Long:  `Disable shard allocation`,
	Run:   RunDisableAllocation,
}

func RunDisableAllocation(cmd *cobra.Command, args []string) {
	disableAllocation(ElasticURL)
}

func disableAllocation(esURL string) {
	req, err := http.NewRequest("PUT", esURL+"/_cluster/settings", bytes.NewBuffer([]byte(`{"transient":{"cluster.routing.allocation.disable_allocation":true}}`)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create request to disable allocation : %s", err)
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to disable allocation : %s", err)
		os.Exit(1)
	}

	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
}

func init() {
	allocationCmd.AddCommand(disableCmd)
}
