package citylist

import (
	"go-crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com)">`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: nil,
		})
	}
	return result
}
