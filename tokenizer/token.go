package tokenizer

type Token string

const (
	ANY      Token = "ANY"
	IPV4     Token = "IPV4"
	DATE     Token = "DATE"
	TIME     Token = "TIME"
	NOTSPACE Token = "NOTSPACE"
	NUMBER   Token = "NUMBER"
	WORD     Token = "WORD"
	ALIGNER  Token = "ALIGNER"
	TAG      Token = "TAG"
)
