// hallmonitor - Intelligent HAproxy for ElasticSearch
//
// node.go - ES node object and methods
//
// (c) 2014, Christopher Snell

package elasticsearch

// A Cluster is a map (hash) of Nodes.  Each Node also has a ClusterView, which is another Cluster
// object that holds that Node's view of the cluster state.   With the ES  Cluster Health API,
// every Node has it's own view of what the cluster looks like: who is the master, who are the members,
//  shards, etc.
type Node struct {
	Name             string         `json:"name"`
	TransportAddress string         `json:"transport_address"`
	Attributes       NodeAttributes `json:"attributes"`
	ClusterView      Cluster
	Unhealthy        bool
	Master           bool
	MasterEligible   bool
	DataEligible     bool
	ActiveInLBPool   bool
}

type NodeAttributes struct {
	MasterEligible bool `json:"master"`
}

// Returns the number of nodes in the cluster, as seen by this Node
func (n *Node) NodesSeen() uint8 {
	return (len(n.ClusterView.Nodes))
}

func (n *Node) IPAddress() string {
	// generate IP address from Node.TransportAddress
}
