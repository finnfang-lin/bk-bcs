/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/apis/network/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudListenerLister helps list CloudListeners.
type CloudListenerLister interface {
	// List lists all CloudListeners in the indexer.
	List(selector labels.Selector) (ret []*v1.CloudListener, err error)
	// CloudListeners returns an object that can list and get CloudListeners.
	CloudListeners(namespace string) CloudListenerNamespaceLister
	CloudListenerListerExpansion
}

// cloudListenerLister implements the CloudListenerLister interface.
type cloudListenerLister struct {
	indexer cache.Indexer
}

// NewCloudListenerLister returns a new CloudListenerLister.
func NewCloudListenerLister(indexer cache.Indexer) CloudListenerLister {
	return &cloudListenerLister{indexer: indexer}
}

// List lists all CloudListeners in the indexer.
func (s *cloudListenerLister) List(selector labels.Selector) (ret []*v1.CloudListener, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CloudListener))
	})
	return ret, err
}

// CloudListeners returns an object that can list and get CloudListeners.
func (s *cloudListenerLister) CloudListeners(namespace string) CloudListenerNamespaceLister {
	return cloudListenerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudListenerNamespaceLister helps list and get CloudListeners.
type CloudListenerNamespaceLister interface {
	// List lists all CloudListeners in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CloudListener, err error)
	// Get retrieves the CloudListener from the indexer for a given namespace and name.
	Get(name string) (*v1.CloudListener, error)
	CloudListenerNamespaceListerExpansion
}

// cloudListenerNamespaceLister implements the CloudListenerNamespaceLister
// interface.
type cloudListenerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudListeners in the indexer for a given namespace.
func (s cloudListenerNamespaceLister) List(selector labels.Selector) (ret []*v1.CloudListener, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CloudListener))
	})
	return ret, err
}

// Get retrieves the CloudListener from the indexer for a given namespace and name.
func (s cloudListenerNamespaceLister) Get(name string) (*v1.CloudListener, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cloudlistener"), name)
	}
	return obj.(*v1.CloudListener), nil
}