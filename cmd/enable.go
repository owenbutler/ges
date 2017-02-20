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

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable shard allocation",
	Long:  `Enable shard allocation`,
	Run:   RunEnableAllocation,
}

func RunEnableAllocation(cmd *cobra.Command, args []string) {
	enableAllocation(ElasticURL)
}

func enableAllocation(esURL string) {
	req, err := http.NewRequest("PUT", esURL+"/_cluster/settings", bytes.NewBuffer([]byte(`{"transient":{"cluster.routing.allocation.disable_allocation":false}}`)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create request to enable allocation : %s", err)
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to enable allocation : %s", err)
		os.Exit(1)
	}

	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
}

func init() {
	allocationCmd.AddCommand(enableCmd)
}
