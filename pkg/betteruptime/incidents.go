package betteruptime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

var (
	incidentsEndpoint = "/api/v2/incidents"
)

func (c *client) ListIncidents(showResolved bool, daysInPast int, showMax int) ([]Incident, error) {
	incidents := []Incident{}

	result := ListIncidentResponse{}
	endpoint := fmt.Sprintf("%s/%s", contextmanager.APIEndpoint(), incidentsEndpoint)

	// oneWeekAgo := ""
	oneWeekAgo := time.Now().Local().AddDate(0, 0, -daysInPast)
	toDate := oneWeekAgo.Format("2006-01-02")

	for {
		resp, err := c.rest.R().
			SetResult(&result).
			SetQueryParam("from", toDate).
			Get(endpoint)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode() != http.StatusOK {
			fmt.Printf("%+v", string(resp.Body()))
			return nil, fmt.Errorf("incorrect status response")
		}

		for _, incident := range result.Data {
			if showResolved || incident.Attributes.Resolved_at == nil {
				incidents = append(incidents, incident)
			}
		}

		if result.Pagination.Next == nil {
			break
		}
		if showMax > 0 && len(incidents) > showMax {
			break
		}
		endpoint = *result.Pagination.Next
	}

	return incidents, nil
}

func (c *client) DeleteIncident(incidentID string) error {
	endpoint := fmt.Sprintf("%s/%s/%s", contextmanager.APIEndpoint(), incidentsEndpoint, incidentID)

	resp, err := c.rest.R().
		Delete(endpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() == http.StatusNotFound {
		return fmt.Errorf("incident not found")
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("incorrect status response from api")
	}
	return nil
}

type ListIncidentResponse struct {
	Data       []Incident
	Pagination Pagination
}

type Incident struct {
	Id   string
	Type string
	// Attributes    map[string]interface{}
	Attributes    IncidentAttributes
	Relationships map[string]IncidentRelationship
}

type IncidentAttributes struct {
	Name                 string
	Url                  string
	Http_method          string
	Cause                string
	Incident_group_id    *int
	Started_at           time.Time
	Acknowledged_at      *time.Time
	Acknowledged_by      *string
	Resolved_at          *time.Time
	Resolved_by          *string
	Response_content     *string
	Response_options     *string
	Regions              []string
	Response_url         *string
	Screenshot_url       *string
	Escalation_policy_id *int
	Call                 bool
	Sms                  bool
	Email                bool
	Push                 bool
}

type IncidentRelationship struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
