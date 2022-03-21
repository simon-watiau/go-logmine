package clustering

import (
	"math"

	"github.com/simon-watiau/go-logmine/tokenizer"
)

type ClusterSet struct {
	maxDistance float64
	Clusters    []*Cluster
}

func NewClusterSet(maxDistance float64) ClusterSet {
	return ClusterSet{
		Clusters:    []*Cluster{},
		maxDistance: maxDistance,
	}
}

func (o *ClusterSet) AppendLog(log tokenizer.TokenizedLog, weight int) {
	var matchedCluster *Cluster = nil
	var distanceToMatchedCluster float64 = math.MaxFloat64

	for _, cluster := range o.Clusters {

		dist := cluster.DistanceToLog(log)

		if dist <= o.maxDistance && dist <= distanceToMatchedCluster {
			matchedCluster = cluster
			distanceToMatchedCluster = dist
		}
	}

	if matchedCluster == nil {
		newCluster := NewCluster(log, weight)
		newCluster.Count = weight
		o.Clusters = append(o.Clusters, &newCluster)
	} else {
		matchedCluster.AppendLog(log, weight)
	}
}
