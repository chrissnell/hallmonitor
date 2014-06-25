// hallmonitor - Intelligent HAproxy for ElasticSearch
//
// clusterstate.go - Objects and methods for polling cluster status API
//
// (c) 2014, Christopher Snell

package elasticsearch

import (
	"encoding/json"
	"log"
	"net/http"
)

// The ClusterState struct is built to match the data structure returned
// by the ElasticSearch Cluster Health API.
type ClusterState struct {
	Master string           `json:"master_node"`
	Nodes  map[string]*Node `json:"nodes"`
}

// Poll the Cluster Health API on this Node for cluster state information.
// We store the cluster state in this Node's ClusterView object.
func GetClusterState(n *Node) error {

}
