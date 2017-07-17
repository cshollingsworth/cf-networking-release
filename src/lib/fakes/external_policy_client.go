// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lib/policy_client"
	"policy-server/models"
	"sync"
)

type ExternalPolicyClient struct {
	GetPoliciesStub        func(token string) ([]models.Policy, error)
	getPoliciesMutex       sync.RWMutex
	getPoliciesArgsForCall []struct {
		token string
	}
	getPoliciesReturns struct {
		result1 []models.Policy
		result2 error
	}
	getPoliciesReturnsOnCall map[int]struct {
		result1 []models.Policy
		result2 error
	}
	GetPoliciesByIDStub        func(token string, ids ...string) ([]models.Policy, error)
	getPoliciesByIDMutex       sync.RWMutex
	getPoliciesByIDArgsForCall []struct {
		token string
		ids   []string
	}
	getPoliciesByIDReturns struct {
		result1 []models.Policy
		result2 error
	}
	getPoliciesByIDReturnsOnCall map[int]struct {
		result1 []models.Policy
		result2 error
	}
	DeletePoliciesStub        func(token string, policies []models.Policy) error
	deletePoliciesMutex       sync.RWMutex
	deletePoliciesArgsForCall []struct {
		token    string
		policies []models.Policy
	}
	deletePoliciesReturns struct {
		result1 error
	}
	deletePoliciesReturnsOnCall map[int]struct {
		result1 error
	}
	AddPoliciesStub        func(token string, policies []models.Policy) error
	addPoliciesMutex       sync.RWMutex
	addPoliciesArgsForCall []struct {
		token    string
		policies []models.Policy
	}
	addPoliciesReturns struct {
		result1 error
	}
	addPoliciesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ExternalPolicyClient) GetPolicies(token string) ([]models.Policy, error) {
	fake.getPoliciesMutex.Lock()
	ret, specificReturn := fake.getPoliciesReturnsOnCall[len(fake.getPoliciesArgsForCall)]
	fake.getPoliciesArgsForCall = append(fake.getPoliciesArgsForCall, struct {
		token string
	}{token})
	fake.recordInvocation("GetPolicies", []interface{}{token})
	fake.getPoliciesMutex.Unlock()
	if fake.GetPoliciesStub != nil {
		return fake.GetPoliciesStub(token)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPoliciesReturns.result1, fake.getPoliciesReturns.result2
}

func (fake *ExternalPolicyClient) GetPoliciesCallCount() int {
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	return len(fake.getPoliciesArgsForCall)
}

func (fake *ExternalPolicyClient) GetPoliciesArgsForCall(i int) string {
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	return fake.getPoliciesArgsForCall[i].token
}

func (fake *ExternalPolicyClient) GetPoliciesReturns(result1 []models.Policy, result2 error) {
	fake.GetPoliciesStub = nil
	fake.getPoliciesReturns = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *ExternalPolicyClient) GetPoliciesReturnsOnCall(i int, result1 []models.Policy, result2 error) {
	fake.GetPoliciesStub = nil
	if fake.getPoliciesReturnsOnCall == nil {
		fake.getPoliciesReturnsOnCall = make(map[int]struct {
			result1 []models.Policy
			result2 error
		})
	}
	fake.getPoliciesReturnsOnCall[i] = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *ExternalPolicyClient) GetPoliciesByID(token string, ids ...string) ([]models.Policy, error) {
	fake.getPoliciesByIDMutex.Lock()
	ret, specificReturn := fake.getPoliciesByIDReturnsOnCall[len(fake.getPoliciesByIDArgsForCall)]
	fake.getPoliciesByIDArgsForCall = append(fake.getPoliciesByIDArgsForCall, struct {
		token string
		ids   []string
	}{token, ids})
	fake.recordInvocation("GetPoliciesByID", []interface{}{token, ids})
	fake.getPoliciesByIDMutex.Unlock()
	if fake.GetPoliciesByIDStub != nil {
		return fake.GetPoliciesByIDStub(token, ids...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPoliciesByIDReturns.result1, fake.getPoliciesByIDReturns.result2
}

func (fake *ExternalPolicyClient) GetPoliciesByIDCallCount() int {
	fake.getPoliciesByIDMutex.RLock()
	defer fake.getPoliciesByIDMutex.RUnlock()
	return len(fake.getPoliciesByIDArgsForCall)
}

func (fake *ExternalPolicyClient) GetPoliciesByIDArgsForCall(i int) (string, []string) {
	fake.getPoliciesByIDMutex.RLock()
	defer fake.getPoliciesByIDMutex.RUnlock()
	return fake.getPoliciesByIDArgsForCall[i].token, fake.getPoliciesByIDArgsForCall[i].ids
}

func (fake *ExternalPolicyClient) GetPoliciesByIDReturns(result1 []models.Policy, result2 error) {
	fake.GetPoliciesByIDStub = nil
	fake.getPoliciesByIDReturns = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *ExternalPolicyClient) GetPoliciesByIDReturnsOnCall(i int, result1 []models.Policy, result2 error) {
	fake.GetPoliciesByIDStub = nil
	if fake.getPoliciesByIDReturnsOnCall == nil {
		fake.getPoliciesByIDReturnsOnCall = make(map[int]struct {
			result1 []models.Policy
			result2 error
		})
	}
	fake.getPoliciesByIDReturnsOnCall[i] = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *ExternalPolicyClient) DeletePolicies(token string, policies []models.Policy) error {
	var policiesCopy []models.Policy
	if policies != nil {
		policiesCopy = make([]models.Policy, len(policies))
		copy(policiesCopy, policies)
	}
	fake.deletePoliciesMutex.Lock()
	ret, specificReturn := fake.deletePoliciesReturnsOnCall[len(fake.deletePoliciesArgsForCall)]
	fake.deletePoliciesArgsForCall = append(fake.deletePoliciesArgsForCall, struct {
		token    string
		policies []models.Policy
	}{token, policiesCopy})
	fake.recordInvocation("DeletePolicies", []interface{}{token, policiesCopy})
	fake.deletePoliciesMutex.Unlock()
	if fake.DeletePoliciesStub != nil {
		return fake.DeletePoliciesStub(token, policies)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deletePoliciesReturns.result1
}

func (fake *ExternalPolicyClient) DeletePoliciesCallCount() int {
	fake.deletePoliciesMutex.RLock()
	defer fake.deletePoliciesMutex.RUnlock()
	return len(fake.deletePoliciesArgsForCall)
}

func (fake *ExternalPolicyClient) DeletePoliciesArgsForCall(i int) (string, []models.Policy) {
	fake.deletePoliciesMutex.RLock()
	defer fake.deletePoliciesMutex.RUnlock()
	return fake.deletePoliciesArgsForCall[i].token, fake.deletePoliciesArgsForCall[i].policies
}

func (fake *ExternalPolicyClient) DeletePoliciesReturns(result1 error) {
	fake.DeletePoliciesStub = nil
	fake.deletePoliciesReturns = struct {
		result1 error
	}{result1}
}

func (fake *ExternalPolicyClient) DeletePoliciesReturnsOnCall(i int, result1 error) {
	fake.DeletePoliciesStub = nil
	if fake.deletePoliciesReturnsOnCall == nil {
		fake.deletePoliciesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deletePoliciesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ExternalPolicyClient) AddPolicies(token string, policies []models.Policy) error {
	var policiesCopy []models.Policy
	if policies != nil {
		policiesCopy = make([]models.Policy, len(policies))
		copy(policiesCopy, policies)
	}
	fake.addPoliciesMutex.Lock()
	ret, specificReturn := fake.addPoliciesReturnsOnCall[len(fake.addPoliciesArgsForCall)]
	fake.addPoliciesArgsForCall = append(fake.addPoliciesArgsForCall, struct {
		token    string
		policies []models.Policy
	}{token, policiesCopy})
	fake.recordInvocation("AddPolicies", []interface{}{token, policiesCopy})
	fake.addPoliciesMutex.Unlock()
	if fake.AddPoliciesStub != nil {
		return fake.AddPoliciesStub(token, policies)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addPoliciesReturns.result1
}

func (fake *ExternalPolicyClient) AddPoliciesCallCount() int {
	fake.addPoliciesMutex.RLock()
	defer fake.addPoliciesMutex.RUnlock()
	return len(fake.addPoliciesArgsForCall)
}

func (fake *ExternalPolicyClient) AddPoliciesArgsForCall(i int) (string, []models.Policy) {
	fake.addPoliciesMutex.RLock()
	defer fake.addPoliciesMutex.RUnlock()
	return fake.addPoliciesArgsForCall[i].token, fake.addPoliciesArgsForCall[i].policies
}

func (fake *ExternalPolicyClient) AddPoliciesReturns(result1 error) {
	fake.AddPoliciesStub = nil
	fake.addPoliciesReturns = struct {
		result1 error
	}{result1}
}

func (fake *ExternalPolicyClient) AddPoliciesReturnsOnCall(i int, result1 error) {
	fake.AddPoliciesStub = nil
	if fake.addPoliciesReturnsOnCall == nil {
		fake.addPoliciesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addPoliciesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ExternalPolicyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	fake.getPoliciesByIDMutex.RLock()
	defer fake.getPoliciesByIDMutex.RUnlock()
	fake.deletePoliciesMutex.RLock()
	defer fake.deletePoliciesMutex.RUnlock()
	fake.addPoliciesMutex.RLock()
	defer fake.addPoliciesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ExternalPolicyClient) recordInvocation(key string, args []interface{}) {
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

var _ policy_client.ExternalPolicyClient = new(ExternalPolicyClient)
