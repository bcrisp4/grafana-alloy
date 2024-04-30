package client

import (
	"context"
	"io"

	alertmanagerconfig "github.com/prometheus/alertmanager/config"
	"gopkg.in/yaml.v3"
)

// CreateAlertmanagerConfig creates a new alertmanager config
func (r *MimirClient) CreateAlertmanagerConfig(ctx context.Context, amc alertmanagerconfig.Config) error {
	payload, err := yaml.Marshal(&amc)
	if err != nil {
		return err
	}

	path := "/api/v1/alerts"

	res, err := r.doRequest(path, path, "POST", payload)
	if err != nil {
		return err
	}

	res.Body.Close()

	return nil
}

// DeleteAlertmanagerConfig deletes an alertmanager config
func (r *MimirClient) DeleteAlertmanagerConfig(ctx context.Context) error {
	path := "/api/v1/alerts"

	res, err := r.doRequest(path, path, "DELETE", nil)
	if err != nil {
		return err
	}

	res.Body.Close()

	return nil

}

// GetAlertmanagerConfig retrieves an alertmanager config
func (r *MimirClient) GetAlertmanagerConfig(ctx context.Context) (alertmanagerconfig.Config, error) {
	path := "/api/v1/alerts"

	res, err := r.doRequest(path, path, "GET", nil)
	if err != nil {
		return alertmanagerconfig.Config{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return alertmanagerconfig.Config{}, err
	}

	amc := alertmanagerconfig.Config{}

	err = yaml.Unmarshal(body, &amc)
	if err != nil {
		return alertmanagerconfig.Config{}, err
	}

	return amc, nil
}
