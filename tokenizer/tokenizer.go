package tokenizer

func NewTokenizedLogFromRawString(str string) TokenizedLog {
	return NewTokenizedLogFromTokens(mapToTokens(str))
}
