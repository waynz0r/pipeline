package feature

import (
	"github.com/banzaicloud/pipeline/api/clustergroup/common"
	cgroup "github.com/banzaicloud/pipeline/internal/clustergroup"
	"github.com/gin-gonic/gin"
	"github.com/goph/emperror"
	"github.com/sirupsen/logrus"
)

const (
	IDParamName = "featureName"
)

type API struct {
	clusterGroupManager *cgroup.Manager
	deploymentManager   *cgroup.CGDeploymentManager
	logger              logrus.FieldLogger
	errorHandler        common.ErrorHandler
}

func NewAPI(
	clusterGroupManager *cgroup.Manager,
	deploymentManager *cgroup.CGDeploymentManager,
	logger logrus.FieldLogger,
	baseErrorHandler emperror.Handler,
) *API {
	return &API{
		clusterGroupManager: clusterGroupManager,
		deploymentManager:   deploymentManager,
		logger:              logger,
		errorHandler: common.ErrorHandler{
			Handler: baseErrorHandler,
		},
	}
}

// AddRoutes adds cluster group features related API routes
func (a *API) AddRoutes(group *gin.RouterGroup) {
	group.GET("", a.List)
	item := group.Group("/:" + IDParamName)
	{
		item.POST("", a.Enable)
		item.GET("", a.Get)
		item.PUT("", a.Update)
		item.DELETE("", a.Disable)
	}
}
