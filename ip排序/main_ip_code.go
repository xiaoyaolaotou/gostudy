package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"

	"strconv"
	"strings"
)

type IPAndStatusCode struct {
	ip   string
	url  string
	code int
}

type IPAndStatusCodeElement struct {
	key   IPAndStatusCode
	value int
}

func parse(line string) (string, string, int, error) {
	elements := strings.Fields(line)
	if len(elements) < 8 {
		return "", "", 0, fmt.Errorf("format error")
	}

	statusCode, err := strconv.Atoi(elements[8])
	if err != nil {
		return "", "", 0, err
	}

	return elements[0], elements[6], statusCode, nil
}

func main() {
	file, err := os.Open("access.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	render := bufio.NewReader(file)

	ipAndStatusCodeStats := make(map[IPAndStatusCode]int)

	for {
		ctx, _, err := render.ReadLine()

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if ip, url, statusCode, err := parse(string(ctx)); err == nil {
			key := IPAndStatusCode{ip, url, statusCode}
			ipAndStatusCodeStats[key]++
		}

	}
	s := make([]IPAndStatusCodeElement, 0)
	for k, v := range ipAndStatusCodeStats {
		s = append(s, IPAndStatusCodeElement{k, v})
	}

	// 排序
	sort.Slice(s, func(i, j int) bool {
		return s[i].value < s[j].value
	})
	fmt.Println(s)

}
