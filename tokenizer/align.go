package tokenizer

func Align(log1 []Token, log2 []Token) []Token {
	align1, align2 := smithWaterman(log1, log2)

	result := make([]Token, len(align1))

	for i := 0; i < len(align1); i++ {
		token1 := align1[i]
		token2 := align2[i]

		if token1 == token2 {
			result[i] = token1
		} else {
			result[i] = ANY
		}
	}

	return result
}
