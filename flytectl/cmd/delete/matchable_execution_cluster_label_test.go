package delete

import (
	"fmt"
	"testing"

	"github.com/flyteorg/flyte/flytectl/cmd/config"
	"github.com/flyteorg/flyte/flytectl/cmd/config/subcommand/executionclusterlabel"
	"github.com/flyteorg/flyte/flytectl/cmd/testutils"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func deleteExecutionClusterLabelSetup() {
	executionclusterlabel.DefaultDelConfig = &executionclusterlabel.AttrDeleteConfig{}
	args = []string{}
}

func TestDeleteExecutionClusterLabels(t *testing.T) {
	t.Run("successful project domain attribute deletion commandline", func(t *testing.T) {
		s := testutils.Setup(t)

		deleteExecutionClusterLabelSetup()
		// Empty attribute file
		executionclusterlabel.DefaultDelConfig.AttrFile = ""
		// No args implying project domain attribute deletion
		s.DeleterExt.EXPECT().DeleteProjectDomainAttributes(mock.Anything, mock.Anything, mock.Anything,
			mock.Anything).Return(nil)
		err := deleteExecutionClusterLabel(s.Ctx, args, s.CmdCtx)
		assert.Nil(t, err)
		s.DeleterExt.AssertCalled(t, "DeleteProjectDomainAttributes",
			s.Ctx, config.GetConfig().Project, config.GetConfig().Domain, admin.MatchableResource_EXECUTION_CLUSTER_LABEL)
	})
	t.Run("failed project domain attribute deletion", func(t *testing.T) {
		s := testutils.Setup(t)

		deleteExecutionClusterLabelSetup()
		// No args implying project domain attribute deletion
		s.DeleterExt.EXPECT().DeleteProjectDomainAttributes(mock.Anything, mock.Anything, mock.Anything,
			mock.Anything).Return(fmt.Errorf("failed to delete project domain attributes"))
		err := deleteExecutionClusterLabel(s.Ctx, []string{}, s.CmdCtx)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("failed to delete project domain attributes"), err)
		s.DeleterExt.AssertCalled(t, "DeleteProjectDomainAttributes",
			s.Ctx, config.GetConfig().Project, config.GetConfig().Domain, admin.MatchableResource_EXECUTION_CLUSTER_LABEL)
	})
	t.Run("successful project domain attribute deletion file", func(t *testing.T) {
		s := testutils.Setup(t)

		deleteExecutionClusterLabelSetup()
		// Empty attribute file
		executionclusterlabel.DefaultDelConfig.AttrFile = "testdata/valid_project_domain_execution_cluster_label.yaml"
		// No args implying project domain attribute deletion
		s.DeleterExt.EXPECT().DeleteProjectDomainAttributes(mock.Anything, mock.Anything, mock.Anything,
			mock.Anything).Return(nil)
		err := deleteExecutionClusterLabel(s.Ctx, []string{}, s.CmdCtx)
		assert.Nil(t, err)
		s.DeleterExt.AssertCalled(t, "DeleteProjectDomainAttributes",
			s.Ctx, "flytesnacks", "development", admin.MatchableResource_EXECUTION_CLUSTER_LABEL)
	})
	t.Run("successful workflow attribute deletion", func(t *testing.T) {
		s := testutils.Setup(t)

		deleteExecutionClusterLabelSetup()
		// Empty attribute file
		executionclusterlabel.DefaultDelConfig.AttrFile = ""
		args := []string{"workflow1"}
		s.DeleterExt.EXPECT().DeleteWorkflowAttributes(mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything).Return(nil)
		err := deleteExecutionClusterLabel(s.Ctx, args, s.CmdCtx)
		assert.Nil(t, err)
		s.DeleterExt.AssertCalled(t, "DeleteWorkflowAttributes",
			s.Ctx, config.GetConfig().Project, config.GetConfig().Domain, "workflow1",
			admin.MatchableResource_EXECUTION_CLUSTER_LABEL)
	})
	t.Run("failed workflow attribute deletion", func(t *testing.T) {
		s := testutils.Setup(t)

		deleteExecutionClusterLabelSetup()
		// Empty attribute file
		executionclusterlabel.DefaultDelConfig.AttrFile = ""
		args := []string{"workflow1"}
		s.DeleterExt.EXPECT().DeleteWorkflowAttributes(mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything).Return(fmt.Errorf("failed to delete workflow attribute"))
		err := deleteExecutionClusterLabel(s.Ctx, args, s.CmdCtx)
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("failed to delete workflow attribute"), err)
		s.DeleterExt.AssertCalled(t, "DeleteWorkflowAttributes",
			s.Ctx, config.GetConfig().Project, config.GetConfig().Domain, "workflow1",
			admin.MatchableResource_EXECUTION_CLUSTER_LABEL)
	})
	t.Run("successful workflow attribute deletion file", func(t *testing.T) {
		s := testutils.Setup(t)

		deleteExecutionClusterLabelSetup()
		// Empty attribute file
		executionclusterlabel.DefaultDelConfig.AttrFile = "testdata/valid_workflow_execution_cluster_label.yaml"
		// No args implying project domain attribute deletion
		s.DeleterExt.EXPECT().DeleteWorkflowAttributes(mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything).Return(nil)
		err := deleteExecutionClusterLabel(s.Ctx, []string{}, s.CmdCtx)
		assert.Nil(t, err)
		s.DeleterExt.AssertCalled(t, "DeleteWorkflowAttributes",
			s.Ctx, "flytesnacks", "development", "core.control_flow.merge_sort.merge_sort",
			admin.MatchableResource_EXECUTION_CLUSTER_LABEL)
	})
	t.Run("workflow attribute deletion non existent file", func(t *testing.T) {
		s := testutils.Setup(t)

		deleteExecutionClusterLabelSetup()
		// Empty attribute file
		executionclusterlabel.DefaultDelConfig.AttrFile = testDataNonExistentFile
		// No args implying project domain attribute deletion
		s.DeleterExt.EXPECT().DeleteWorkflowAttributes(mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything).Return(nil)
		err := deleteExecutionClusterLabel(s.Ctx, []string{}, s.CmdCtx)
		assert.NotNil(t, err)
		s.DeleterExt.AssertNotCalled(t, "DeleteWorkflowAttributes",
			s.Ctx, "flytesnacks", "development", "core.control_flow.merge_sort.merge_sort",
			admin.MatchableResource_EXECUTION_CLUSTER_LABEL)
	})
	t.Run("attribute deletion invalid file", func(t *testing.T) {
		s := testutils.Setup(t)

		deleteExecutionClusterLabelSetup()
		// Empty attribute file
		executionclusterlabel.DefaultDelConfig.AttrFile = testDataInvalidAttrFile
		// No args implying project domain attribute deletion
		err := deleteExecutionClusterLabel(s.Ctx, []string{}, s.CmdCtx)
		assert.NotNil(t, err)
		assert.Equal(t,
			"error unmarshaling JSON: while decoding JSON: json: unknown field \"InvalidDomain\"",
			err.Error())
		s.DeleterExt.AssertNotCalled(t, "DeleteProjectDomainAttributes",
			s.Ctx, "flytesnacks", "development", admin.MatchableResource_EXECUTION_CLUSTER_LABEL)
	})
}
