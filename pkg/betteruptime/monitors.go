package betteruptime

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

var (
	monitorEndpoint = "/api/v2/monitors"
)

func (c *client) CreateMonitor(ctx context.Context, name string) (*Monitor, error) {
	result := GetMonitorResponse{}

	resp, err := c.rest.R().
		SetResult(&result).
		SetBody(CreateMonitorRequest{
			Name: name,
		}).
		Post(fmt.Sprintf("%s/%s", contextmanager.APIEndpoint(), monitorEndpoint))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("incorrect status response")
	}
	return &result.Data, nil
}

func (c *client) DeleteMonitor(ctx context.Context, id string) error {
	result := GetMonitorResponse{}

	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), monitorEndpoint, id)

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

func (c *client) UpdateMonitor(ctx context.Context, monitor *Monitor) error {
	result := GetMonitorResponse{}

	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), monitorEndpoint, monitor.Id)

	resp, err := c.rest.R().
		SetContext(ctx).
		SetResult(&result).
		SetBody(monitor).
		Patch(endpoint)
	if err != nil {
		return fmt.Errorf("error updating monitor: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("incorrect status response: %d", resp.StatusCode())
	}

	fmt.Printf("monitor \"%s\" updated\n", monitor.Id)

	return nil
}

func (c *client) GetMonitor(ctx context.Context, id string) (*Monitor, error) {
	monitor := Monitor{}

	result := Monitor{}
	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), monitorEndpoint, id)
	resp, err := c.rest.R().
		SetContext(ctx).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("incorrect status response")
	}

	fmt.Printf("%s", string(resp.Body()))

	return &monitor, nil
}

func (c *client) ListMonitors(ctx context.Context) ([]Monitor, error) {
	monitors := []Monitor{}

	result := ListMonitorResponse{}
	endpoint := fmt.Sprintf("%s/%s", contextmanager.APIEndpoint(), monitorEndpoint)

	for {
		resp, err := c.rest.R().
			SetResult(&result).
			SetContext(ctx).
			Get(endpoint)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode() != http.StatusOK {
			return nil, fmt.Errorf("incorrect status response")
		}

		monitors = append(monitors, result.Data...)

		if result.Pagination.Next == nil {
			break
		}
		if len(monitors) > 50 {
			break
		}
		endpoint = *result.Pagination.Next
	}
	return monitors, nil
}

func (c *client) PauseMonitor(id string) error {
	result := GetMonitorResponse{}

	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), monitorEndpoint, id)

	updateRequest := UpdateMonitorRequest{
		Paused: true,
	}

	resp, err := c.rest.R().
		// SetContext(ctx).
		SetResult(&result).
		SetBody(updateRequest).
		Patch(endpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("incorrect status response")
	}

	fmt.Printf("monitor \"%s\" paused\n", id)

	return nil
}

func (c *client) UnpauseMonitor(id string) error {
	result := GetMonitorResponse{}

	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), monitorEndpoint, id)

	updateRequest := UpdateMonitorRequest{
		Paused: false,
	}

	resp, err := c.rest.R().
		// SetContext(ctx).
		SetResult(&result).
		SetBody(updateRequest).
		Patch(endpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("incorrect status response")
	}

	fmt.Printf("monitor \"%s\" unpaused\n", id)

	return nil
}

type CreateMonitorRequest struct {
	Name      string `json:"name"`
	Paused    bool   `json:"paused"`
	SortIndex *int   `json:"sort_index"`
}

type UpdateMonitorRequest struct {
	Paused bool `json:"paused"`
}

type ListMonitorResponse struct {
	Data       []Monitor
	Pagination Pagination
}

type GetMonitorResponse struct {
	Data Monitor
}

type Monitor struct {
	Id         string
	Type       string
	Attributes MonitorAttributes
}

type MonitorAttributes struct {
	Url               string    `json:"url"`
	PronounceableName string    `json:"pronounceable_name"`
	Monitor_type      string    `json:"monitor_type"`
	Monitor_group_id  int       `json:"monitor_group_id"`
	LastCheckedAt     time.Time `json:"last_checked_at"`
	Status            string    `json:"status"`
	Policy_id         string    `json:"policy_id"`
	Required_keyword  string    `json:"required_keyword"`
	Verify_ssl        bool      `json:"verify_ssl"`
	Check_frequency   int       `json:"check_frequency"`
	Call              bool      `json:"call"`
	Sms               bool      `json:"sms"`
	Email             bool      `json:"email"`
	Push              bool      `json:"push"`
	Team_wait         bool      `json:"team_wait"`
	Http_method       string    `json:"http_method"`
	Request_timeout   int       `json:"request_timeout"`
	Recovery_period   int       `json:"recovery_period"`
	// Request_headers     string     `json:"request_headers"`
	Request_body        string     `json:"request_body"`
	Follow_redirects    bool       `json:"follow_redirects"`
	Remember_cookies    bool       `json:"remember_cookies"`
	Ssl_expiration      int        `json:"ssl_expiration"`
	Domain_expiration   string     `json:"domain_expiration"`
	Regions             []string   `json:"regions"`
	ExpectedStatusCodes []int      `json:"expected_status_codes"`
	Port                string     `json:"port"`
	ConfirmationPeriod  int        `json:"confirmation_period"`
	PausedAt            *time.Time `json:"paused_at"`
	Paused              bool       `json:"paused"`
	MaintenanceFrom     *time.Time `json:"maintenance_from"`
	MaintenanceTo       *time.Time `json:"maintenance_to"`
	MaintenanceTimezone string     `json:"maintenance_timezone"`
	MaintenanceDays     []string   `json:"maintenance_days"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdateAt            *time.Time `json:"updated_at"`
}

// string `json:""`
