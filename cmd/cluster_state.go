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

type ClusterNodeState struct {
	Nodes map[string]ClusterNode `json:"nodes"`
}

type ClusterNode struct {
	Name string   `json:"name"`
	Jvm  JvmStats `json:"jvm"`
}

type JvmStats struct {
	Mem MemStats `json:"mem"`
}

type MemStats struct {
	Pools map[string]MemoryPool `json:"pools"`
}

type MemoryPool struct {
	Used            string `json:"used"`
	UsedInBytes     int    `json:"used_in_bytes"`
	Max             string `json:"max"`
	MaxInBytes      int    `json:"max_in_bytes"`
	PeakUsed        string `json:"peak_used"`
	PeakUsedInBytes int    `json:"peak_used_in_bytes"`
	PeakMax         string `json:"peak_max"`
	PeakMaxInBytes  int    `json:"peak_max_in_bytes"`
}

type IndicesState struct {
	Indices map[string]Index `json:"indices"`
}

type Index struct {
	IndexSize IndexSizeStat          `json:"index"`
	Docs      DocSizeStat            `json:"docs"`
	Shards    map[string][]ShardStat `json:"shards"`
}

type IndexSizeStat struct {
	Size      string `json:"size"`
	SizeBytes int64  `json:"size_in_bytes"`
}

type DocSizeStat struct {
	NumDocs     int64 `json:"num_docs"`
	MaxDoc      int64 `json:"max_doc"`
	DeletedDocs int64 `json:"deleted_docs"`
}

type ShardStat struct {
	Routing RoutingStat   `json:"routing"`
	State   string        `json:"state"`
	Size    IndexSizeStat `json:"index"`
	Docs    DocSizeStat   `json:"docs"`
}

type RoutingStat struct {
	State          string `json:"state"`
	Primary        bool   `json:"primary"`
	Node           string `json:"node"`
	RelocatingNode string `json:"relocating_node"`
}

func getBytesForUrl(base, path string) []byte {
	fullUrl := fmt.Sprintf("%s%s", base, path)
	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading from (%s) : %s\n", fullUrl, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading from (%s) : : %s\n", fullUrl, err)
		os.Exit(1)
	}
	return respBytes
}

func getClusterState(url string) ClusterState {
	state := ClusterState{}
	stateBytes := getBytesForUrl(url, "/_cluster/state")

	json.Unmarshal(stateBytes, &state)
	return state
}

func getClusterNodeState(url string) ClusterNodeState {
	clusterNodeState := ClusterNodeState{}
	bytes := getBytesForUrl(url, "/_cluster/nodes/stats?jvm=true")

	json.Unmarshal(bytes, &clusterNodeState)
	return clusterNodeState
}

func getIndicesState(url string) IndicesState {
	state := IndicesState{}
	stateBytes := getBytesForUrl(url, "/_status")

	json.Unmarshal(stateBytes, &state)
	return state
}
