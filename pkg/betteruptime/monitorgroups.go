package betteruptime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

var (
	monitorGroupEndpoint = "/api/v2/monitor-groups"
)

func (c *client) CreateMonitorGroup(name string) (*MonitorGroup, error) {
	result := GetMonitorGroupResponse{}

	resp, err := c.rest.R().
		SetResult(&result).
		SetBody(CreateMonitorGroupRequest{
			Name: name,
		}).
		Post(fmt.Sprintf("%s/%s", contextmanager.APIEndpoint(), monitorGroupEndpoint))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("incorrect status response")
	}
	return &result.Data, nil
}

func (c *client) DeleteMonitorGroup(id string) error {
	result := GetMonitorGroupResponse{}

	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), monitorGroupEndpoint, id)

	resp, err := c.rest.R().
		SetResult(&result).
		Delete(endpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() == http.StatusNotFound {
		fmt.Printf("monitor group not found\n")
		return nil
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("incorrect status response")
	}
	fmt.Printf("monitor group deleted\n")
	return nil
}

func (c *client) ListMonitoringGroups() ([]MonitorGroup, error) {
	monitorGroups := []MonitorGroup{}

	result := ListMonitorGroupResponse{}
	endpoint := fmt.Sprintf("%s/%s", contextmanager.APIEndpoint(), monitorGroupEndpoint)

	for {
		resp, err := c.rest.R().
			SetResult(&result).
			Get(endpoint)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode() != http.StatusOK {
			return nil, fmt.Errorf("incorrect status response")
		}

		monitorGroups = append(monitorGroups, result.Data...)

		if result.Pagination.Next == nil {
			break
		}
		if len(monitorGroups) > 50 {
			break
		}
		endpoint = *result.Pagination.Next
	}
	return monitorGroups, nil
}

type CreateMonitorGroupRequest struct {
	Name      string `json:"name"`
	Paused    bool   `json:"paused"`
	SortIndex *int   `json:"sort_index"`
}

type ListMonitorGroupResponse struct {
	Data       []MonitorGroup
	Pagination Pagination
}

type GetMonitorGroupResponse struct {
	Data MonitorGroup
}

type MonitorGroup struct {
	Id         string
	Type       string
	Attributes MonitorGroupAttributes
}

type MonitorGroupAttributes struct {
	Name      string     `json:"name"`
	SortIndex int        `json:"sort_index"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at"`
	Paused    bool       `json:"paused"`
}
