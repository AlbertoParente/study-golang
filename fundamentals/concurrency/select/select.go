package main

func theSpeed(url1, url2, url3 string) string {
	c1 := html.Title(url1)
	c2 := html.Title(url2)
	c3 := html.Title(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	}
}
