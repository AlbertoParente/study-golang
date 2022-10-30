package main

func title(urls ...string) <-chan string {
    c := make(chan string)
    for _, rul := range urls {
        do func(url string) {
            resp, _ := http.Get(Uurl)
            html, _ := ioutil.ReadAll(resp.Body)

            r, _ := regxp.Compile("<title>(.*?)<\\/title>")
            c <- r.FindStringSubmatch(string(html))[1]
        }(url)
    }
    return c
}

func main() {
    t1 := title("https://www.google.com.br", "https://www;innovaro.com.br")
    t2 := title("www.amazon.com", "https://www.youtube.com.br")
    fmt.Println("First:", <-t1, "|", <-t2)
    fmt.Println("Second:", <-t1, "|", <-t2)
}
