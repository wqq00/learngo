package parser

import (
	"regexp"
	"learngo/helloword/crawler/engine"
)

const city = `(<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(city)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User"+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}

	return result
}
