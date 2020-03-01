package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kpfaulkner/precious/models"
	"io/ioutil"
	"net/http"
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

func sendSlackMessage( title string, author string) error {

	fields := []models.Field{
		{
			Title: "Wiki Changed",
			Value: title,
			Short: false,
		},
	}

	msg := models.Webhook{
		UserName: "precious",
		Attachments: []models.Attachment{
			{
				AuthorName: author,
				Fields:     fields,
			},
		},
	}

	endpoint := os.Getenv("SLACK_WEBHOOK")
	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "URL is required")
		os.Exit(1)
	}

	err := sendToSlack(endpoint, msg)
	return err
}

func sendToSlack(endpoint string, msg models.Webhook) error {
	enc, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(enc)
	_, err = http.Post(endpoint, "application/json", b)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
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
	githubActor := os.Getenv("GITHUB_ACTOR")
	titleList := strings.Split(strings.ToLower(pageTitles), ",")

	for _,page := range ev.Pages {
		if contains(titleList, strings.ToLower(page.Title)) {
			sendSlackMessage(page.Title, githubActor)
			fmt.Printf("page title: %s , pagename: %s, change by: %s\n", page.Title, page.PageName, githubActor )
		}
  }
	//fmt.Println(fmt.Sprintf(`::set-output name=myOutput::%s`, output))
}
