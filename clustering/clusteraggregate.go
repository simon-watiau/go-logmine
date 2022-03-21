package clustering

import (
	"context"

	"github.com/simon-watiau/go-logmine/tokenizer"
)

type ClusterAggregate struct {
	distances   []float64
	clusterSets []*ClusterSet
}

type WeightedLogs struct {
	Pattern tokenizer.TokenizedLog
	Weight  int
}

func NewClusterAggregate(ctx context.Context, distances []float64) ClusterAggregate {
	clusterSets := make([]*ClusterSet, len(distances))

	for i := 0; i < len(distances); i++ {
		newClusterSet := NewClusterSet(distances[i])
		clusterSets[i] = &newClusterSet
	}

	return ClusterAggregate{
		distances:   distances,
		clusterSets: clusterSets,
	}
}

func (o *ClusterAggregate) AddLog(log tokenizer.TokenizedLog) {
	o.clusterSets[0].AppendLog(log, 1)
}

func (o *ClusterAggregate) Aggregate() []WeightedLogs {
	for i := 1; i < len(o.clusterSets); i++ {
		previousClusters := o.clusterSets[i-1]
		clusters := previousClusters.Clusters
		for _, cluster := range clusters {
			o.clusterSets[i].AppendLog(cluster.Pattern(), cluster.Count)
		}
	}

	var lastClusterPatterns []WeightedLogs
	for _, cluster := range o.clusterSets[len(o.clusterSets)-1].Clusters {
		lastClusterPatterns = append(lastClusterPatterns, WeightedLogs{
			Pattern: cluster.Pattern(),
			Weight:  cluster.Count,
		})
	}

	return lastClusterPatterns
}
