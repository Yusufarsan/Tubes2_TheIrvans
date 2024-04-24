package main

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Node struct {
	Value    string
	Children map[string]*Node
	mu       sync.Mutex
}

func NewNode(value string) *Node {
	return &Node{
		Value:    value,
		Children: make(map[string]*Node),
	}
}

func (n *Node) AddChild(childValue string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	child := NewNode(childValue)
	n.Children[childValue] = child
}

func (n *Node) DFS(baseURL string, endURL string) ([][]string, int) {
	visited := make(map[string]bool)		// Map untuk nyimpen node yang udah dikunjungi
	paths := [][]string{}
	// countArticle := 1

	// n.mu.Lock()
	// defer n.mu.Unlock()

	n.dfsRecursive(baseURL, endURL, visited, []string{}, &paths)
	// fmt.Println(visited)
	return paths, len(visited)
}

func (n *Node) dfsRecursive(baseURL string, endURL string, visited map[string]bool, currentPath []string, paths *[][]string) {
	// Kalo node nya nil, ga bakal nge cek anak anaknya lagi
	if n == nil {
		return
	}

	visited[n.Value] = true			
	// fmt.Println(n.Value)
	currentPath = append(currentPath, n.Value)

	allVisited := true
	for _, child := range n.Children {
		// n.mu.Lock()
		if !visited[child.Value] {
			allVisited = false
			// n.mu.Unlock()
			child.dfsRecursive(baseURL, endURL, visited, currentPath, paths)
		}
		// else {
		// 	n.mu.Unlock()
		// }
	}

	if allVisited {
		if n.Value == endURL {
			fmt.Println("Found!")
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			*paths = append(*paths, pathCopy)
		}
	}

	visited[n.Value] = false
}

func ids(startURL string, endURL string, baseURL string) ([][]string, int) {
	// visitedURL := SafeMap[bool]{data: make(map[string]bool)}
	queriedURL := SafeMap[bool]{data: make(map[string]bool)}
	root := NewNode(startURL)
	// root.AddChild("https://en.wikipedia.org/wiki/Indonesia")
	// root.AddChild("https://en.wikipedia.org/wiki/Java")
	// root.Children["https://en.wikipedia.org/wiki/Indonesia"].AddChild("https://en.wikipedia.org/wiki/Asia")
	// root.Children["https://en.wikipedia.org/wiki/Indonesia"].AddChild("https://en.wikipedia.org/wiki/Java")
	// root.Children["https://en.wikipedia.org/wiki/Java"].AddChild("https://en.wikipedia.org/wiki/Philosophy")
	// root.Children["https://en.wikipedia.org/wiki/Java"].AddChild("https://en.wikipedia.org/wiki/Asia")

	// fmt.Println("Depth-First Search:")
	var paths = make([][]string, 0)
	var visitedCount int
	// var countArticle int

	var currentLastNodes SafeArray[*Node]
	currentLastNodes.Add(root)

	depth := 0

	// Selama belom ketemu satupun solusi, bakalan di loop
	for len(paths) == 0 {
		var newLastNodes SafeArray[*Node]
		var wg sync.WaitGroup
		wg.Add(len(currentLastNodes.Get()))

		maxConcurrentRequests := 250

		sem := make(chan struct{}, maxConcurrentRequests)

		for _, node := range currentLastNodes.array {
			sem <- struct{}{}
			go func(node *Node) {
				defer func() { <-sem }()
				defer wg.Done()

				if _, ok := queriedURL.Get(node.Value); ok {
					fmt.Println("Already queried")
					return
				}

				doc := makeRequest(node.Value)
				if doc != nil {
					duplicateURL := make(map[string]bool)
					doc.Find("a").Each(func(_ int, s *goquery.Selection) {
						link, _ := s.Attr("href")
						matched, _ := regexp.MatchString("^/wiki/", link)

						if matched && !duplicateURL[baseURL+link] {
							duplicateURL[baseURL+link] = true
							// fmt.Println(baseURL + link)


							// _, ok := visitedURL.Get(baseURL + link)
							// if !ok {
								node.AddChild(baseURL + link)				// nambahin semua anakan dari node
								newLastNodes.Add(node.Children[baseURL+link])
							// 	visitedURL.Add(baseURL + link, true)
							// }
								
						}
					})
				}
			}(node)
		}

		wg.Wait()

		currentLastNodes.Set(newLastNodes.Get())

		fmt.Println("Recursing")
		paths, visitedCount = root.DFS(baseURL, endURL)
		fmt.Println("Finished Recursing")

		depth++
		fmt.Println("Depth:", depth)
	}

	// fmt.Println(paths)
	// fmt.Println(visitedCount)
	// fmt.Println(countArticle)
	// visitedCount = len(visitedURL.data)

	return paths, visitedCount
}

// func dfs () {
// 	///

// 	solusi := array of string
// 	for childnode {
// 		solusi append dfs()
// 	}

// 	retun solusi
// }

// func ids(startURL string, endURL string, baseURL string) ([][]string, int) {
// 	visitedURL := SafeMap[bool]{data: make(map[string]bool)}

// 	found := false
// 	depth := 1

// 	for !found {
// 		var newPaths SafeArray[[]string]
// 		pathsSize := len(paths.Get())
// 		var wg sync.WaitGroup
// 		wg.Add(pathsSize)

// 		depth++

// 		maxConcurrentRequests := 250
// 	}

// solution := dfs(startURL, endURL, baseURL, 0)

// Inisialisasi Root Node (startURL)
// makeRequest(startURL)
// Append root node dengan setiap link yang ada di dalamnya
// Traverse dengan DFS dengan batasan depth
// Jika ditemukan endURL, maka return path
// Jika tidak ditemukan, maka tambahkan depth dan ulangi proses

// }
