package todo

import (
	"fmt"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
)

// Issue represents a Todo issue
type Issue struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	CreateAt int64  `json:"create_at"`
	PostID   string `json:"post_id"`
}

// ExtendedIssue extends the information on Issue to be used on the front-end
type ExtendedIssue struct {
	Issue
	ForeignUser     string `json:"user"`
	ForeignList     string `json:"list"`
	ForeignPosition int    `json:"position"`
}

type ExtendedIssues struct {
	Issues []*ExtendedIssue
}

func newIssue(message string, postID string) *Issue {
	return &Issue{
		ID:       model.NewId(),
		CreateAt: model.GetMillis(),
		Message:  message,
		PostID:   postID,
	}
}

func (issues ExtendedIssues) ToString() string {
	if len(issues.Issues) == 0 {
		return "Nothing to do!"
	}

	str := "\n\n"

	for _, issue := range issues.Issues {
		createAt := time.Unix(issue.CreateAt/1000, 0)
		str += fmt.Sprintf("* %s\n  * (%s)\n", issue.Message, createAt.Format("January 2, 2006 at 15:04"))
	}

	return str
}