// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"

// ServiceCatalogSyncer is an autogenerated mock type for the ServiceCatalogSyncer type
type ServiceCatalogSyncer struct {
	mock.Mock
}

// Sync provides a mock function with given fields: name, retries
func (_m *ServiceCatalogSyncer) Sync(name string, retries int) error {
	ret := _m.Called(name, retries)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(name, retries)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}