package main

import (
	"fmt"
	// "regexp"
	// "sync"

)

// IDS yang Nge cek tanpa scrapping, jadi nge bikin tree sendiri
func idstest(startURL string, endURL string, baseURL string) ([][]string, int) {
	visitedURL := SafeMap[bool]{data: make(map[string]bool)}
	// queriedURL := SafeMap[bool]{data: make(map[string]bool)}
	root := NewNode("https://en.wikipedia.org/wiki/Jokowi")
	root.AddChild("https://en.wikipedia.org/wiki/Indonesia")
	root.AddChild("https://en.wikipedia.org/wiki/Java")
	root.Children["https://en.wikipedia.org/wiki/Indonesia"].AddChild("https://en.wikipedia.org/wiki/Asia")
	root.Children["https://en.wikipedia.org/wiki/Indonesia"].AddChild("https://en.wikipedia.org/wiki/Java")
	root.Children["https://en.wikipedia.org/wiki/Indonesia"].Children["https://en.wikipedia.org/wiki/Java"].AddChild("https://en.wikipedia.org/wiki/Philosophy")
	root.Children["https://en.wikipedia.org/wiki/Indonesia"].Children["https://en.wikipedia.org/wiki/Java"].AddChild("https://en.wikipedia.org/wiki/Asia")
	root.Children["https://en.wikipedia.org/wiki/Java"].AddChild("https://en.wikipedia.org/wiki/Philosophy")
	root.Children["https://en.wikipedia.org/wiki/Java"].AddChild("https://en.wikipedia.org/wiki/Asia")
	root.Children["https://en.wikipedia.org/wiki/Java"].Children["https://en.wikipedia.org/wiki/Asia"].AddChild("https://en.wikipedia.org/wiki/Satu")
	root.Children["https://en.wikipedia.org/wiki/Java"].Children["https://en.wikipedia.org/wiki/Asia"].AddChild("https://en.wikipedia.org/wiki/Dua")

	// fmt.Println("Depth-First Search:")
	var paths = make([][]string, 0)
	var visitedCount int
	// var countArticle int

	var currentLastNodes SafeArray[*Node]
	currentLastNodes.Add(root)

	// depth := 0

	paths, visitedCount = root.DFS(baseURL, endURL)
	return paths, visitedCount

	// Selama belom ketemu satupun solusi, bakalan di loop
	// for len(paths) == 0 {
	// 	var newLastNodes SafeArray[*Node]
	// 	var wg sync.WaitGroup
	// 	wg.Add(len(currentLastNodes.Get()))

	// 	maxConcurrentRequests := 250

	// 	sem := make(chan struct{}, maxConcurrentRequests)

	// 	for _, node := range currentLastNodes.array {
	// 		sem <- struct{}{}
	// 		go func(node *Node) {
	// 			defer func() { <-sem }()
	// 			defer wg.Done()

	// 			if _, ok := queriedURL.Get(node.Value); ok {
	// 				fmt.Println("Already queried")
	// 				return
	// 			}

	// 			doc := makeRequest(node.Value)
	// 			if doc != nil {
	// 				duplicateURL := make(map[string]bool)
	// 				doc.Find("a").Each(func(_ int, s *goquery.Selection) {
	// 					link, _ := s.Attr("href")
	// 					matched, _ := regexp.MatchString("^/wiki/", link)

	// 					if matched && !duplicateURL[baseURL+link] {
	// 						duplicateURL[baseURL+link] = true
	// 						// fmt.Println(baseURL + link)

	// 						_, ok := visitedURL.Get(baseURL + link)
	// 						if !ok {
	// 							node.AddChild(baseURL + link)				// nambahin semua anakan dari node
	// 							newLastNodes.Add(node.Children[baseURL+link])
	// 							visitedURL.Add(baseURL + link, true)
	// 						} else {
	// 							// fmt.Println("Already visited")
	// 						}
								
	// 					}
	// 				})
	// 			}
	// 		}(node)
	// 	}

	// 	wg.Wait()

	// 	currentLastNodes.Set(newLastNodes.Get())

	// 	fmt.Println("Recursing")
	// 	paths, visitedCount = root.DFS(baseURL, endURL)
	// 	fmt.Println("Finished Recursing")

	// 	depth++
	// 	fmt.Println("Depth:", depth)
	// }

	// fmt.Println(paths)
	// fmt.Println(visitedCount)
	// fmt.Println(countArticle)
	fmt.Println("Checked articles from DFS func:", visitedCount)
	visitedCount = len(visitedURL.data)

	return paths, visitedCount
}