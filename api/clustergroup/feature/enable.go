package feature

import (
	"context"
	"net/http"

	cgroupIAPI "github.com/banzaicloud/pipeline/internal/clustergroup/api"
	ginutils "github.com/banzaicloud/pipeline/internal/platform/gin/utils"
	pkgCommon "github.com/banzaicloud/pipeline/pkg/common"
	"github.com/gin-gonic/gin"
)

func (n *API) Enable(c *gin.Context) {
	ctx := ginutils.Context(context.Background(), c)

	var req cgroupIAPI.FeatureRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkgCommon.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error parsing request",
			Error:   err.Error(),
		})
		return
	}

	clusterGroupId, ok := ginutils.UintParam(c, "id")
	if !ok {
		return
	}

	clusterGroup, err := n.clusterGroupManager.GetClusterGroupById(ctx, clusterGroupId)
	if err != nil {
		n.errorHandler.Handle(c, err)
		return
	}

	featureName := c.Param("featureName")

	err = n.clusterGroupManager.SetFeatureParams(featureName, clusterGroup, true, req.Properties)
	if err != nil {
		n.errorHandler.Handle(c, err)
		return
	}

	err = n.clusterGroupManager.ReconcileFeatures(*clusterGroup, false)
	if err != nil {
		n.errorHandler.Handle(c, err)
		return
	}

	c.Status(http.StatusOK)
}
