/*package main

import (
	"fmt"
	"go_code/fund/crawl/links"
	"log"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	//breadthFirst(crawl, os.Args[1:])
	links.Extract("https://image.baidu.com/search/index?tn=baiduimage&word=黑丝")
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	dic := map[byte]int{
		'0': 1, '1': 1, '2': 1, '3': 1, '4': 1, '5': 1, '6': 1,
		'7': 1, '8': 1, '9': 1, 'A': 1, 'B': 1, 'C': 1, 'D': 1,
		'E': 1, 'F': 1, 'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1,
		'f': 1,
	}
	res := handler(dic)
	fmt.Println(res)
}

func handler(dic map[byte]int) string {
	s, s1, s2 := "", "", ""
	bytes1, bytes2 := make([]byte, 0), make([]byte, 0)
	fmt.Scanf("%s %s", &s1, &s2)
	s = s1 + s2
	if len(s) == 0 {
		return ""
	}
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if i%2 == 0 {
			bytes1 = append(bytes1, bytes[i])
		} else {
			bytes2 = append(bytes2, bytes[i])
		}
	}
	sort.Slice(bytes1, func(i, j int) bool {
		return bytes1[i] < bytes1[j]
	})
	sort.Slice(bytes2, func(i, j int) bool {
		return bytes2[i] < bytes2[j]
	})
	j, k := 0, 0
	for i := 0; i < len(bytes); i++ {
		if i%2 == 0 {
			bytes[i] = bytes1[j]
			j++
		} else {
			bytes[i] = bytes2[k]
			k++
		}
	}
	res := ""
	for _, c := range bytes {
		if dic[c] == 1 {
			var b string
			if c >= '0' && c <= '9' {
				b = fmt.Sprintf("%4b", int(c)-48)
			} else if c >= 'a' && c <= 'f' {
				b = fmt.Sprintf("%4b", int(c)-87)
			} else if c >= 'A' && c <= 'F' {
				b = fmt.Sprintf("%4b", int(c)-55)
			}
			sum := 0
			for i := len(b) - 1; i >= 0; i-- {
				if b[i] == '1' {
					sum = sum*2 + 1
				} else {
					sum = sum * 2
				}
			}
			if sum >= 0 && sum <= 9 {
				res += strconv.Itoa(sum)
			} else if sum >= 10 && sum <= 25 {
				res += string('A' + byte(sum-10))
			} else {
				res += string('a' + byte(sum-26))
			}
		} else {
			res += string(c)
		}
	}
	return res
}
