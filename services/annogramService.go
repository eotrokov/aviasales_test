package services

import (
	"aviasales/store"
	"sort"
	"strings"
)

func GetAnnogram(findWord string) []string {
	findWordSort := sortStr(findWord)
	var list []string
	//list := []string{}

	for _, word := range store.GetStore() {
		key := sortStr(word)
		if key == findWordSort {
			list = append(list, word)
		}
	}
	return list
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
