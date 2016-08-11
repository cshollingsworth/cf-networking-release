// This file was generated by counterfeiter
package fakes

import "sync"

type Converger struct {
	ConvergeStub        func() error
	convergeMutex       sync.RWMutex
	convergeArgsForCall []struct{}
	convergeReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Converger) Converge() error {
	fake.convergeMutex.Lock()
	fake.convergeArgsForCall = append(fake.convergeArgsForCall, struct{}{})
	fake.recordInvocation("Converge", []interface{}{})
	fake.convergeMutex.Unlock()
	if fake.ConvergeStub != nil {
		return fake.ConvergeStub()
	} else {
		return fake.convergeReturns.result1
	}
}

func (fake *Converger) ConvergeCallCount() int {
	fake.convergeMutex.RLock()
	defer fake.convergeMutex.RUnlock()
	return len(fake.convergeArgsForCall)
}

func (fake *Converger) ConvergeReturns(result1 error) {
	fake.ConvergeStub = nil
	fake.convergeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Converger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convergeMutex.RLock()
	defer fake.convergeMutex.RUnlock()
	return fake.invocations
}

func (fake *Converger) recordInvocation(key string, args []interface{}) {
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
