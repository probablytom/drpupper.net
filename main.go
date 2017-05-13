package main

import (
  "fmt"
  "net/http"
  "github.com/yhat/scrape"
  "golang.org/x/net/html"
  "golang.org/x/net/html/atom"
)

func main() {

  // Parse the subreddit
  resp, err := http.Get("https://reddit.com/r/aww/")
  if err != nil {
    panic(err)
  }

  root, err := html.Parse(resp.Body)
  if err != nil {
    panic(err)
  }

  matcher := func(n *html.Node) bool {
    // must check for nil values!
    return n.DataAtom == atom.A && n.Parent.DataAtom == atom.P// && scrape.Attr(n.Parent, "class") == "title"
  }

  // Pick up all nodes filtered by the matcher
  image_urls := scrape.FindAll(root, matcher)
  fmt.Printf("Finished scraping!\n")
  for i, url_node := range image_urls {
    url := scrape.Attr(url_node, "href")
    if url[:2] == "/r" {
      url = "https://reddit.com" + url
    }
    fmt.Printf("%2d: %s\n", i, url)
//    fmt.Printf("parent class: %s \n\n", scrape.Attr(url_node, "class"))
  }

}
