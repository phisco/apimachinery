/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/client-go/tools/cache"
)

// ScopeableSharedIndexInformer is an informer that knows how to scope itself down to one cluster,
// or act as an informer across clusters.
type ScopeableSharedIndexInformer interface {
	Cluster(clusterName logicalcluster.Name) cache.SharedIndexInformer
	ClusterWithContext(ctx context.Context, clusterName logicalcluster.Name) cache.SharedIndexInformer
	cache.SharedIndexInformer
}
