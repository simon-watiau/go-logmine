package cluster

import (
	"context"

	"github.com/simon-watiau/logcop/tokenizer"
)

type ClusterAggregate struct {
	distances   []float64
	clusterSets []*ClusterSet
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
	o.clusterSets[0].AppendLog(log)
}

func (o *ClusterAggregate) Aggregate() []tokenizer.TokenizedLog {
	for i := 1; i < len(o.clusterSets); i++ {
		previousClusters := o.clusterSets[i-1]
		patterns := previousClusters.Patterns()
		for _, pattern := range patterns {
			o.clusterSets[i].AppendLog(pattern)
		}
	}

	return o.clusterSets[len(o.clusterSets)-1].Patterns()
}
