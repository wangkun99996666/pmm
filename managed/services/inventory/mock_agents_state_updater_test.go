// Code generated by mockery v1.0.0. DO NOT EDIT.

package inventory

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockAgentsStateUpdater is an autogenerated mock type for the agentsStateUpdater type
type mockAgentsStateUpdater struct {
	mock.Mock
}

// RequestStateUpdate provides a mock function with given fields: ctx, pmmAgentID
func (_m *mockAgentsStateUpdater) RequestStateUpdate(ctx context.Context, pmmAgentID string) {
	_m.Called(ctx, pmmAgentID)
}