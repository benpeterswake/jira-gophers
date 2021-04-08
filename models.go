package jira

import (
	"time"

	"github.com/trivago/tgo/tcontainer"
)

// SearchOptions specifies the optional parameters to various List methods that
// support pagination.
// Pagination is used for the Jira REST APIs to conserve server resources and limit
// response size for resources that return potentially large collection of items.
// A request to a pages API will result in a values array wrapped in a JSON object with some paging metadata
// Default Pagination options
type SearchOptions struct {
	// StartAt: The starting index of the returned projects. Base index: 0.
	StartAt int `url:"startAt,omitempty"`
	// MaxResults: The maximum number of projects to return per page. Default: 50.
	MaxResults int `url:"maxResults,omitempty"`
	// Expand: Expand specific sections in the returned issues
	Expand string `url:"expand,omitempty"`
	Fields []string
	// ValidateQuery: The validateQuery param offers control over whether to validate and how strictly to treat the validation. Default: strict.
	ValidateQuery string `url:"validateQuery,omitempty"`
}

// ChangelogItems reflects one single changelog item of a history item
type ChangelogItems struct {
	Field      string      `json:"field" structs:"field"`
	FieldType  string      `json:"fieldtype" structs:"fieldtype"`
	From       interface{} `json:"from" structs:"from"`
	FromString string      `json:"fromString" structs:"fromString"`
	To         interface{} `json:"to" structs:"to"`
	ToString   string      `json:"toString" structs:"toString"`
}

// ChangelogHistory reflects one single changelog history entry
type ChangelogHistory struct {
	Id      string           `json:"id" structs:"id"`
	Author  User             `json:"author" structs:"author"`
	Created string           `json:"created" structs:"created"`
	Items   []ChangelogItems `json:"items" structs:"items"`
}

// Changelog reflects the change log of an issue
type Changelog struct {
	Histories []ChangelogHistory `json:"histories,omitempty"`
}

// Attachment represents a Jira attachment
type Attachment struct {
	Self      string `json:"self,omitempty" structs:"self,omitempty"`
	ID        string `json:"id,omitempty" structs:"id,omitempty"`
	Filename  string `json:"filename,omitempty" structs:"filename,omitempty"`
	Author    *User  `json:"author,omitempty" structs:"author,omitempty"`
	Created   string `json:"created,omitempty" structs:"created,omitempty"`
	Size      int    `json:"size,omitempty" structs:"size,omitempty"`
	MimeType  string `json:"mimeType,omitempty" structs:"mimeType,omitempty"`
	Content   string `json:"content,omitempty" structs:"content,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty" structs:"thumbnail,omitempty"`
}

// Epic represents the epic to which an issue is associated
// Not that this struct does not process the returned "color" value
type Epic struct {
	ID      int    `json:"id" structs:"id"`
	Key     string `json:"key" structs:"key"`
	Self    string `json:"self" structs:"self"`
	Name    string `json:"name" structs:"name"`
	Summary string `json:"summary" structs:"summary"`
	Done    bool   `json:"done" structs:"done"`
}

// IssueFields represents single fields of a Jira issue.
// Every Jira issue has several fields attached.
type IssueFields struct {
	// TODO Missing fields
	//      * "workratio": -1,
	//      * "lastViewed": null,
	//      * "environment": null,
	Expand                        string            `json:"expand,omitempty" structs:"expand,omitempty"`
	Type                          IssueType         `json:"issuetype,omitempty" structs:"issuetype,omitempty"`
	Project                       Project           `json:"project,omitempty" structs:"project,omitempty"`
	Resolution                    *Resolution       `json:"resolution,omitempty" structs:"resolution,omitempty"`
	Priority                      *Priority         `json:"priority,omitempty" structs:"priority,omitempty"`
	Resolutiondate                Time              `json:"resolutiondate,omitempty" structs:"resolutiondate,omitempty"`
	Created                       Time              `json:"created,omitempty" structs:"created,omitempty"`
	Duedate                       Date              `json:"duedate,omitempty" structs:"duedate,omitempty"`
	Watches                       *Watches          `json:"watches,omitempty" structs:"watches,omitempty"`
	Assignee                      *User             `json:"assignee,omitempty" structs:"assignee,omitempty"`
	Updated                       Time              `json:"updated,omitempty" structs:"updated,omitempty"`
	Description                   string            `json:"description,omitempty" structs:"description,omitempty"`
	Summary                       string            `json:"summary,omitempty" structs:"summary,omitempty"`
	Creator                       *User             `json:"Creator,omitempty" structs:"Creator,omitempty"`
	Reporter                      *User             `json:"reporter,omitempty" structs:"reporter,omitempty"`
	Components                    []*Component      `json:"components,omitempty" structs:"components,omitempty"`
	Status                        *Status           `json:"status,omitempty" structs:"status,omitempty"`
	Progress                      *Progress         `json:"progress,omitempty" structs:"progress,omitempty"`
	AggregateProgress             *Progress         `json:"aggregateprogress,omitempty" structs:"aggregateprogress,omitempty"`
	TimeTracking                  *TimeTracking     `json:"timetracking,omitempty" structs:"timetracking,omitempty"`
	TimeSpent                     int               `json:"timespent,omitempty" structs:"timespent,omitempty"`
	TimeEstimate                  int               `json:"timeestimate,omitempty" structs:"timeestimate,omitempty"`
	TimeOriginalEstimate          int               `json:"timeoriginalestimate,omitempty" structs:"timeoriginalestimate,omitempty"`
	Worklog                       *Worklog          `json:"worklog,omitempty" structs:"worklog,omitempty"`
	IssueLinks                    []*IssueLink      `json:"issuelinks,omitempty" structs:"issuelinks,omitempty"`
	Comments                      *Comments         `json:"comment,omitempty" structs:"comment,omitempty"`
	FixVersions                   []*FixVersion     `json:"fixVersions,omitempty" structs:"fixVersions,omitempty"`
	AffectsVersions               []*AffectsVersion `json:"versions,omitempty" structs:"versions,omitempty"`
	Labels                        []string          `json:"labels,omitempty" structs:"labels,omitempty"`
	Subtasks                      []*Subtasks       `json:"subtasks,omitempty" structs:"subtasks,omitempty"`
	Attachments                   []*Attachment     `json:"attachment,omitempty" structs:"attachment,omitempty"`
	Epic                          *Epic             `json:"epic,omitempty" structs:"epic,omitempty"`
	Sprint                        *Sprint           `json:"sprint,omitempty" structs:"sprint,omitempty"`
	Parent                        *Parent           `json:"parent,omitempty" structs:"parent,omitempty"`
	AggregateTimeOriginalEstimate int               `json:"aggregatetimeoriginalestimate,omitempty" structs:"aggregatetimeoriginalestimate,omitempty"`
	AggregateTimeSpent            int               `json:"aggregatetimespent,omitempty" structs:"aggregatetimespent,omitempty"`
	AggregateTimeEstimate         int               `json:"aggregatetimeestimate,omitempty" structs:"aggregatetimeestimate,omitempty"`
	Unknowns                      tcontainer.MarshalMap
}

// Parent represents the parent of a Jira issue, to be used with subtask issue types.
type Parent struct {
	ID  string `json:"id,omitempty" structs:"id"`
	Key string `json:"key,omitempty" structs:"key"`
}

// Subtasks represents all issues of a parent issue.
type Subtasks struct {
	ID     string      `json:"id" structs:"id"`
	Key    string      `json:"key" structs:"key"`
	Self   string      `json:"self" structs:"self"`
	Fields IssueFields `json:"fields" structs:"fields"`
}

// AffectsVersion represents a software release which is affected by an issue.
type AffectsVersion Version

// FixVersion represents a software release in which an issue is fixed.
type FixVersion struct {
	Self            string `json:"self,omitempty" structs:"self,omitempty"`
	ID              string `json:"id,omitempty" structs:"id,omitempty"`
	Name            string `json:"name,omitempty" structs:"name,omitempty"`
	Description     string `json:"description,omitempty" structs:"description,omitempty"`
	Archived        *bool  `json:"archived,omitempty" structs:"archived,omitempty"`
	Released        *bool  `json:"released,omitempty" structs:"released,omitempty"`
	ReleaseDate     string `json:"releaseDate,omitempty" structs:"releaseDate,omitempty"`
	UserReleaseDate string `json:"userReleaseDate,omitempty" structs:"userReleaseDate,omitempty"`
	ProjectID       int    `json:"projectId,omitempty" structs:"projectId,omitempty"` // Unlike other IDs, this is returned as a number
	StartDate       string `json:"startDate,omitempty" structs:"startDate,omitempty"`
}

// IssueLink represents a link between two issues in Jira.
type IssueLink struct {
	ID           string        `json:"id,omitempty" structs:"id,omitempty"`
	Self         string        `json:"self,omitempty" structs:"self,omitempty"`
	Type         IssueLinkType `json:"type" structs:"type"`
	OutwardIssue *Issue        `json:"outwardIssue" structs:"outwardIssue"`
	InwardIssue  *Issue        `json:"inwardIssue" structs:"inwardIssue"`
	Comment      *Comment      `json:"comment,omitempty" structs:"comment,omitempty"`
}

// IssueLinkType represents a type of a link between to issues in Jira.
// Typical issue link types are "Related to", "Duplicate", "Is blocked by", etc.
type IssueLinkType struct {
	ID      string `json:"id,omitempty" structs:"id,omitempty"`
	Self    string `json:"self,omitempty" structs:"self,omitempty"`
	Name    string `json:"name" structs:"name"`
	Inward  string `json:"inward" structs:"inward"`
	Outward string `json:"outward" structs:"outward"`
}

// Progress represents the progress of a Jira issue.
type Progress struct {
	Progress int `json:"progress" structs:"progress"`
	Total    int `json:"total" structs:"total"`
	Percent  int `json:"percent" structs:"percent"`
}

// Component represents a "component" of a Jira issue.
// Components can be user defined in every Jira instance.
type Component struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
}

// Watches represents a type of how many and which user are "observing" a Jira issue to track the status / updates.
type Watches struct {
	Self       string     `json:"self,omitempty" structs:"self,omitempty"`
	WatchCount int        `json:"watchCount,omitempty" structs:"watchCount,omitempty"`
	IsWatching bool       `json:"isWatching,omitempty" structs:"isWatching,omitempty"`
	Watchers   []*Watcher `json:"watchers,omitempty" structs:"watchers,omitempty"`
}

// Watcher represents a simplified user that "observes" the issue
type Watcher struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	AccountID   string `json:"accountId,omitempty" structs:"accountId,omitempty"`
	DisplayName string `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active      bool   `json:"active,omitempty" structs:"active,omitempty"`
}

// Time represents the Time definition of Jira as a time.Time of go
type Time time.Time

func (t Time) Equal(u Time) bool {
	return time.Time(t).Equal(time.Time(u))
}

// Date represents the Date definition of Jira as a time.Time of go
type Date time.Time

// UnmarshalJSON will transform the Jira time into a time.Time
// during the transformation of the Jira JSON response
func (t *Time) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the main JSON package.
	if string(b) == "null" {
		return nil
	}
	ti, err := time.Parse("\"2006-01-02T15:04:05.999-0700\"", string(b))
	if err != nil {
		return err
	}
	*t = Time(ti)
	return nil
}

// MarshalJSON will transform the time.Time into a Jira time
// during the creation of a Jira request
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format("\"2006-01-02T15:04:05.000-0700\"")), nil
}

// UnmarshalJSON will transform the Jira date into a time.Time
// during the transformation of the Jira JSON response
func (t *Date) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the main JSON package.
	if string(b) == "null" {
		return nil
	}
	ti, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return err
	}
	*t = Date(ti)
	return nil
}

// MarshalJSON will transform the Date object into a short
// date string as Jira expects during the creation of a
// Jira request
func (t Date) MarshalJSON() ([]byte, error) {
	time := time.Time(t)
	return []byte(time.Format("\"2006-01-02\"")), nil
}

// Priority represents a priority of a Jira issue.
// Typical types are "Normal", "Moderate", "Urgent", ...
type Priority struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	StatusColor string `json:"statusColor,omitempty" structs:"statusColor,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
}

// TimeTracking represents the timetracking fields of a Jira issue.
type TimeTracking struct {
	OriginalEstimate         string `json:"originalEstimate,omitempty" structs:"originalEstimate,omitempty"`
	RemainingEstimate        string `json:"remainingEstimate,omitempty" structs:"remainingEstimate,omitempty"`
	TimeSpent                string `json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
	OriginalEstimateSeconds  int    `json:"originalEstimateSeconds,omitempty" structs:"originalEstimateSeconds,omitempty"`
	RemainingEstimateSeconds int    `json:"remainingEstimateSeconds,omitempty" structs:"remainingEstimateSeconds,omitempty"`
	TimeSpentSeconds         int    `json:"timeSpentSeconds,omitempty" structs:"timeSpentSeconds,omitempty"`
}

// WorklogRecord represents one entry of a Worklog
type WorklogRecord struct {
	Self             string           `json:"self,omitempty" structs:"self,omitempty"`
	Author           *User            `json:"author,omitempty" structs:"author,omitempty"`
	UpdateAuthor     *User            `json:"updateAuthor,omitempty" structs:"updateAuthor,omitempty"`
	Comment          string           `json:"comment,omitempty" structs:"comment,omitempty"`
	Created          *Time            `json:"created,omitempty" structs:"created,omitempty"`
	Updated          *Time            `json:"updated,omitempty" structs:"updated,omitempty"`
	Started          *Time            `json:"started,omitempty" structs:"started,omitempty"`
	TimeSpent        string           `json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
	TimeSpentSeconds int              `json:"timeSpentSeconds,omitempty" structs:"timeSpentSeconds,omitempty"`
	ID               string           `json:"id,omitempty" structs:"id,omitempty"`
	IssueID          string           `json:"issueId,omitempty" structs:"issueId,omitempty"`
	Properties       []EntityProperty `json:"properties,omitempty"`
}

type EntityProperty struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// Worklog represents the work log of a Jira issue.
// One Worklog contains zero or n WorklogRecords
// Jira Wiki: https://confluence.atlassian.com/jira/logging-work-on-an-issue-185729605.html
type Worklog struct {
	StartAt    int             `json:"startAt" structs:"startAt"`
	MaxResults int             `json:"maxResults" structs:"maxResults"`
	Total      int             `json:"total" structs:"total"`
	Worklogs   []WorklogRecord `json:"worklogs" structs:"worklogs"`
}

// Resolution represents a resolution of a Jira issue.
// Typical types are "Fixed", "Suspended", "Won't Fix", ...
type Resolution struct {
	Self        string `json:"self" structs:"self"`
	ID          string `json:"id" structs:"id"`
	Description string `json:"description" structs:"description"`
	Name        string `json:"name" structs:"name"`
}

// ProjectCategory represents a single project category
type ProjectCategory struct {
	Self        string `json:"self" structs:"self,omitempty"`
	ID          string `json:"id" structs:"id,omitempty"`
	Name        string `json:"name" structs:"name,omitempty"`
	Description string `json:"description" structs:"description,omitempty"`
}

// Project represents a Jira Project.
type Project struct {
	Expand          string             `json:"expand,omitempty" structs:"expand,omitempty"`
	Self            string             `json:"self,omitempty" structs:"self,omitempty"`
	ID              string             `json:"id,omitempty" structs:"id,omitempty"`
	Key             string             `json:"key,omitempty" structs:"key,omitempty"`
	Description     string             `json:"description,omitempty" structs:"description,omitempty"`
	Lead            User               `json:"lead,omitempty" structs:"lead,omitempty"`
	Components      []ProjectComponent `json:"components,omitempty" structs:"components,omitempty"`
	IssueTypes      []IssueType        `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
	URL             string             `json:"url,omitempty" structs:"url,omitempty"`
	Email           string             `json:"email,omitempty" structs:"email,omitempty"`
	AssigneeType    string             `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
	Versions        []Version          `json:"versions,omitempty" structs:"versions,omitempty"`
	Name            string             `json:"name,omitempty" structs:"name,omitempty"`
	Roles           map[string]string  `json:"roles,omitempty" structs:"roles,omitempty"`
	AvatarUrls      AvatarUrls         `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	ProjectCategory ProjectCategory    `json:"projectCategory,omitempty" structs:"projectCategory,omitempty"`
}

// Version represents a single release version of a project
type Version struct {
	Self            string `json:"self,omitempty" structs:"self,omitempty"`
	ID              string `json:"id,omitempty" structs:"id,omitempty"`
	Name            string `json:"name,omitempty" structs:"name,omitempty"`
	Description     string `json:"description,omitempty" structs:"description,omitempty"`
	Archived        *bool  `json:"archived,omitempty" structs:"archived,omitempty"`
	Released        *bool  `json:"released,omitempty" structs:"released,omitempty"`
	ReleaseDate     string `json:"releaseDate,omitempty" structs:"releaseDate,omitempty"`
	UserReleaseDate string `json:"userReleaseDate,omitempty" structs:"userReleaseDate,omitempty"`
	ProjectID       int    `json:"projectId,omitempty" structs:"projectId,omitempty"` // Unlike other IDs, this is returned as a number
	StartDate       string `json:"startDate,omitempty" structs:"startDate,omitempty"`
}

// ProjectComponent represents a single component of a project
type ProjectComponent struct {
	Self                string `json:"self" structs:"self,omitempty"`
	ID                  string `json:"id" structs:"id,omitempty"`
	Name                string `json:"name" structs:"name,omitempty"`
	Description         string `json:"description" structs:"description,omitempty"`
	Lead                User   `json:"lead,omitempty" structs:"lead,omitempty"`
	AssigneeType        string `json:"assigneeType" structs:"assigneeType,omitempty"`
	Assignee            User   `json:"assignee" structs:"assignee,omitempty"`
	RealAssigneeType    string `json:"realAssigneeType" structs:"realAssigneeType,omitempty"`
	RealAssignee        User   `json:"realAssignee" structs:"realAssignee,omitempty"`
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid" structs:"isAssigneeTypeValid,omitempty"`
	Project             string `json:"project" structs:"project,omitempty"`
	ProjectID           int    `json:"projectId" structs:"projectId,omitempty"`
}

// Pe

// Transition represents an issue transition in Jira
type Transition struct {
	ID     string                     `json:"id" structs:"id"`
	Name   string                     `json:"name" structs:"name"`
	To     Status                     `json:"to" structs:"status"`
	Fields map[string]TransitionField `json:"fields" structs:"fields"`
}

// StatusCategory represents the category a status belongs to.
// Those categories can be user defined in every Jira instance.
type StatusCategory struct {
	Self      string `json:"self" structs:"self"`
	ID        int    `json:"id" structs:"id"`
	Name      string `json:"name" structs:"name"`
	Key       string `json:"key" structs:"key"`
	ColorName string `json:"colorName" structs:"colorName"`
}

// Status represents the current status of a Jira issue.
// Typical status are "Open", "In Progress", "Closed", ...
// Status can be user defined in every Jira instance.
type Status struct {
	Self           string         `json:"self" structs:"self"`
	Description    string         `json:"description" structs:"description"`
	IconURL        string         `json:"iconUrl" structs:"iconUrl"`
	Name           string         `json:"name" structs:"name"`
	ID             string         `json:"id" structs:"id"`
	StatusCategory StatusCategory `json:"statusCategory" structs:"statusCategory"`
}

// TransitionField represents the value of one Transition
type TransitionField struct {
	Required bool `json:"required" structs:"required"`
}

// Issue represents a Jira issue.
type Issue struct {
	Expand         string               `json:"expand,omitempty" structs:"expand,omitempty"`
	ID             string               `json:"id,omitempty" structs:"id,omitempty"`
	Self           string               `json:"self,omitempty" structs:"self,omitempty"`
	Key            string               `json:"key,omitempty" structs:"key,omitempty"`
	Fields         *IssueFields         `json:"fields,omitempty" structs:"fields,omitempty"`
	RenderedFields *IssueRenderedFields `json:"renderedFields,omitempty" structs:"renderedFields,omitempty"`
	Changelog      *Changelog           `json:"changelog,omitempty" structs:"changelog,omitempty"`
	Transitions    []Transition         `json:"transitions,omitempty" structs:"transitions,omitempty"`
	Names          map[string]string    `json:"names,omitempty" structs:"names,omitempty"`
}

// IssueRenderedFields represents rendered fields of a Jira issue.
// Not all IssueFields are rendered.
type IssueRenderedFields struct {
	// TODO Missing fields
	//      * "aggregatetimespent": null,
	//      * "workratio": -1,
	//      * "lastViewed": null,
	//      * "aggregatetimeoriginalestimate": null,
	//      * "aggregatetimeestimate": null,
	//      * "environment": null,
	Resolutiondate string    `json:"resolutiondate,omitempty" structs:"resolutiondate,omitempty"`
	Created        string    `json:"created,omitempty" structs:"created,omitempty"`
	Duedate        string    `json:"duedate,omitempty" structs:"duedate,omitempty"`
	Updated        string    `json:"updated,omitempty" structs:"updated,omitempty"`
	Comments       *Comments `json:"comment,omitempty" structs:"comment,omitempty"`
	Description    string    `json:"description,omitempty" structs:"description,omitempty"`
}

// IssueType represents a type of a Jira issue.
// Typical types are "Request", "Bug", "Story", ...
type IssueType struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	Subtask     bool   `json:"subtask,omitempty" structs:"subtask,omitempty"`
	AvatarID    int    `json:"avatarId,omitempty" structs:"avatarId,omitempty"`
}

// Comments represents a list of Comment.
type Comments struct {
	Comments []*Comment `json:"comments,omitempty" structs:"comments,omitempty"`
}

// Comment represents a comment by a person to an issue in Jira.
type Comment struct {
	ID           string            `json:"id,omitempty" structs:"id,omitempty"`
	Self         string            `json:"self,omitempty" structs:"self,omitempty"`
	Name         string            `json:"name,omitempty" structs:"name,omitempty"`
	Author       User              `json:"author,omitempty" structs:"author,omitempty"`
	Body         string            `json:"body,omitempty" structs:"body,omitempty"`
	UpdateAuthor User              `json:"updateAuthor,omitempty" structs:"updateAuthor,omitempty"`
	Updated      string            `json:"updated,omitempty" structs:"updated,omitempty"`
	Created      string            `json:"created,omitempty" structs:"created,omitempty"`
	Visibility   CommentVisibility `json:"visibility,omitempty" structs:"visibility,omitempty"`
}

// CommentVisibility represents he visibility of a comment.
// E.g. Type could be "role" and Value "Administrators"
type CommentVisibility struct {
	Type  string `json:"type,omitempty" structs:"type,omitempty"`
	Value string `json:"value,omitempty" structs:"value,omitempty"`
}

// Sprint represents a sprint on Jira agile board
type Sprint struct {
	ID            int        `json:"id" structs:"id"`
	Name          string     `json:"name" structs:"name"`
	CompleteDate  *time.Time `json:"completeDate" structs:"completeDate"`
	EndDate       *time.Time `json:"endDate" structs:"endDate"`
	StartDate     *time.Time `json:"startDate" structs:"startDate"`
	OriginBoardID int        `json:"originBoardId" structs:"originBoardId"`
	Self          string     `json:"self" structs:"self"`
	State         string     `json:"state" structs:"state"`
}

// searchResult is only a small wrapper around the Search (with JQL) method
// to be able to parse the results
type searchResult struct {
	Issues     []Issue `json:"issues" structs:"issues"`
	StartAt    int     `json:"startAt" structs:"startAt"`
	MaxResults int     `json:"maxResults" structs:"maxResults"`
	Total      int     `json:"total" structs:"total"`
}

// User represents a Jira user.
type User struct {
	Self            string     `json:"self,omitempty" structs:"self,omitempty"`
	AccountID       string     `json:"accountId,omitempty" structs:"accountId,omitempty"`
	AccountType     string     `json:"accountType,omitempty" structs:"accountType,omitempty"`
	Name            string     `json:"name,omitempty" structs:"name,omitempty"`
	Key             string     `json:"key,omitempty" structs:"key,omitempty"`
	Password        string     `json:"-"`
	EmailAddress    string     `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	AvatarUrls      AvatarUrls `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	DisplayName     string     `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active          bool       `json:"active,omitempty" structs:"active,omitempty"`
	TimeZone        string     `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Locale          string     `json:"locale,omitempty" structs:"locale,omitempty"`
	ApplicationKeys []string   `json:"applicationKeys,omitempty" structs:"applicationKeys,omitempty"`
}

// AvatarUrls represents different dimensions of avatars / images
type AvatarUrls struct {
	Four8X48  string `json:"48x48,omitempty" structs:"48x48,omitempty"`
	Two4X24   string `json:"24x24,omitempty" structs:"24x24,omitempty"`
	One6X16   string `json:"16x16,omitempty" structs:"16x16,omitempty"`
	Three2X32 string `json:"32x32,omitempty" structs:"32x32,omitempty"`
}

type Error struct {
	StatusCode int
	Body       map[string]interface{}
}

type IssueTimeTrackingUpdate struct {
	TimeTracking []TimeTrackingUpdate `json:"timetracking"`
}

type TimeTrackingUpdate struct {
	Edit TimeTrackingEdit `json:"edit"`
}

type TimeTrackingEdit struct {
	OriginalEstimate string`json:"originalEstimate,omitempty" structs:"originalEstimate,omitempty"`
	RemainingEstimate string`json:"remainingEstimate,omitempty" structs:"remainingEstimate,omitempty"`
	TimeSpent string`json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
}

type IssueUpdate struct {
	Update IssueTimeTrackingUpdate `json:"update"`
}


