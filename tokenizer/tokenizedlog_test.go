package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromToBytes(t *testing.T) {
	tokens := []Token{
		DATE,
		ALIGNER,
		ANY,
		IPV4,
	}

	tokenizedLog := NewTokenizedLogFromTokens(tokens)

	bytes := tokenizedLog.ToBytes()

	newTokenizedLog, err := NewTokenizedLogFromBytes(bytes)

	assert.Nil(t, err)
	assert.Equal(t, tokens, newTokenizedLog.Tokens)
}

func TestDistanceZero(t *testing.T) {
	tokenizedLog1 := NewTokenizedLogFromTokens([]Token{
		DATE,
		ALIGNER,
		ANY,
		IPV4,
	})

	tokenizedLog2 := NewTokenizedLogFromTokens([]Token{
		DATE,
		ALIGNER,
		ANY,
		IPV4,
	})

	assert.Equal(t, float64(0), tokenizedLog1.Distance(tokenizedLog2))
	assert.Equal(t, float64(0), tokenizedLog2.Distance(tokenizedLog1))
}

func TestDistanceNotZero(t *testing.T) {
	tokenizedLog1 := NewTokenizedLogFromTokens([]Token{
		DATE,
		ALIGNER,
		ANY,
		IPV4,
	})

	tokenizedLog2 := NewTokenizedLogFromTokens([]Token{
		DATE,
		ALIGNER,
		IPV4,
		ANY,
	})

	assert.Equal(t, float64(0.5), tokenizedLog1.Distance(tokenizedLog2))
	assert.Equal(t, float64(0.5), tokenizedLog2.Distance(tokenizedLog1))
}
