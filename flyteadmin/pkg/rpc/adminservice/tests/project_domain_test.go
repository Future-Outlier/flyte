package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/flyteorg/flyte/flyteadmin/pkg/manager/mocks"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
)

func TestUpdateProjectDomain(t *testing.T) {
	ctx := context.Background()

	mockProjectDomainManager := mocks.ResourceInterface{}
	var updateCalled bool
	mockProjectDomainManager.EXPECT().UpdateProjectDomainAttributes(mock.Anything, mock.Anything).RunAndReturn(
		func(ctx context.Context,
			request *admin.ProjectDomainAttributesUpdateRequest) (*admin.ProjectDomainAttributesUpdateResponse, error) {
			updateCalled = true
			return &admin.ProjectDomainAttributesUpdateResponse{}, nil
		},
	)
	mockServer := NewMockAdminServer(NewMockAdminServerInput{
		resourceManager: &mockProjectDomainManager,
	})

	resp, err := mockServer.UpdateProjectDomainAttributes(ctx, &admin.ProjectDomainAttributesUpdateRequest{
		Attributes: &admin.ProjectDomainAttributes{
			Project: "project",
			Domain:  "domain",
		},
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.True(t, updateCalled)
}

func TestUpdateProjectAttr(t *testing.T) {
	ctx := context.Background()

	mockProjectDomainManager := mocks.ResourceInterface{}
	var updateCalled bool
	mockProjectDomainManager.EXPECT().UpdateProjectAttributes(mock.Anything, mock.Anything).RunAndReturn((func(ctx context.Context,
		request *admin.ProjectAttributesUpdateRequest) (*admin.ProjectAttributesUpdateResponse, error) {
		updateCalled = true
		return &admin.ProjectAttributesUpdateResponse{}, nil
	}),
	)
	mockServer := NewMockAdminServer(NewMockAdminServerInput{
		resourceManager: &mockProjectDomainManager,
	})

	resp, err := mockServer.UpdateProjectAttributes(ctx, &admin.ProjectAttributesUpdateRequest{
		Attributes: &admin.ProjectAttributes{
			Project: "project",
		},
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.True(t, updateCalled)
}

func TestDeleteProjectAttr(t *testing.T) {
	ctx := context.Background()

	mockProjectDomainManager := mocks.ResourceInterface{}
	var deleteCalled bool
	mockProjectDomainManager.EXPECT().DeleteProjectAttributes(mock.Anything, mock.Anything).RunAndReturn((func(ctx context.Context,
		request *admin.ProjectAttributesDeleteRequest) (*admin.ProjectAttributesDeleteResponse, error) {
		deleteCalled = true
		return &admin.ProjectAttributesDeleteResponse{}, nil
	}),
	)
	mockServer := NewMockAdminServer(NewMockAdminServerInput{
		resourceManager: &mockProjectDomainManager,
	})

	resp, err := mockServer.DeleteProjectAttributes(ctx, &admin.ProjectAttributesDeleteRequest{
		Project:      "project",
		ResourceType: admin.MatchableResource_WORKFLOW_EXECUTION_CONFIG,
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.True(t, deleteCalled)
}

func TestGetProjectAttr(t *testing.T) {
	ctx := context.Background()

	mockProjectDomainManager := mocks.ResourceInterface{}
	var getCalled bool
	mockProjectDomainManager.EXPECT().GetProjectAttributes(mock.Anything, mock.Anything).RunAndReturn((func(ctx context.Context,
		request *admin.ProjectAttributesGetRequest) (*admin.ProjectAttributesGetResponse, error) {
		getCalled = true
		return &admin.ProjectAttributesGetResponse{}, nil
	}),
	)
	mockServer := NewMockAdminServer(NewMockAdminServerInput{
		resourceManager: &mockProjectDomainManager,
	})

	resp, err := mockServer.GetProjectAttributes(ctx, &admin.ProjectAttributesGetRequest{
		Project:      "project",
		ResourceType: admin.MatchableResource_WORKFLOW_EXECUTION_CONFIG,
	})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
	assert.True(t, getCalled)
}
