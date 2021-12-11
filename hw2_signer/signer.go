package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ExecutePipeline(jobs ...job) {
	wg := sync.WaitGroup{}
	wg.Add(len(jobs))

	var in chan interface{}
	for _, j := range jobs {
		out := make(chan interface{})
		go func(j job, in, out chan interface{}) {
			j(in, out)
			close(out)
			wg.Done()
		}(j, in, out)
		in = out
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	mu := sync.Mutex{}
	var results []chan interface{}

	for v := range in {
		results = append(results, make(chan interface{}))
		data := strconv.Itoa(v.(int))

		go func(data string, out chan interface{}) {
			md5ch := make(chan string)
			go func() {
				mu.Lock()
				md5ch <- DataSignerMd5(data)
				mu.Unlock()
			}()
			crc32ch := make(chan string)
			go func() {
				crc32ch <- DataSignerCrc32(data)
			}()
			crc32md5ch := make(chan string)
			go func() {
				crc32md5ch <- DataSignerCrc32(<-md5ch)
			}()
			result := <-crc32ch + "~" + <-crc32md5ch
			out <- result
		}(data, results[len(results)-1])
	}

	for i := range results {
		out <- <-results[i]
	}
}

func MultiHash(in, out chan interface{}) {
	var results []chan interface{}
	for v := range in {
		results = append(results, make(chan interface{}))
		data := v.(string)

		go func(data string, out chan interface{}) {
			var result string
			var crc32s []chan string
			for i := 0; i <= 5; i++ {
				crc32s = append(crc32s, make(chan string))
				go func(i int) {
					crc32s[i] <- DataSignerCrc32(strconv.Itoa(i) + data)
				}(i)
			}
			for i := 0; i <= 5; i++ {
				result += <-crc32s[i]
			}
			out <- result
		}(data, results[len(results)-1])
	}
	for i := range results {
		out <- <-results[i]
	}
}

func CombineResults(in, out chan interface{}) {
	var results []string
	for v := range in {
		results = append(results, v.(string))
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})
	dataRaw := strings.Join(results, "_")
	out <- dataRaw
}
