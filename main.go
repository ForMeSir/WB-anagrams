package main

import (
	"fmt"
	"strings"
)

func main() {
	fad := []string{"пар", "раП", "пра", "реп", "огр", "пер", "afh1", "фарш", "fha1", "fha2"}
	gar := [][]byte{}
	for i := 0; i < len(fad); i++ {
		gar = append(gar, []byte(fad[i]))
	}
	//fmt.Println(fad[0][0]==fad[1][2])
	maper := AnagramMap(&gar)
	 mao:=maper
	 fmt.Println(*mao["пар"])
	fmt.Println(maper)
}

func AnagramMap(array *[][]byte) (map[string]*[]string) {
	mape := make(map[string][]string)
	for i := 0; i < len(*array); i++ {
		far := *array
		word := string(far[i][:])
		var k int
		word=strings.ToLower(word)
		for key := range mape {
			if IsAnagram([]byte(key), []byte(word)) {
				mape[key] = append(mape[key], word)
				k++
			}
		}
		if k == 0 {
			mape[word] = []string{word}
		}
	}
	for key, value := range mape {
		if len(value) == 1 {
			delete(mape, key)
		}
	}
	maperu:=make(map[string]*[]string)
	for key:=range mape{
		k:= mape[key]
		maperu[key]=&k
	}
	return maperu
}

func IsAnagram(wordOne []byte, wordTwo []byte) bool {
	if len(wordOne) != len(wordTwo) {
		return false
	}
	mapNew := make(map[byte]int)
	for _, value := range wordOne {
		mapNew[value]++
	}
	for _, value := range wordTwo {
		mapNew[value]--
	}
	for _, value := range mapNew {
		if value != 0 {
			return false
		}
	}
	return true
}
