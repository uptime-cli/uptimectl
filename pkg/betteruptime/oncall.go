package betteruptime

import (
	"fmt"
	"net/http"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

var (
	onCallsEndpoint = "/api/v2/on-calls"
)

func (c *client) GetOnCall() (*ListOnCallResponse, error) {
	result := ListOnCallResponse{}

	endpoint := fmt.Sprintf("%s/%s", contextmanager.APIEndpoint(), onCallsEndpoint)

	resp, err := c.rest.R().
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("on calls not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("incorrect status response")
	}
	return &result, nil
}

type ListOnCallResponse struct {
	Data       []OnCall         `json:"data"`
	Included   []UserReferences `json:"included"`
	Pagination Pagination       `json:"pagination"`
}

type OnCall struct {
	Id            string              `json:"id"`
	Type          string              `json:"type"`
	Attributes    OnCallAttributes    `json:"attributes"`
	Relationships OnCallRelationships `json:"relationships"`
}

type OnCallAttributes struct {
	Name            string `json:"name" yaml:"name"`
	DefaultCalendar bool   `json:"default_calendar" yaml:"default_calendar"`
}

type OnCallRelationships struct {
	OnCallUsers OnCallUserReferences `json:"on_call_users"`
}

type OnCallUserReferences struct {
	Data []struct {
		Id   string
		Type string
	}
}

type UserReferences struct {
	Id         string         `json:"id"`
	Attributes UserAttributes `json:"attributes"`
}
type UserAttributes struct {
	FirstName    string   `json:"first_name" yaml:"first_name"`
	LastName     string   `json:"last_name" yaml:"last_name"`
	Email        string   `json:"email" yaml:"email"`
	PhoneNumbers []string `json:"phone_numbers" yaml:"phone_numbers"`
}
