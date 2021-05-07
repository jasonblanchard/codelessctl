package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type story struct {
	CaseNumber string
	Title      string
	Text       string
}

func main() {
	file, err := ioutil.ReadFile("../../stories.json")

	if err != nil {
		panic(err)
	}

	var stories []story
	err = json.Unmarshal(file, &stories)

	if err != nil {
		panic(err)
	}

	var ids []int
	storiesById := make(map[int]story)

	for i := 0; i < len(stories); i++ {
		if stories[i].CaseNumber == "" {
			continue
		}
		split := strings.Split(stories[i].CaseNumber, " ")
		idString := split[1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println(err)
		}
		ids = append(ids, id)
		storiesById[id] = stories[i]
	}

	sort.Ints(ids)
	fmt.Println(ids)
}
