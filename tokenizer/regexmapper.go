package tokenizer

import (
	"regexp"
)

type regexMapper struct {
	token  Token
	regexs []*regexp.Regexp
}

func newRegexMapper(dataType Token, regexStrs []string) regexMapper {
	var regexs = make([]*regexp.Regexp, len(regexStrs))

	for idx, regexStr := range regexStrs {
		var err error
		regexs[idx], err = regexp.Compile(regexStr)
		if err != nil {
			panic("configuration error")
		}
	}

	return regexMapper{
		token:  dataType,
		regexs: regexs,
	}
}

func (o regexMapper) supports(token string) bool {
	for _, regex := range o.regexs {
		if regex.MatchString(token) {
			return true
		}
	}

	return false
}

func (o regexMapper) getToken() Token {
	return o.token
}
