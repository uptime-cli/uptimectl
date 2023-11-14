package betteruptime

import (
	"fmt"
	"net/http"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

func (c *client) GetStatusPageResources(id string) ([]StatusPageResource, error) {
	result := ListStatusPageResourcesResponse{}

	endpoint := fmt.Sprintf("%s/%s/%s/resources", contextmanager.APIEndpoint(), statusPageEndpoint, id)

	resp, err := c.rest.R().
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == http.StatusNotFound {
		fmt.Printf("status page not found\n")
		return nil, fmt.Errorf("status page not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("incorrect status response")
	}
	fmt.Printf("status page found\n")
	return result.Data, nil
}

type ListStatusPageResourcesResponse struct {
	Data       []StatusPageResource `json:"data"`
	Pagination Pagination           `json:"pagination"`
}

type StatusPageResource struct {
	Id         string                       `json:"id"`
	Type       string                       `json:"type"`
	Attributes StatusPageResourceAttributes `json:"attributes"`
}

type StatusPageResourceAttributes struct {
	StatusPageSectionId int    `json:"status_page_section_id" yaml:"status_page_section_id"`
	ResourceId          int    `json:"resource_id" yaml:"resource_id"`
	ResourceType        string `json:"resource_type" yaml:"resource_type"`
	History             bool   `json:"history" yaml:"history"`
	WidgetType          string `json:"widget_type" yaml:"widget_type"`
	PublicName          string `json:"public_name" yaml:"public_name"`
	Explanation         string `json:"explanation" yaml:"explanation"`
	Position            int    `json:"position" yaml:"position"`
}
