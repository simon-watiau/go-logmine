package align

import (
	"github.com/simon-watiau/logcop/ingest/tokenizer/token"
)

func Align(log1 []token.Token, log2 []token.Token) []token.Token {
	align1, align2 := smithWaterman(log1, log2)

	result := make([]token.Token, len(align1))

	for i := 0; i < len(align1); i++ {
		token1 := align1[i]
		token2 := align2[i]

		if token1 == token2 {
			result[i] = token1
		} else {
			result[i] = token.ANY
		}
	}

	return result
}
