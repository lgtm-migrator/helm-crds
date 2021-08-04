/*
Copyright The Kubernetes Authors.

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
	"fmt"

	appv1 "github.com/alauda/helm-crds/pkg/apis/app/v1"
	v1alpha1 "github.com/alauda/helm-crds/pkg/apis/app/v1alpha1"
	v1beta1 "github.com/alauda/helm-crds/pkg/apis/app/v1beta1"
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
	// Group=app.alauda.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("charts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1alpha1().Charts().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("chartrepos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1alpha1().ChartRepos().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("helmrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1alpha1().HelmRequests().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("releases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1alpha1().Releases().Informer()}, nil

		// Group=app.alauda.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("charts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1beta1().Charts().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("chartrepos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1beta1().ChartRepos().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("helmrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1beta1().HelmRequests().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("releases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1beta1().Releases().Informer()}, nil

		// Group=app.cpaas.io, Version=v1
	case appv1.SchemeGroupVersion.WithResource("charts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1().Charts().Informer()}, nil
	case appv1.SchemeGroupVersion.WithResource("chartrepos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1().ChartRepos().Informer()}, nil
	case appv1.SchemeGroupVersion.WithResource("helmrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1().HelmRequests().Informer()}, nil
	case appv1.SchemeGroupVersion.WithResource("releases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.App().V1().Releases().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
