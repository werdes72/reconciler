// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	pod "github.com/kyma-incubator/reconciler/pkg/reconciler/instances/istio/reset/pod"
	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *Handler) Execute(_a0 pod.CustomObject, _a1 pod.GetSyncWG) {
	_m.Called(_a0, _a1)
}