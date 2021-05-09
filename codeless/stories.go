package codeless

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Story struct {
	CaseNumber string
	Title      string
	Text       string
}

var stories []Story
var storiesById map[int]Story
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

	storiesById = make(map[int]Story)

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

func GetStoryById(id int) Story {
	return storiesById[id]
}

func DecorateStory(story Story) string {
	return fmt.Sprintf("# %s: %s\n\n%s", story.CaseNumber, story.Title, story.Text)
}

func GetAllStories() []Story {
	var sortedStories []Story

	for i := 0; i < len(storyIds); i++ {
		story := storiesById[i]
		if story.CaseNumber == "" {
			continue
		}
		sortedStories = append(sortedStories, story)
	}

	return sortedStories
}
