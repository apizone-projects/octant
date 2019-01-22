package fake

import (
	"context"

	"github.com/heptio/developer-dash/internal/hcli"
	"github.com/heptio/developer-dash/internal/view/component"
)

// SimpleClusterOverview is a fake that implements overview.Interface.
type SimpleClusterOverview struct{}

// NewSimpleClusterOverview creates an instance of SimpleClusterOverview.
func NewSimpleClusterOverview() *SimpleClusterOverview {
	return &SimpleClusterOverview{}
}

// Name is the module name.
func (sco *SimpleClusterOverview) Name() string {
	return "overview"
}

// Content generates content
func (sco *SimpleClusterOverview) Content(ctx context.Context, contentPath, prefix, namespace string) (component.ContentResponse, error) {
	return component.ContentResponse{}, nil
}

// ContentPath returns the content path for mounting this module.
func (sco *SimpleClusterOverview) ContentPath() string {
	return "/overview"
}

// Navigation is a no-op.
func (sco *SimpleClusterOverview) Navigation(root string) (*hcli.Navigation, error) {
	return nil, nil
}

// SetNamespace sets the namespace for this module. It is a no-op.
func (sco *SimpleClusterOverview) SetNamespace(namespace string) error {
	return nil
}

// Start starts the module. It is a no-op.
func (sco *SimpleClusterOverview) Start() error {
	return nil
}

// Stop stops the module. It is a no-op.
func (sco *SimpleClusterOverview) Stop() {
}
