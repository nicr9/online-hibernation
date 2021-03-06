package localresourceaccessreview

import (
	api "github.com/openshift/origin/pkg/authorization/api"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/runtime"
)

type Registry interface {
	CreateLocalResourceAccessReview(ctx kapi.Context, resourceAccessReview *api.LocalResourceAccessReview) (*api.ResourceAccessReviewResponse, error)
}

type Storage interface {
	Create(ctx kapi.Context, obj runtime.Object) (runtime.Object, error)
}

type storage struct {
	Storage
}

func NewRegistry(s Storage) Registry {
	return &storage{s}
}

func (s *storage) CreateLocalResourceAccessReview(ctx kapi.Context, resourceAccessReview *api.LocalResourceAccessReview) (*api.ResourceAccessReviewResponse, error) {
	obj, err := s.Create(ctx, resourceAccessReview)
	if err != nil {
		return nil, err
	}
	return obj.(*api.ResourceAccessReviewResponse), nil
}
