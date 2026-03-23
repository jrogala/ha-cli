package ops

import "github.com/jrogala/ha-cli/client"

// ServiceEntry represents a service in listing results.
type ServiceEntry struct {
	Domain  string `json:"domain"`
	Service string `json:"service"`
}

// ListServices returns available services, optionally filtered by domain.
func ListServices(c *client.Client, domain string) ([]ServiceEntry, error) {
	services, err := c.GetServices()
	if err != nil {
		return nil, err
	}

	var entries []ServiceEntry
	for d, svcs := range services {
		if domain != "" && d != domain {
			continue
		}
		for svc := range svcs {
			entries = append(entries, ServiceEntry{
				Domain:  d,
				Service: svc,
			})
		}
	}
	return entries, nil
}

// CallResult holds the result of a service call.
type CallResult struct {
	Domain  string `json:"domain"`
	Service string `json:"service"`
	Success bool   `json:"success"`
}

// CallService calls a Home Assistant service with the given data.
func CallService(c *client.Client, domain, service string, data map[string]any) (*CallResult, error) {
	_, err := c.CallService(domain, service, data)
	if err != nil {
		return nil, err
	}
	return &CallResult{
		Domain:  domain,
		Service: service,
		Success: true,
	}, nil
}
