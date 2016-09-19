package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Node struct {
	Name             string `json:"name"`
	TransportAddress string `json:"transport_address"`
}

type ClusterState struct {
	MasterNode string          `json:"master_node"`
	Nodes      map[string]Node `json:"nodes"`
}

func getClusterState(url string) ClusterState {
	resp, err := http.Get(fmt.Sprintf("%s/_cluster/state", url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading cluster state : %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading cluster state : %s\n", err)
		os.Exit(1)
	}

	state := ClusterState{}
	json.Unmarshal(respBytes, &state)
	return state
}
