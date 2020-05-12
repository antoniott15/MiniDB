package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)


func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}


func mid(length int) int {
	if length%2 == 0 {
		return length / 2
	}

	return length/2 + 1
}

func getPositions(line string) []string {
	return SplitAny(line, "|")
}

func WalkDir(root string) (files []string) {
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	}); err != nil {
		return nil
	}

	return
}

func getPositionsAndPK(line string) ([]string, int) {
	pos := SplitAny(line, "|")
	for _, elements := range pos {
		pk, err := strconv.Atoi(elements)
		if err == nil {
			return pos, pk
		}
	}
	c := rand.Intn(1000000-0) + 10
	return pos, c
}


func toFormat(elem Structure) string {
	var finalWord string
	for i,_ := range elem.Headers {
		finalWord += fmt.Sprint(elem.Attribs[elem.Headers[i]]) + "|"
	}
	return finalWord[:len(finalWord)-1]
}

func headers(line []string) string {
	var word string
	for _, elements := range line{
		word += elements + "|"
	}
	return  word[:len(word)-1]
}