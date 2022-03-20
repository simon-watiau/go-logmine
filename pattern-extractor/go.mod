module github.com/simon-watiau/logcop/pattern-extractor

go 1.13

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/simon-watiau/logcop/dto v0.0.0-20210714215057-1218c1057aaa
	github.com/simon-watiau/logcop/subscriber v0.0.0-00010101000000-000000000000
	github.com/simon-watiau/logcop/tokenizer v0.0.0-00010101000000-000000000000
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.18.1
)

replace github.com/simon-watiau/logcop/dto => ../dto

replace github.com/simon-watiau/logcop/subscriber => ../subscriber

replace github.com/simon-watiau/logcop/tokenizer => ../tokenizer
