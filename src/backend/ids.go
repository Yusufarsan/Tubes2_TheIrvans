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
	visited := make(map[string]bool) // Map untuk nyimpen node yang udah dikunjungi
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
			fmt.Println("Found endURL when traversing!")
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			*paths = append(*paths, pathCopy)
		}
	}

	visited[n.Value] = false
}

func ids(startURL string, endURL string, baseURL string) ([][]string, int) {
	visitedURL := SafeMap[bool]{data: make(map[string]bool)}
	// Inisialisasi root sebagai startURL nya dan querriedURL sebagai map untuk nyimpen URL yang udah di query
	queriedURL := SafeMap[bool]{data: make(map[string]bool)} // Agak ga penting soalnya hampir ga kepake
	root := NewNode(startURL)

	// fmt.Println("Depth-First Search:")
	// Inisialisasi paths untuk nyimpen solusi dan visitedCount untuk ngitung berapa artikel yang diperiksa
	var paths = make([][]string, 0)
	var visitedCount int
	// var countArticle int

	// Inisialisasi currentLastNodes untuk nyimpen semua simpul pada satu layer tertentu (bergantung di depth mana sehingga nilainya akan berubah ubah)
	var currentLastNodes SafeArray[*Node]
	currentLastNodes.Add(root)

	// Inisialisasi depth untuk nyimpen data kedalaman yang bakal ditelusuri, ini akan di increment trs sampe ketemu
	depth := 0

	// Selama belom ketemu satupun solusi, bakalan di loop
	for len(paths) == 0 {
		var newLastNodes SafeArray[*Node] // Inisialisasi newLastNodes untuk memperbarui currentLastNodes (newLastNodes adalah semua simpul setelah layer nya currentLastNodes)

		var wg sync.WaitGroup
		wg.Add(len(currentLastNodes.Get()))

		maxConcurrentRequests := 250

		sem := make(chan struct{}, maxConcurrentRequests)

		for _, node := range currentLastNodes.array { // iterasi semua simpul pada layer tertentu (pada currentLastNodes)
			sem <- struct{}{}
			go func(node *Node) {
				defer func() { <-sem }()
				defer wg.Done()

				if _, ok := queriedURL.Get(node.Value); ok { // Kalo udah pernah di query, ga perlu di query lagi (Tapi jarang masuk ke kondisi ini makanya rada ga penting)
					fmt.Println("Already queried")
					return
				}

				doc := makeRequest(node.Value) // Nge scrap ke Wikipedia
				if doc != nil {
					duplicateURL := make(map[string]bool)

					bodyContent := doc.Find("#bodyContent")

					bodyContent.Find("a").Each(func(_ int, s *goquery.Selection) { // Menemukan semua link yang ada di suatu halaman Wikipedia
						link, _ := s.Attr("href")
						matched, _ := regexp.MatchString("^/wiki/[^:]+$", link)

						if matched && !duplicateURL[baseURL+link] { // Kalo link nya adalah link ke artikel Wikipedia lain dan belum pernah di cek
							duplicateURL[baseURL+link] = true
							// fmt.Println(baseURL + link)

							if endURL == baseURL+link {
								fmt.Println("Found endURL when querying!")
								node.AddChild(baseURL + link)
								newLastNodes.Add(node.Children[baseURL+link])
							}

							_, ok := visitedURL.Get(baseURL + link)
							if !ok {
								node.AddChild(baseURL + link)                 // nambahin semua anakan dari node
								newLastNodes.Add(node.Children[baseURL+link]) // Membentuk layer baru
								visitedURL.Add(baseURL+link, true)
							}
						}
					})
				}
			}(node)
		}

		wg.Wait()

		currentLastNodes.Set(newLastNodes.Get()) // Update currentLastNodes dengan newLastNodes, yaitu layer berikutnya

		fmt.Println("Recursing")
		paths, visitedCount = root.DFS(baseURL, endURL)
		fmt.Println("Finished Recursing")

		depth++ // Increment depth
		fmt.Println("Depth:", depth)
	}

	// fmt.Println(paths)
	// fmt.Println(visitedCount)
	// fmt.Println(countArticle)

	return paths, visitedCount
}

func (n *Node) DFSSingle(baseURL string, endURL string) ([][]string, int) {
	visited := make(map[string]bool) // Map untuk nyimpen node yang udah dikunjungi
	paths := [][]string{}
	// countArticle := 1

	// n.mu.Lock()
	// defer n.mu.Unlock()

	found := false
	n.dfsRecursiveSingle(baseURL, endURL, visited, []string{}, &paths, &found)
	// fmt.Println(visited)
	return paths, len(visited)
}

func (n *Node) dfsRecursiveSingle(baseURL string, endURL string, visited map[string]bool, currentPath []string, paths *[][]string, found *bool) {
	// If node is nil, stop recursion
	if n == nil {
		return
	}

	visited[n.Value] = true
	currentPath = append(currentPath, n.Value)

	allVisited := true
	for _, child := range n.Children {
		if !visited[child.Value] {
			allVisited = false
			child.dfsRecursiveSingle(baseURL, endURL, visited, currentPath, paths, found)
			if *found {
				break // Stop recursing if a solution has been found
			}
		}
	}

	if allVisited && !*found {
		if n.Value == endURL {
			fmt.Println("Found endURL when traversing!")
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			*paths = append(*paths, pathCopy)
			*found = true // Mark that a solution has been found
		}
	}

	visited[n.Value] = false
}

func ids_single(startURL string, endURL string, baseURL string) ([][]string, int) {
	visitedURL := SafeMap[bool]{data: make(map[string]bool)}
	// Inisialisasi root sebagai startURL nya dan querriedURL sebagai map untuk nyimpen URL yang udah di query
	queriedURL := SafeMap[bool]{data: make(map[string]bool)} // Agak ga penting soalnya hampir ga kepake
	root := NewNode(startURL)

	// fmt.Println("Depth-First Search:")
	// Inisialisasi paths untuk nyimpen solusi dan visitedCount untuk ngitung berapa artikel yang diperiksa
	var paths = make([][]string, 0)
	var visitedCount int
	// var countArticle int

	// Inisialisasi currentLastNodes untuk nyimpen semua simpul pada satu layer tertentu (bergantung di depth mana sehingga nilainya akan berubah ubah)
	var currentLastNodes SafeArray[*Node]
	currentLastNodes.Add(root)

	// Inisialisasi depth untuk nyimpen data kedalaman yang bakal ditelusuri, ini akan di increment trs sampe ketemu
	depth := 0

	// Inisialisasi found untuk nyimpen apakah udah ketemu solusi atau belum
	found := false

	// Selama belom ketemu satupun solusi, bakalan di loop
	for len(paths) == 0 {
		var newLastNodes SafeArray[*Node] // Inisialisasi newLastNodes untuk memperbarui currentLastNodes (newLastNodes adalah semua simpul setelah layer nya currentLastNodes)

		var wg sync.WaitGroup
		wg.Add(len(currentLastNodes.Get()))

		maxConcurrentRequests := 250

		sem := make(chan struct{}, maxConcurrentRequests)

		for _, node := range currentLastNodes.array { // iterasi semua simpul pada layer tertentu (pada currentLastNodes)
			sem <- struct{}{}
			go func(node *Node) {
				defer func() { <-sem }()
				defer wg.Done()

				if found {
					return
				}

				if _, ok := queriedURL.Get(node.Value); ok { // Kalo udah pernah di query, ga perlu di query lagi (Tapi jarang masuk ke kondisi ini makanya rada ga penting)
					fmt.Println("Already queried")
					return
				}

				doc := makeRequest(node.Value) // Nge scrap ke Wikipedia
				if doc != nil {
					duplicateURL := make(map[string]bool)

					bodyContent := doc.Find("#bodyContent")

					bodyContent.Find("a").Each(func(_ int, s *goquery.Selection) { // Menemukan semua link yang ada di suatu halaman Wikipedia
						link, _ := s.Attr("href")
						matched, _ := regexp.MatchString("^/wiki/[^:]+$", link)

						if matched && !duplicateURL[baseURL+link] { // Kalo link nya adalah link ke artikel Wikipedia lain dan belum pernah di cek
							duplicateURL[baseURL+link] = true
							// fmt.Println(baseURL + link)

							if endURL == baseURL+link && !found {
								fmt.Println("Found endURL when querying!")
								node.AddChild(baseURL + link)
								newLastNodes.Add(node.Children[baseURL+link])
								found = true
							}

							_, ok := visitedURL.Get(baseURL + link)
							if !ok {
								node.AddChild(baseURL + link)                 // nambahin semua anakan dari node
								newLastNodes.Add(node.Children[baseURL+link]) // Membentuk layer baru
								visitedURL.Add(baseURL+link, true)
							}
						}
					})
				}
			}(node)
		}

		wg.Wait()

		currentLastNodes.Set(newLastNodes.Get()) // Update currentLastNodes dengan newLastNodes, yaitu layer berikutnya

		fmt.Println("Recursing")
		paths, visitedCount = root.DFSSingle(baseURL, endURL)
		fmt.Println("Finished Recursing")

		depth++ // Increment depth
		fmt.Println("Depth:", depth)
	}

	// fmt.Println(paths)
	// fmt.Println(visitedCount)
	// fmt.Println(countArticle)

	return paths, visitedCount
}
