package apicontroller

import (
	"context"

	"github.com/golang/glog"
	"github.com/kyma-project/kyma/components/ui-api-layer/internal/domain/apicontroller/pretty"
	"github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlerror"
	"github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"
	"github.com/pkg/errors"
)

type apiResolver struct {
	apiLister    apiLister
	apiConverter apiConverter
}

func newApiResolver(lister apiLister) (*apiResolver, error) {
	if lister == nil {
		return nil, errors.New("Nil pointer for apiLister")
	}

	return &apiResolver{
		apiLister:    lister,
		apiConverter: apiConverter{},
	}, nil
}

func (ar *apiResolver) APIsQuery(ctx context.Context, environment *string, namespace *string, serviceName *string, hostname *string) ([]gqlschema.API, error) {
	if namespace == nil && environment == nil {
		return nil, errors.New("One of environment or namespace is required")
	}

	var ns string
	if namespace != nil {
		ns = *namespace
	} else {
		ns = *environment
	}
	apis, err := ar.apiLister.List(ns, serviceName, hostname)
	if err != nil {
		glog.Error(errors.Wrapf(err, "while listing %s for service name %v, hostname %v", pretty.APIs, serviceName, hostname))
		return nil, gqlerror.New(err, pretty.APIs, gqlerror.WithEnvironment(ns))
	}

	return ar.apiConverter.ToGQLs(apis), nil
}
