package main

import (
	"fmt"
	"html"
	"time"
)

func theSpeed(url1, url2, url3, url4 string) string {
	c1 := html.Title(url1)
	c2 := html.Title(url2)
	c3 := html.Title(url3)
	c4 := html.Title(url4)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case t4 := <-c4:
		return t4
	case <-time.After(1000 * time.Millisecond):
		return "All loser"
	}
}

func main() {
	winner := theSpeed(
		"http://www.youtube.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.innovaro.com.br",
	)
	fmt.Println(winner)
}
