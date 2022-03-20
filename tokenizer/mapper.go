package tokenizer

import "strings"

type mapper interface {
	supports(str string) bool
	getToken() Token
}

var delimiters []string = []string{"=", ":"}

var tokenizers = []mapper{
	newRegexMapper(DATE, []string{
		"^[0-9]{2}/[0-9]{2}/[0-9]{4}$",
		`^[0-9]{4}-[0-9]{2}-[0-9]{2}(T[0-9]{2}:[0-9]{2}:[0-9]{2}(Z|([\+-]?[0-9]{2}:[0-9]{2})))?$`,
		"^[0-9]{4}/[0-9]{2}/[0-9]{2}$",
	}),
	newRegexMapper(IPV4, []string{
		`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`,
	}),
	newRegexMapper(NUMBER, []string{
		`^((-[0-9]*\.?)?([0-9]+)|(0x[0-9A-F]+))$`,
	}),
	newRegexMapper(TIME, []string{
		"^[0-9]{2}:[0-9]{2}:[0-9]{2}$",
	}),
	newRegexMapper(WORD, []string{
		"^[a-zAA-Z0-9]*[0-9]+[a-zA-Z0-9]*.*$",
	}),
}

func isolateDelimiters(log string) string {
	l := log
	for _, ss := range delimiters {
		l = strings.ReplaceAll(l, ss, " "+ss+" ")
	}

	return l
}

func convertToToken(str string) Token {

	for _, handler := range tokenizers {
		if handler.supports(str) {
			return handler.getToken()
		}
	}

	return Token(str)
}

func mapToTokens(log string) []Token {
	log = isolateDelimiters(log)

	components := strings.Split(log, " ")

	tokens := []Token{}
	for _, token := range components {
		trimmedToken := strings.TrimSpace(token)

		if trimmedToken == "" {
			continue
		}

		tokens = append(
			tokens,
			convertToToken(trimmedToken),
		)
	}

	return tokens
}
