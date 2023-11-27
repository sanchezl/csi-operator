// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OLMLister helps list OLMs.
// All objects returned here must be treated as read-only.
type OLMLister interface {
	// List lists all OLMs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OLM, err error)
	// Get retrieves the OLM from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OLM, error)
	OLMListerExpansion
}

// oLMLister implements the OLMLister interface.
type oLMLister struct {
	indexer cache.Indexer
}

// NewOLMLister returns a new OLMLister.
func NewOLMLister(indexer cache.Indexer) OLMLister {
	return &oLMLister{indexer: indexer}
}

// List lists all OLMs in the indexer.
func (s *oLMLister) List(selector labels.Selector) (ret []*v1alpha1.OLM, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OLM))
	})
	return ret, err
}

// Get retrieves the OLM from the index for a given name.
func (s *oLMLister) Get(name string) (*v1alpha1.OLM, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("olm"), name)
	}
	return obj.(*v1alpha1.OLM), nil
}