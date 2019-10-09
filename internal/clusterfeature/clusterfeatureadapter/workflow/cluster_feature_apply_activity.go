// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workflow

import (
	"context"

	"go.uber.org/cadence"

	"github.com/banzaicloud/pipeline/internal/clusterfeature"
)

const ClusterFeatureApplyActivityName = "cluster-feature-apply"

type ClusterFeatureApplyActivityInput struct {
	ClusterID   uint
	FeatureName string
	FeatureSpec clusterfeature.FeatureSpec
}

type ClusterFeatureApplyActivity struct {
	features clusterfeature.FeatureOperatorRegistry
}

func MakeClusterFeatureApplyActivity(features clusterfeature.FeatureOperatorRegistry) ClusterFeatureApplyActivity {
	return ClusterFeatureApplyActivity{
		features: features,
	}
}

func (a ClusterFeatureApplyActivity) Execute(ctx context.Context, input ClusterFeatureApplyActivityInput) error {
	f, err := a.features.GetFeatureOperator(input.FeatureName)
	if err != nil {
		return err
	}

	err = f.Apply(ctx, input.ClusterID, input.FeatureSpec)
	if ok := shouldRetry(err); ok {
		return cadence.NewCustomError(shouldRetryReason)
	}

	return err
}
