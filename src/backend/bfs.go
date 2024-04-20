package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SafeMap[T any] struct {
	mu   sync.Mutex
	data map[string]T
}

func (s *SafeMap[T]) Add(key string, value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *SafeMap[T]) Get(key string) (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.data[key]
	return value, ok
}

type SafeArray[T any] struct {
	mu    sync.Mutex
	array []T
}

func (s *SafeArray[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.array = append(s.array, value)
}

func (s *SafeArray[T]) Get() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.array
}

func (s *SafeArray[T]) Set(array []T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.array = array
}

func makeRequest(url string) *goquery.Document {
	res, err := http.Get(url)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	if err != nil {
		return nil
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil
	}

	return doc
}

func bfs(startURL string, endURL string, baseURL string) [][]string {
	visitedURL := SafeMap[bool]{data: make(map[string]bool)}
	correctURL := SafeMap[[]string]{data: make(map[string][]string)}
	queriedURL := SafeMap[bool]{data: make(map[string]bool)}
	found := false

	var paths SafeArray[[]string]
	paths.Add([]string{startURL})

	var solutions SafeArray[[]string]

	for !found {
		var newPaths SafeArray[[]string]
		pathsSize := len(paths.Get())
		var wg sync.WaitGroup
		wg.Add(pathsSize)

		maxConcurrentRequests := 250

		sem := make(chan struct{}, maxConcurrentRequests)

		for i := 0; i < pathsSize; i++ {
			sem <- struct{}{}
			go func(i int) {
				defer func() { <-sem }()
				defer wg.Done()
				p := paths.Get()[i]
				node := p[len(p)-1]

				if value, ok := correctURL.Get(node); ok {
					fmt.Println("Correct URL")
					for _, path := range value {
						newPath := append([]string{}, p...)
						newPath = append(newPath, path)
						newPaths.Add(newPath)
					}
					return
				}

				if _, ok := queriedURL.Get(node); ok {
					fmt.Println("Already queried")
					return
				}

				doc := makeRequest(node)

				if doc != nil {
					duplicateURL := make(map[string]bool)
					doc.Find("a").Each(func(_ int, s *goquery.Selection) {
						link, _ := s.Attr("href")
						matched, _ := regexp.MatchString("^/wiki/", link)

						if matched && !duplicateURL[baseURL+link] {
							duplicateURL[baseURL+link] = true

							if baseURL+link == endURL {
								fmt.Println("Found!")
								value, _ := correctURL.Get(node)
								correctURL.Add(node, append(value, baseURL+link))
								found = true
								solutions.Add(append(p, baseURL+link))
							}

							_, ok := visitedURL.Get(baseURL + link)
							if !found && !ok {
								visitedURL.Add(baseURL+link, true)
								newPath := make([]string, len(p))
								copy(newPath, p)
								newPath = append(newPath, baseURL+link)
								newPaths.Add(newPath)
							}
						}
					})
					queriedURL.Add(node, true)
				}
			}(i)
		}

		wg.Wait()
		paths.Set(newPaths.Get())
	}

	return solutions.Get()
}

func bfs_single(startURL string, endURL string, baseURL string) [][]string {
	visitedURL := SafeMap[bool]{data: make(map[string]bool)}
	correctURL := SafeMap[[]string]{data: make(map[string][]string)}
	queriedURL := SafeMap[bool]{data: make(map[string]bool)}

	solutions := SafeArray[[]string]{}

	found := make(chan struct{})
	stopSearch := make(chan struct{})
	var once sync.Once

	go func() {
		for {
			select {
			case <-stopSearch:
				return
			default:
				var paths SafeArray[[]string]
				paths.Add([]string{startURL})

				for len(paths.Get()) > 0 {
					var newPaths SafeArray[[]string]
					pathsSize := len(paths.Get())
					var wg sync.WaitGroup
					wg.Add(pathsSize)

					maxConcurrentRequests := 250
					sem := make(chan struct{}, maxConcurrentRequests)

					for i := 0; i < pathsSize; i++ {
						sem <- struct{}{}
						go func(i int) {
							defer func() { <-sem }()
							defer wg.Done()
							p := paths.Get()[i]
							node := p[len(p)-1]

							if value, ok := correctURL.Get(node); ok {
								fmt.Println("Correct URL")
								for _, path := range value {
									newPath := append([]string{}, p...)
									newPath = append(newPath, path)
									newPaths.Add(newPath)
								}
								return
							}

							if _, ok := queriedURL.Get(node); ok {
								fmt.Println("Already queried")
								return
							}

							doc := makeRequest(node)

							if doc != nil {
								duplicateURL := make(map[string]bool)
								doc.Find("a").Each(func(_ int, s *goquery.Selection) {
									link, _ := s.Attr("href")
									matched, _ := regexp.MatchString("^/wiki/", link)

									if matched && !duplicateURL[baseURL+link] {
										duplicateURL[baseURL+link] = true

										if baseURL+link == endURL {
											fmt.Println("Found!")
											value, _ := correctURL.Get(node)
											correctURL.Add(node, append(value, baseURL+link))
											solution := append(p, baseURL+link)
											solutions.Add(solution)
											once.Do(func() { close(found) })
										}

										_, ok := visitedURL.Get(baseURL + link)
										if !ok {
											visitedURL.Add(baseURL+link, true)
											newPath := make([]string, len(p))
											copy(newPath, p)
											newPath = append(newPath, baseURL+link)
											newPaths.Add(newPath)
										}
									}
								})
								queriedURL.Add(node, true)
							}
						}(i)
					}

					wg.Wait()
					paths.Set(newPaths.Get())
				}
			}
		}
	}()

	<-found
	close(stopSearch)
	return solutions.Get()
}
