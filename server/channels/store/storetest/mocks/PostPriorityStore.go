// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// PostPriorityStore is an autogenerated mock type for the PostPriorityStore type
type PostPriorityStore struct {
	mock.Mock
}

// GetForPost provides a mock function with given fields: postId
func (_m *PostPriorityStore) GetForPost(postId string) (*model.PostPriority, error) {
	ret := _m.Called(postId)

	var r0 *model.PostPriority
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.PostPriority, error)); ok {
		return rf(postId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.PostPriority); ok {
		r0 = rf(postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostPriority)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForPosts provides a mock function with given fields: ids
func (_m *PostPriorityStore) GetForPosts(ids []string) ([]*model.PostPriority, error) {
	ret := _m.Called(ids)

	var r0 []*model.PostPriority
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*model.PostPriority, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]string) []*model.PostPriority); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PostPriority)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostPriorityStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostPriorityStore creates a new instance of PostPriorityStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostPriorityStore(t mockConstructorTestingTNewPostPriorityStore) *PostPriorityStore {
	mock := &PostPriorityStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
