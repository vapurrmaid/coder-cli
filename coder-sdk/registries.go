package coder

import (
	"context"
	"net/http"
	"time"
)

// Registry defines an image registry configuration.
type Registry struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	FriendlyName   string    `json:"friendly_name"`
	Registry       string    `json:"registry"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Registries fetches all registries in an organization.
func (c Client) Registries(ctx context.Context, orgID string) ([]Registry, error) {
	var r []Registry
	if err := c.requestBody(ctx, http.MethodGet, "/api/private/orgs/"+orgID+"/registries", nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// RegistryByID fetches a registry resource by its ID.
func (c Client) RegistryByID(ctx context.Context, orgID, registryID string) (*Registry, error) {
	var r Registry
	if err := c.requestBody(ctx, http.MethodGet, "/api/private/orgs/"+orgID+"/registries/"+registryID, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// UpdateRegistryReq defines the requests parameters for a partial update of a registry resource.
type UpdateRegistryReq struct {
	Registry     *string `json:"registry"`
	FriendlyName *string `json:"friendly_name"`
	Username     *string `json:"username"`
	Password     *string `json:"password"`
}

// UpdateRegistry applies a partial update to a registry resource.
func (c Client) UpdateRegistry(ctx context.Context, orgID, registryID string, req UpdateRegistryReq) error {
	return c.requestBody(ctx, http.MethodPatch, "/api/private/orgs/"+orgID+"/registries/"+registryID, req, nil)
}

// DeleteRegistry deletes a registry resource by its ID.
func (c Client) DeleteRegistry(ctx context.Context, orgID, registryID string) error {
	return c.requestBody(ctx, http.MethodDelete, "/api/private/orgs/"+orgID+"/registries/"+registryID, nil, nil)
}
