package cluster

import (
	"github.com/simon-watiau/logcop/tokenizer"
)

type Cluster struct {
	logs  []tokenizer.TokenizedLog
	Count int
}

func NewCluster(log tokenizer.TokenizedLog) Cluster {
	return Cluster{
		Count: 0,
		logs: []tokenizer.TokenizedLog{
			log,
		},
	}
}

func (o Cluster) DistanceToLog(log tokenizer.TokenizedLog) float64 {
	return o.logs[0].LogDistance(log)
}

func (o *Cluster) AppendLog(log tokenizer.TokenizedLog) {
	o.logs = append(o.logs, log)
	o.Count = o.Count + 1
}

func (o *Cluster) Pattern() tokenizer.TokenizedLog {
	existingEntry := o.logs[0].Tokens()
	for _, entry := range o.logs[1:] {
		existingEntry = tokenizer.Align(existingEntry, entry.Tokens())
	}

	return tokenizer.NewFromDataTypes(existingEntry)
}
