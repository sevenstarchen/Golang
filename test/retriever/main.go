package main

import (
	"fmt"
	"time"
)
import "test/retriever/real"

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retiever
	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T,%v", r, r)
}
