/*
Copyright 2023 The KServe Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	fmt "fmt"

	v1alpha1 "github.com/kserve/kserve/pkg/apis/serving/v1alpha1"
	v1beta1 "github.com/kserve/kserve/pkg/apis/serving/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=serving.kserve.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("clusterservingruntimes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().ClusterServingRuntimes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterstoragecontainers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().ClusterStorageContainers().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("inferencegraphs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().InferenceGraphs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("localmodelcaches"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().LocalModelCaches().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("localmodelnodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().LocalModelNodes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("localmodelnodegroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().LocalModelNodeGroups().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("servingruntimes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().ServingRuntimes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("trainedmodels"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().TrainedModels().Informer()}, nil

		// Group=serving.kserve.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("inferenceservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1beta1().InferenceServices().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
