package tokenizer

import (
	"bytes"
	"encoding/gob"
	"log"
	"math"
)

type TokenizedLog struct {
	Tokens []Token
}

func NewTokenizedLogFromBytes(input []byte) (TokenizedLog, error) {
	buf := bytes.NewBuffer(input)
	dec := gob.NewDecoder(buf)

	var log TokenizedLog

	err := dec.Decode(&log)
	if err != nil {
		return TokenizedLog{}, err
	}

	return log, nil
}

func NewTokenizedLogFromTokens(dataTypes []Token) TokenizedLog {
	return TokenizedLog{
		Tokens: dataTypes,
	}
}

func (t TokenizedLog) ToBytes() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(t); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func (t TokenizedLog) Distance(log1 TokenizedLog) float64 {
	var dist float64 = 1
	minLen := int(math.Min(float64(len(t.Tokens)), float64(len(log1.Tokens))))
	maxLen := int(math.Max(float64(len(t.Tokens)), float64(len(log1.Tokens))))
	for i := 0; i < int(minLen); i++ {
		if t.Tokens[i] == log1.Tokens[i] {
			dist -= float64(1) / float64(maxLen)
		}
	}

	return dist
}
