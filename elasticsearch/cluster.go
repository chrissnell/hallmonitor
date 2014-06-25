// hallmonitor - Intelligent HAproxy for ElasticSearch
//
// cluster.go - ES Cluster object and methods
//
// (c) 2014, Christopher Snell

package elasticsearch

import (
	"log"
)

type Cluster struct {
	Nodes map[string]*Node
}

func (c *Cluster) GetClusterStateForAllNodes() err {
	for name, node := range c.Nodes {
		err := GetClusterState(node)
		if err != nil {
			log.Printf("GetClusterState() failed for %v / %v: %v\n", node.Name, node.TransportAddress, err)
			node.Unhealthy = true
		}
	}
}

func NewCluster(nodes []string) c *Cluster {
	for _, v := range nodes {
		c.Nodes[v].
	}
}
