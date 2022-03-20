package cluster

import (
	"math"

	"github.com/simon-watiau/logcop/tokenizer"
)

type ClusterSet struct {
	maxDistance float64
	clusters    []*Cluster
}

func NewClusterSet(maxDistance float64) ClusterSet {
	return ClusterSet{
		clusters:    []*Cluster{},
		maxDistance: maxDistance,
	}
}

func (o *ClusterSet) AppendLog(log tokenizer.TokenizedLog) {
	var matchedCluster *Cluster = nil
	var distanceToMatchedCluster float64 = math.MaxFloat64

	for _, cluster := range o.clusters {

		dist := cluster.DistanceToLog(log)

		if dist <= o.maxDistance && dist <= distanceToMatchedCluster {
			matchedCluster = cluster
			distanceToMatchedCluster = dist
		}
	}

	if matchedCluster == nil {
		newCluster := NewCluster(log)
		o.clusters = append(o.clusters, &newCluster)
	} else {
		matchedCluster.AppendLog(log)
	}
}

func (o *ClusterSet) Patterns() []tokenizer.TokenizedLog {
	patterns := make([]tokenizer.TokenizedLog, len(o.clusters))
	for index, cluster := range o.clusters {
		patterns[index] = cluster.Pattern()
	}
	return patterns
}
