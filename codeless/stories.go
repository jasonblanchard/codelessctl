package codeless

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

var stories []story
var storiesById map[int]story
var storyIds []int

func init() {
	file, err := ioutil.ReadFile("stories.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &stories)

	if err != nil {
		panic(err)
	}

	storiesById = make(map[int]story)

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
		storyIds = append(storyIds, id)
		storiesById[id] = stories[i]
	}

	sort.Ints(storyIds)
}

func GetStoryById(id int) story {
	return storiesById[id]
}

func GetAllStories() []story {
	var sortedStories []story

	for i := 0; i < len(storyIds); i++ {
		story := storiesById[i]
		if story.CaseNumber == "" {
			continue
		}
		sortedStories = append(sortedStories, story)
	}

	return sortedStories
}
