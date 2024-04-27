package main

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func bfs(startURL string, endURL string, baseURL string) ([][]string, int) {
	// Inisialisasi variabel penentu algoritma
	visitedURL := SafeMap[bool]{data: make(map[string]bool)}
	queriedURL := SafeMap[bool]{data: make(map[string]bool)}
	found := false

	var paths SafeArray[[]string]
	paths.Add([]string{startURL})

	var solutions SafeArray[[]string]

	depth := 1

	// Looping hingga ketemu
	for !found {
		var newPaths SafeArray[[]string]
		pathsSize := len(paths.Get())
		var wg sync.WaitGroup
		wg.Add(pathsSize)

		fmt.Println("Depth:", depth)
		depth++

		maxConcurrentRequests := 250

		sem := make(chan struct{}, maxConcurrentRequests)

		for i := 0; i < pathsSize; i++ {
			sem <- struct{}{}
			go func(i int) {
				defer func() { <-sem }()
				defer wg.Done()

				// Mendapatkan path
				p := paths.Get()[i]

				// Mendapatkan node terakhir
				node := p[len(p)-1]

				// Jika sudah ketemu, maka hentikan
				if _, ok := queriedURL.Get(node); ok {
					return
				}

				doc := makeRequest(node)

				if doc != nil {
					duplicateURL := make(map[string]bool)

					// Hanya mengecek pada bagian yang memiliki id bodyContent
					bodyContent := doc.Find("#bodyContent")

					bodyContent.Find("a").Each(func(_ int, s *goquery.Selection) {
						link, _ := s.Attr("href")
						matched, _ := regexp.MatchString("^/wiki/[^:]+$", link)

						// Jika link sesuai dan tidak duplikat
						if matched && !duplicateURL[baseURL+link] {
							duplicateURL[baseURL+link] = true

							// Jika link adalah tujuan
							if baseURL+link == endURL {
								fmt.Println("Found!")
								found = true
								solutions.Add(append(p, baseURL+link))
							}

							_, ok := visitedURL.Get(baseURL + link)
							if !ok {
								visitedURL.Add(baseURL+link, true)
							}

							// Jika belum ketemu, maka tambahkan path baru ke newPaths
							if !found {
								newPath := make([]string, len(p))
								copy(newPath, p)
								newPath = append(newPath, baseURL+link)
								newPaths.Add(newPath)
							}
						}
					})
					// Tandai node sudah diquery
					queriedURL.Add(node, true)
				}
			}(i)
		}

		wg.Wait()

		// Set paths dengan newPaths
		paths.Set(newPaths.Get())
	}

	return solutions.Get(), len(visitedURL.data)
}

func bfs_single(startURL string, endURL string, baseURL string) ([][]string, int) {
	// Inisialisasi variabel penentu algoritma
	visitedURL := SafeMap[bool]{data: make(map[string]bool)}
	queriedURL := SafeMap[bool]{data: make(map[string]bool)}
	found := false

	var paths SafeArray[[]string]
	paths.Add([]string{startURL})

	var solutions SafeArray[[]string]

	depth := 1

	// Looping hingga ketemu
	for !found {
		var newPaths SafeArray[[]string]
		pathsSize := len(paths.Get())
		var wg sync.WaitGroup
		wg.Add(pathsSize)

		fmt.Println("Depth:", depth)
		depth++

		maxConcurrentRequests := 250

		sem := make(chan struct{}, maxConcurrentRequests)

		for i := 0; i < pathsSize; i++ {
			sem <- struct{}{}
			go func(i int) {
				defer func() { <-sem }()
				defer wg.Done()

				// Mendapatkan path
				p := paths.Get()[i]

				// Mendapatkan node terakhir
				node := p[len(p)-1]

				// Jika sudah ketemu, maka hentikan
				if found {
					return
				}

				// Jika sudah diquery, maka skip
				if _, ok := queriedURL.Get(node); ok {
					fmt.Println("Already queried")
					return
				}

				doc := makeRequest(node)

				if doc != nil {
					duplicateURL := make(map[string]bool)

					// Hanya mengecek pada bagian yang memiliki id bodyContent
					bodyContent := doc.Find("#bodyContent")

					bodyContent.Find("a").Each(func(_ int, s *goquery.Selection) {
						link, _ := s.Attr("href")
						matched, _ := regexp.MatchString("^/wiki/[^:]+$", link)

						// Jika link sesuai dan tidak duplikat
						if matched && !duplicateURL[baseURL+link] {
							duplicateURL[baseURL+link] = true

							// Jika link adalah tujuan dan belum ditemukan
							if baseURL+link == endURL && !found {
								fmt.Println("Found!")
								found = true
								solutions.Add(append(p, baseURL+link))
							}

							_, ok := visitedURL.Get(baseURL + link)
							if !ok {
								visitedURL.Add(baseURL+link, true)
							}

							// Jika belum ketemu, maka tambahkan path baru ke newPaths
							if !found {
								newPath := make([]string, len(p))
								copy(newPath, p)
								newPath = append(newPath, baseURL+link)
								newPaths.Add(newPath)
							}
						}
					})
					// Tandai node sudah diquery
					queriedURL.Add(node, true)
				}
			}(i)
		}

		wg.Wait()

		// Set paths dengan newPaths
		paths.Set(newPaths.Get())
	}

	return solutions.Get(), len(visitedURL.data)
}
