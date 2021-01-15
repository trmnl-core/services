// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	m3o_alert "github.com/trmnl-core/services/alert/proto/alert"
	"github.com/micro/micro/v3/service/client"
)

type FakeAlertService struct {
	ReportEventStub        func(context.Context, *m3o_alert.ReportEventRequest, ...client.CallOption) (*m3o_alert.ReportEventResponse, error)
	reportEventMutex       sync.RWMutex
	reportEventArgsForCall []struct {
		arg1 context.Context
		arg2 *m3o_alert.ReportEventRequest
		arg3 []client.CallOption
	}
	reportEventReturns struct {
		result1 *m3o_alert.ReportEventResponse
		result2 error
	}
	reportEventReturnsOnCall map[int]struct {
		result1 *m3o_alert.ReportEventResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAlertService) ReportEvent(arg1 context.Context, arg2 *m3o_alert.ReportEventRequest, arg3 ...client.CallOption) (*m3o_alert.ReportEventResponse, error) {
	fake.reportEventMutex.Lock()
	ret, specificReturn := fake.reportEventReturnsOnCall[len(fake.reportEventArgsForCall)]
	fake.reportEventArgsForCall = append(fake.reportEventArgsForCall, struct {
		arg1 context.Context
		arg2 *m3o_alert.ReportEventRequest
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.ReportEventStub
	fakeReturns := fake.reportEventReturns
	fake.recordInvocation("ReportEvent", []interface{}{arg1, arg2, arg3})
	fake.reportEventMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAlertService) ReportEventCallCount() int {
	fake.reportEventMutex.RLock()
	defer fake.reportEventMutex.RUnlock()
	return len(fake.reportEventArgsForCall)
}

func (fake *FakeAlertService) ReportEventCalls(stub func(context.Context, *m3o_alert.ReportEventRequest, ...client.CallOption) (*m3o_alert.ReportEventResponse, error)) {
	fake.reportEventMutex.Lock()
	defer fake.reportEventMutex.Unlock()
	fake.ReportEventStub = stub
}

func (fake *FakeAlertService) ReportEventArgsForCall(i int) (context.Context, *m3o_alert.ReportEventRequest, []client.CallOption) {
	fake.reportEventMutex.RLock()
	defer fake.reportEventMutex.RUnlock()
	argsForCall := fake.reportEventArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAlertService) ReportEventReturns(result1 *m3o_alert.ReportEventResponse, result2 error) {
	fake.reportEventMutex.Lock()
	defer fake.reportEventMutex.Unlock()
	fake.ReportEventStub = nil
	fake.reportEventReturns = struct {
		result1 *m3o_alert.ReportEventResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeAlertService) ReportEventReturnsOnCall(i int, result1 *m3o_alert.ReportEventResponse, result2 error) {
	fake.reportEventMutex.Lock()
	defer fake.reportEventMutex.Unlock()
	fake.ReportEventStub = nil
	if fake.reportEventReturnsOnCall == nil {
		fake.reportEventReturnsOnCall = make(map[int]struct {
			result1 *m3o_alert.ReportEventResponse
			result2 error
		})
	}
	fake.reportEventReturnsOnCall[i] = struct {
		result1 *m3o_alert.ReportEventResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeAlertService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.reportEventMutex.RLock()
	defer fake.reportEventMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAlertService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ m3o_alert.AlertService = new(FakeAlertService)
