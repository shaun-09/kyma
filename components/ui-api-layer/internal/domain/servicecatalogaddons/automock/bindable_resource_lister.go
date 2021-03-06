// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

// bindableResourceLister is an autogenerated mock type for the bindableResourceLister type
type bindableResourceLister struct {
	mock.Mock
}

// ListResources provides a mock function with given fields: environment
func (_m *bindableResourceLister) ListResources(environment string) ([]gqlschema.BindableResourcesOutputItem, error) {
	ret := _m.Called(environment)

	var r0 []gqlschema.BindableResourcesOutputItem
	if rf, ok := ret.Get(0).(func(string) []gqlschema.BindableResourcesOutputItem); ok {
		r0 = rf(environment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.BindableResourcesOutputItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(environment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
