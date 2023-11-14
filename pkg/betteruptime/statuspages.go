package betteruptime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

var (
	statusPageEndpoint = "/api/v2/status-pages"
)

func (c *client) CreateStatusPage(createStatusPage CreateStatusPageRequest) (*StatusPage, error) {
	result := GetStatusPageResponse{}

	resp, err := c.rest.R().
		SetResult(&result).
		SetBody(createStatusPage).
		Post(fmt.Sprintf("%s/%s", contextmanager.APIEndpoint(), statusPageEndpoint))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("incorrect status response")
	}
	return &result.Data, nil
}

func (c *client) UpdateStatusPage(id string, createStatusPage CreateStatusPageRequest) (*StatusPage, error) {
	result := GetStatusPageResponse{}

	resp, err := c.rest.R().
		SetResult(&result).
		SetBody(createStatusPage).
		Patch(fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), statusPageEndpoint, id))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("incorrect status response")
	}
	return &result.Data, nil
}

func (c *client) GetStatusPage(id string) (*StatusPage, error) {
	result := GetStatusPageResponse{}

	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), statusPageEndpoint, id)

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
	return &result.Data, nil
}

func (c *client) DeleteStatusPage(id string) error {
	result := GetStatusPageResponse{}

	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), statusPageEndpoint, id)

	resp, err := c.rest.R().
		SetResult(&result).
		Delete(endpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() == http.StatusNotFound {
		fmt.Printf("status page not found\n")
		return nil
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("incorrect status response")
	}
	fmt.Printf("status page deleted\n")
	return nil
}

func (c *client) ListStatusPages() ([]StatusPage, error) {
	statusPages := []StatusPage{}

	result := ListStatusPageResponse{}
	endpoint := fmt.Sprintf("%s/%s", contextmanager.APIEndpoint(), statusPageEndpoint)

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

		statusPages = append(statusPages, result.Data...)

		if result.Pagination.Next == nil {
			break
		}
		if len(statusPages) > 50 {
			break
		}
		endpoint = *result.Pagination.Next
	}
	return statusPages, nil
}

type ListStatusPageResponse struct {
	Data       []StatusPage
	Pagination Pagination
}

type GetStatusPageResponse struct {
	Data StatusPage
}

type StatusPage struct {
	Id         string
	Type       string
	Attributes StatusPageAttributes
}

type StatusPageAttributes struct {
	CompanyName  string `json:"company_name"`
	CompanyURL   string `json:"company_url"`
	ContactURL   string `json:"contact_url"`
	LogoURL      string `json:"logo_url"`
	CustomDomain string `json:"custom_domain"`

	Announcement          string     `json:"announcement"`
	Subscribable          bool       `json:"subscribable"`
	HideFromSearchEngines bool       `json:"hide_from_search_engines"`
	PasswordEnabled       bool       `json:"password_enabled"`
	History               int        `json:"history"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdateAt              *time.Time `json:"updated_at"`
}

type CreateStatusPageRequest struct {
	CompanyName           *string    `json:"company_name,omitempty"`
	CompanyURL            *string    `json:"company_url,omitempty"`
	ContactURL            *string    `json:"contact_url,omitempty"`
	LogoURL               *string    `json:"logo_url,omitempty"`
	Timezone              *string    `json:"timezone,omitempty"`
	CustomDomain          *string    `json:"custom_domain,omitempty"`
	SubDomain             *string    `json:"subdomain,omitempty"`
	Announcement          *string    `json:"announcement,omitempty"`
	Subscribable          *bool      `json:"subscribable,omitempty"`
	HideFromSearchEngines *bool      `json:"hide_from_search_engines,omitempty"`
	PasswordEnabled       *bool      `json:"password_enabled,omitempty"`
	Password              *string    `json:"password,omitempty"`
	History               *int       `json:"history,omitempty"`
	CreatedAt             *time.Time `json:"created_at,omitempty"`
	UpdateAt              *time.Time `json:"updated_at,omitempty"`
	MinIncidentLength     *int       `json:"min_incident_length,omitempty"`
}
