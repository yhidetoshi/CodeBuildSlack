package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	CHANNEL  = "dev"
	USERNAME = "CodeBuild"
	SLACKURL = os.Getenv("SLACKURL")
)

type CodeBuildPhaseStatus string

type CodeBuildEventDetail struct {
	BuildStatus CodeBuildPhaseStatus `json:"build-status"`
	ProjectName string               `json:"project-name"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event events.CloudWatchEvent) {
	resInfo := &CodeBuildEventDetail{}

	err := json.Unmarshal(event.Detail, &resInfo)
	if err != nil {
		fmt.Println(err)
	}

	PostSlack(resInfo.ProjectName, string(resInfo.BuildStatus))
}

func checkStatus(status string) string {
	var color string

	if status == "SUCCEEDED" {
		color = "#00ff00"
	} else if status == "IN_PROGRESS" {
		color = "#0000ff"
	} else {
		color = "#dc143c"
	}
	return color
}

func PostSlack(pjtName string, status string) {

	statusColor := checkStatus(status)

	field1 := slack.Field{Title: "ProjectName", Value: pjtName}
	field2 := slack.Field{Title: "BuildStatus", Value: status}

	attachment := slack.Attachment{}
	attachment.AddField(field1).AddField(field2)
	color := statusColor
	attachment.Color = &color
	payload := slack.Payload{
		Username:    USERNAME,
		Channel:     CHANNEL,
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(SLACKURL, "", payload)
	if err != nil {
		os.Exit(1)
	}
}