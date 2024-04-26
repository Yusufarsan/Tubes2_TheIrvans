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

func (n *Node) DFS(endURL string) ([][]string, int) {
	visited := make(map[string]bool) // Map untuk nyimpen node yang udah dikunjungi
	paths := [][]string{}

	n.dfsIterative(endURL, visited, []string{}, &paths) // eak = visitedCount
	return paths, len(visited)
}

func (n *Node) dfsIterative(endURL string, visited map[string]bool, currentPath []string, paths *[][]string) {
	if n == nil {
		return
	}

	stack := []*Node{n}
	pathStack := [][]string{currentPath}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currentPath = pathStack[len(pathStack)-1]
		pathStack = pathStack[:len(pathStack)-1]

		visited[node.Value] = true
		currentPath = append(currentPath, node.Value)

		if node.Value == endURL {
			fmt.Println("Found endURL when traversing!")
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			*paths = append(*paths, pathCopy)
		}

		allVisited := true
		for _, child := range node.Children {
			if !visited[child.Value] {
				allVisited = false
				stack = append(stack, child)
				newPath := make([]string, len(currentPath))
				copy(newPath, currentPath)
				pathStack = append(pathStack, newPath)
			}
		}

		if allVisited {
			visited[node.Value] = false
		}
	}
}

func ids(startURL string, endURL string, baseURL string) ([][]string, int) {
	// Inisialisasi root sebagai startURL nya dan querriedURL sebagai map untuk nyimpen URL yang udah di query
	queriedURL := SafeMap[bool]{data: make(map[string]bool)} // Agak ga penting soalnya hampir ga kepake
	root := NewNode(startURL)

	// Inisialisasi paths untuk nyimpen solusi dan visitedCount untuk ngitung berapa artikel yang diperiksa
	var paths = make([][]string, 0)
	var visitedCount int

	// Inisialisasi currentLastNodes untuk nyimpen semua simpul pada satu layer tertentu (bergantung di depth mana sehingga nilainya akan berubah ubah)
	var currentLastNodes SafeArray[*Node]
	currentLastNodes.Add(root)

	currentLastNodesMap := SafeMap[bool]{data: make(map[string]bool)}
	currentLastNodesMap.Add(startURL, true)

	// Inisialisasi depth untuk nyimpen data kedalaman yang bakal ditelusuri, ini akan di increment trs sampe ketemu
	depth := 0

	// Selama belom ketemu satupun solusi, bakalan di loop
	for len(paths) == 0 {
		var newLastNodes SafeArray[*Node] // Inisialisasi newLastNodes untuk memperbarui currentLastNodes (newLastNodes adalah semua simpul setelah layer nya currentLastNodes)
		newLastNodesMap := SafeMap[bool]{data: make(map[string]bool)}

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

							if endURL == baseURL+link {
								fmt.Println("Found endURL when querying!")
							}

							// Only add node if it's not in currentLastNodes and hasnt been queried
							_, okNodeMap := currentLastNodesMap.Get(baseURL + link)
							_, okQuery := queriedURL.Get(baseURL + link)
							if !okNodeMap && !okQuery {
								node.AddChild(baseURL + link)                 // nambahin semua anakan dari node
								newLastNodes.Add(node.Children[baseURL+link]) // Membentuk layer baru
								newLastNodesMap.Add(baseURL+link, true)
							}
						}
					})
					queriedURL.Add(node.Value, true)
				}
			}(node)
		}

		wg.Wait()

		currentLastNodes.Set(newLastNodes.Get()) // Update currentLastNodes dengan newLastNodes, yaitu layer berikutnya
		currentLastNodesMap.Replace(newLastNodesMap.data)

		fmt.Println("Recursing")
		paths, visitedCount = root.DFS(endURL)
		fmt.Println("Finished Recursing")

		depth++ // Increment depth
		fmt.Println("Depth:", depth)
	}

	return paths, visitedCount
}

func (n *Node) DFSSingle(endURL string) ([][]string, int) {
	visited := make(map[string]bool) // Map untuk nyimpen node yang udah dikunjungi
	paths := [][]string{}

	n.dfsIterativeSingle(endURL, visited, []string{}, &paths)

	return paths, len(visited)
}

func (n *Node) dfsIterativeSingle(endURL string, visited map[string]bool, currentPath []string, paths *[][]string) {
	if n == nil {
		return
	}

	stack := []*Node{n}
	pathStack := [][]string{currentPath}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currentPath = pathStack[len(pathStack)-1]
		pathStack = pathStack[:len(pathStack)-1]

		visited[node.Value] = true
		currentPath = append(currentPath, node.Value)

		if node.Value == endURL {
			fmt.Println("Found endURL when traversing!")
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			*paths = append(*paths, pathCopy)
			break
		}

		allVisited := true
		for _, child := range node.Children {
			if !visited[child.Value] {
				allVisited = false
				stack = append(stack, child)
				newPath := make([]string, len(currentPath))
				copy(newPath, currentPath)
				pathStack = append(pathStack, newPath)
			}
		}

		if allVisited {
			visited[node.Value] = false
		}
	}
}

func ids_single(startURL string, endURL string, baseURL string) ([][]string, int) {
	// Inisialisasi root sebagai startURL nya dan querriedURL sebagai map untuk nyimpen URL yang udah di query
	queriedURL := SafeMap[bool]{data: make(map[string]bool)} // Agak ga penting soalnya hampir ga kepake
	root := NewNode(startURL)

	// Inisialisasi paths untuk nyimpen solusi dan visitedCount untuk ngitung berapa artikel yang diperiksa
	var paths = make([][]string, 0)
	var visitedCount int

	// Inisialisasi currentLastNodes untuk nyimpen semua simpul pada satu layer tertentu (bergantung di depth mana sehingga nilainya akan berubah ubah)
	var currentLastNodes SafeArray[*Node]
	currentLastNodes.Add(root)

	currentLastNodesMap := SafeMap[bool]{data: make(map[string]bool)}
	currentLastNodesMap.Add(startURL, true)

	// Inisialisasi depth untuk nyimpen data kedalaman yang bakal ditelusuri, ini akan di increment trs sampe ketemu
	depth := 0

	// Inisialisasi found untuk nyimpen apakah udah ketemu solusi atau belum
	found := false

	// Selama belom ketemu satupun solusi, bakalan di loop
	for len(paths) == 0 {
		var newLastNodes SafeArray[*Node] // Inisialisasi newLastNodes untuk memperbarui currentLastNodes (newLastNodes adalah semua simpul setelah layer nya currentLastNodes)
		newLastNodesMap := SafeMap[bool]{data: make(map[string]bool)}

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

							if endURL == baseURL+link && !found {
								fmt.Println("Found endURL when querying!")
								node.AddChild(baseURL + link)
								newLastNodes.Add(node.Children[baseURL+link])
								found = true
							}

							// Only add node if it's not in currentLastNodes and hasnt been queried
							_, okNodeMap := currentLastNodesMap.Get(baseURL + link)
							_, okQuery := queriedURL.Get(baseURL + link)
							if !okNodeMap && !okQuery && !found {
								node.AddChild(baseURL + link)                 // nambahin semua anakan dari node
								newLastNodes.Add(node.Children[baseURL+link]) // Membentuk layer baru
								newLastNodesMap.Add(baseURL+link, true)
							}
						}
					})
					queriedURL.Add(node.Value, true)
				}
			}(node)
		}

		wg.Wait()

		currentLastNodes.Set(newLastNodes.Get()) // Update currentLastNodes dengan newLastNodes, yaitu layer berikutnya
		currentLastNodesMap.Replace(newLastNodesMap.data)

		fmt.Println("Recursing")
		paths, visitedCount = root.DFSSingle(endURL)
		fmt.Println("Finished Recursing")

		depth++ // Increment depth
		fmt.Println("Depth:", depth)
	}

	return paths, visitedCount
}
