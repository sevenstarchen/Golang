package main

import (
	"test/crawler/engine"
	"test/crawler/zhenai/parser"
)

/*
func printCityList(contents []byte) {
	//re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a> `)
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {

		fmt.Printf("City: %s,	URL: %s\n", m[2], m[1])

	}
	fmt.Printf("Matches found:%d\n", len(matches))
}*/

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
