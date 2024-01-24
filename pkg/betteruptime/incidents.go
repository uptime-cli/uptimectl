package betteruptime

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

var (
	incidentsEndpoint = "/api/v2/incidents"
	incidentIDRegex   = regexp.MustCompile(`incidents/(\d*)[/]?`)
	numericCheckRegex = regexp.MustCompile(`^[0-9]+$`)
)

var (
	ErrIncidentAlreadyAcknowledged = errors.New("incident is already acknowledged")
	ErrIncidentAlreadyResolved     = errors.New("incident is already resolved")
	ErrIncidentNotFound            = errors.New("incident not found")
	ErrUnexpectedStatusCodeFromAPI = errors.New("unexpected status received from better-uptime API")
)

// Extracts a betteruptime incident ID from the URL.
// If an incidentID is provided (and not a URL), it returns that without any further processing
func IncidentIDFromURL(incidentStr string) (string, error) {
	// Check if passed-in incidentStr is entirely numeric, if so, assume it's an incidentID and return it
	if numericCheckRegex.MatchString(incidentStr) {
		return incidentStr, nil
	}

	incidentURL, err := url.Parse(incidentStr)
	if err != nil {
		return "", err
	}

	matches := incidentIDRegex.FindStringSubmatch(incidentURL.Path)
	if len(matches) != 2 {
		return "", fmt.Errorf("invalid incident URL: %s", incidentURL)
	}

	return matches[1], nil
}

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

	switch resp.StatusCode() {
	case http.StatusNotFound:
		return ErrIncidentNotFound
	case http.StatusNoContent:
		return nil
	default:
		return ErrUnexpectedStatusCodeFromAPI
	}
}

func (c *client) AcknowledgeIncident(ctx context.Context, incidentID, acknowledgedBy string) error {
	endpoint := fmt.Sprintf("%s/%s/%s/acknowledge", contextmanager.APIEndpoint(), incidentsEndpoint, incidentID)

	resp, err := c.rest.R().
		SetContext(ctx).
		SetBody(Acknowledge{
			AcknowledgedBy: acknowledgedBy,
		}).
		Post(endpoint)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusNotFound:
		return ErrIncidentNotFound
	case http.StatusConflict:
		return ErrIncidentAlreadyAcknowledged
	case http.StatusOK:
		return nil // incident resolved successfully
	default:
		return ErrUnexpectedStatusCodeFromAPI
	}
}

func (c *client) ResolveIncident(incidentID string, resolvedBy string) error {
	endpoint := fmt.Sprintf("%s/%s/%s/resolve", contextmanager.APIEndpoint(), incidentsEndpoint, incidentID)

	resp, err := c.rest.R().
		SetBody(Resolve{
			ResolvedBy: resolvedBy,
		}).
		Post(endpoint)
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusNotFound:
		return ErrIncidentNotFound
	case http.StatusConflict:
		return ErrIncidentAlreadyResolved
	case http.StatusOK:
		return nil // incident resolved successfully
	default:
		return ErrUnexpectedStatusCodeFromAPI
	}
}

type Resolve struct {
	ResolvedBy string `json:"resolved_by"`
}

type Acknowledge struct {
	AcknowledgedBy string `json:"acknowledged_by"`
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
