package align

import (
	"github.com/simon-watiau/logcop/ingest/tokenizer/token"
)

const (
	MatchAward      int = 10
	MismatchPenalty int = 1
	GapPenalty      int = 0
)

func zeros(rows int, cols int) [][]int {
	retVal := make([][]int, rows)
	for r := range retVal {
		retVal[r] = make([]int, cols)
	}

	return retVal
}

func matchScore(a token.Token, b token.Token) int {
	if a == b {
		return MatchAward
	}

	if a == token.ALIGNER || b == token.ALIGNER {
		return GapPenalty
	}

	return MismatchPenalty
}

func maxInt(nums ...int) int {
	max := nums[0]
	for _, i := range nums {
		if max < i {
			max = i
		}
	}

	return max
}

func reverse(tokens []token.Token) []token.Token {
	for i, j := 0, len(tokens)-1; i < j; i, j = i+1, j-1 {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	}
	return tokens
}

func finalize(align1 []token.Token, align2 []token.Token) ([]token.Token, []token.Token) {

	align1 = reverse(align1)
	align2 = reverse(align2)

	symbol := token.Token("")

	score := 0
	identity := float64(0)
	for i := 0; i < len(align1); i++ {
		if align1[i] == align2[i] {
			symbol = symbol + align1[i]
			identity = identity + 1
			score += matchScore(align1[i], align2[i])
		} else if align1[i] != align2[i] && align1[i] != token.ALIGNER && align2[i] != token.ALIGNER {
			score += matchScore(align1[i], align2[i])
			symbol += " "
		} else if align1[i] == token.ALIGNER || align2[i] == token.ALIGNER {
			symbol += " "
			score += GapPenalty
		}
	}

	return align1, align2
}

func smithWaterman(seq1 []token.Token, seq2 []token.Token) ([]token.Token, []token.Token) {

	m := len(seq1)
	n := len(seq2)

	score := zeros(m+1, n+1)
	pointer := zeros(m+1, n+1)
	maxI := 0
	maxJ := 0

	maxScore := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			scoreDiagonal := score[i-1][j-1] + matchScore(seq1[i-1], seq2[j-1])
			scoreUp := score[i][j-1] + GapPenalty
			scoreLeft := score[i-1][j] + GapPenalty
			score[i][j] = maxInt(0, scoreLeft, scoreUp, scoreDiagonal)
			if score[i][j] == 0 {
				pointer[i][j] = 0
			}

			if score[i][j] == scoreLeft {
				pointer[i][j] = 1
			}

			if score[i][j] == scoreUp {
				pointer[i][j] = 2
			}
			if score[i][j] == scoreDiagonal {
				pointer[i][j] = 3
			}

			if score[i][j] >= maxScore {
				maxI = i
				maxJ = j
				maxScore = score[i][j]
			}
		}
	}

	align1 := []token.Token{}
	align2 := []token.Token{}
	i := maxI
	j := maxJ

	for pointer[i][j] != 0 {
		if pointer[i][j] == 3 {
			align1 = append(align1, seq1[i-1])
			align2 = append(align2, seq2[j-1])
			i--
			j--
		} else if pointer[i][j] == 2 {
			align1 = append(align1, token.ALIGNER)
			align2 = append(align2, seq2[j-1])
			j--
		} else if pointer[i][j] == 1 {
			align1 = append(align1, seq1[i-1])
			align2 = append(align2, token.ALIGNER)
			i--
		}
	}

	a1, a2 := finalize(align1, align2)

	return a1, a2
}
