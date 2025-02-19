// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// RemoteClusterStore is an autogenerated mock type for the RemoteClusterStore type
type RemoteClusterStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: remoteClusterId
func (_m *RemoteClusterStore) Delete(remoteClusterId string) (bool, error) {
	ret := _m.Called(remoteClusterId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(remoteClusterId)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(remoteClusterId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(remoteClusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: remoteClusterId
func (_m *RemoteClusterStore) Get(remoteClusterId string) (*model.RemoteCluster, error) {
	ret := _m.Called(remoteClusterId)

	var r0 *model.RemoteCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.RemoteCluster, error)); ok {
		return rf(remoteClusterId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.RemoteCluster); ok {
		r0 = rf(remoteClusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RemoteCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(remoteClusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: filter
func (_m *RemoteClusterStore) GetAll(filter model.RemoteClusterQueryFilter) ([]*model.RemoteCluster, error) {
	ret := _m.Called(filter)

	var r0 []*model.RemoteCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(model.RemoteClusterQueryFilter) ([]*model.RemoteCluster, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.RemoteClusterQueryFilter) []*model.RemoteCluster); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RemoteCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(model.RemoteClusterQueryFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: rc
func (_m *RemoteClusterStore) Save(rc *model.RemoteCluster) (*model.RemoteCluster, error) {
	ret := _m.Called(rc)

	var r0 *model.RemoteCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.RemoteCluster) (*model.RemoteCluster, error)); ok {
		return rf(rc)
	}
	if rf, ok := ret.Get(0).(func(*model.RemoteCluster) *model.RemoteCluster); ok {
		r0 = rf(rc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RemoteCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.RemoteCluster) error); ok {
		r1 = rf(rc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetLastPingAt provides a mock function with given fields: remoteClusterId
func (_m *RemoteClusterStore) SetLastPingAt(remoteClusterId string) error {
	ret := _m.Called(remoteClusterId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(remoteClusterId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: rc
func (_m *RemoteClusterStore) Update(rc *model.RemoteCluster) (*model.RemoteCluster, error) {
	ret := _m.Called(rc)

	var r0 *model.RemoteCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.RemoteCluster) (*model.RemoteCluster, error)); ok {
		return rf(rc)
	}
	if rf, ok := ret.Get(0).(func(*model.RemoteCluster) *model.RemoteCluster); ok {
		r0 = rf(rc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RemoteCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.RemoteCluster) error); ok {
		r1 = rf(rc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTopics provides a mock function with given fields: remoteClusterId, topics
func (_m *RemoteClusterStore) UpdateTopics(remoteClusterId string, topics string) (*model.RemoteCluster, error) {
	ret := _m.Called(remoteClusterId, topics)

	var r0 *model.RemoteCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*model.RemoteCluster, error)); ok {
		return rf(remoteClusterId, topics)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.RemoteCluster); ok {
		r0 = rf(remoteClusterId, topics)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RemoteCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(remoteClusterId, topics)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRemoteClusterStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewRemoteClusterStore creates a new instance of RemoteClusterStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRemoteClusterStore(t mockConstructorTestingTNewRemoteClusterStore) *RemoteClusterStore {
	mock := &RemoteClusterStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
