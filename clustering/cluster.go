package clustering

import (
	"github.com/simon-watiau/go-logmine/tokenizer"
)

type Cluster struct {
	logs  []tokenizer.TokenizedLog
	Count int
}

func NewCluster(log tokenizer.TokenizedLog, weight int) Cluster {
	return Cluster{
		Count: weight,
		logs: []tokenizer.TokenizedLog{
			log,
		},
	}
}

func (o Cluster) DistanceToLog(log tokenizer.TokenizedLog) float64 {
	return o.logs[0].Distance(log)
}

func (o *Cluster) AppendLog(log tokenizer.TokenizedLog, weight int) {
	o.logs = append(o.logs, log)
	o.Count += weight
}

func (o *Cluster) Pattern() tokenizer.TokenizedLog {
	existingEntry := o.logs[0].Tokens
	for _, entry := range o.logs[1:] {
		existingEntry = tokenizer.Align(existingEntry, entry.Tokens)
	}

	return tokenizer.NewTokenizedLogFromTokens(existingEntry)
}
