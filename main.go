package main

import (
	"encoding/json"
	"fmt"
	"github.com/kpfaulkner/precious/models"
	"io/ioutil"
	"os"
	"strings"
)

func contains(titleList []string, title string) bool {
	for _, t := range titleList {
		if t == title {
			return true
		}
	}
	return false
}

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
  fmt.Printf("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx\n")

	dat, err := ioutil.ReadFile(eventPath)
	if err != nil {
    fmt.Printf("unable to read event")
    return
	}

	// have the data, deserialise
  var ev models.GollumEventModel
	err = json.Unmarshal(dat, &ev)
	if err != nil {
		fmt.Printf("cannot unmarshal event data")
		return
	}

	pageTitles := os.Getenv("WIKI_TITLES_TO_ALERT")
	titleList := strings.Split(strings.ToLower(pageTitles), ",")

	for _,page := range ev.Pages {
		if contains(titleList, strings.ToLower(page.Title)) {
			fmt.Printf("page title: %s , pagename: %s, star: %d\n", page.Title, page.PageName, ev.Repository.StargazersCount)
		}
  }
	//fmt.Println(fmt.Sprintf(`::set-output name=myOutput::%s`, output))
}
