package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromRawString(t *testing.T) {
	tokenizedLog := NewTokenizedLogFromRawString(
		"hello user#2  127.0.0.1:user#4   toto123 ",
	)
	tokens := []Token{
		Token("hello"),
		Token("user#2"),
		IPV4,
		Token(":"),
		Token("user#4"),
		WORD,
	}

	assert.Equal(t, tokens, tokenizedLog.Tokens)
}
