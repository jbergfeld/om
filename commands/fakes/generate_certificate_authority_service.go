// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type GenerateCertificateAuthorityService struct {
	GenerateCertificateAuthorityStub        func() (api.CA, error)
	generateCertificateAuthorityMutex       sync.RWMutex
	generateCertificateAuthorityArgsForCall []struct {
	}
	generateCertificateAuthorityReturns struct {
		result1 api.CA
		result2 error
	}
	generateCertificateAuthorityReturnsOnCall map[int]struct {
		result1 api.CA
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *GenerateCertificateAuthorityService) GenerateCertificateAuthority() (api.CA, error) {
	fake.generateCertificateAuthorityMutex.Lock()
	ret, specificReturn := fake.generateCertificateAuthorityReturnsOnCall[len(fake.generateCertificateAuthorityArgsForCall)]
	fake.generateCertificateAuthorityArgsForCall = append(fake.generateCertificateAuthorityArgsForCall, struct {
	}{})
	fake.recordInvocation("GenerateCertificateAuthority", []interface{}{})
	fake.generateCertificateAuthorityMutex.Unlock()
	if fake.GenerateCertificateAuthorityStub != nil {
		return fake.GenerateCertificateAuthorityStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateCertificateAuthorityReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *GenerateCertificateAuthorityService) GenerateCertificateAuthorityCallCount() int {
	fake.generateCertificateAuthorityMutex.RLock()
	defer fake.generateCertificateAuthorityMutex.RUnlock()
	return len(fake.generateCertificateAuthorityArgsForCall)
}

func (fake *GenerateCertificateAuthorityService) GenerateCertificateAuthorityCalls(stub func() (api.CA, error)) {
	fake.generateCertificateAuthorityMutex.Lock()
	defer fake.generateCertificateAuthorityMutex.Unlock()
	fake.GenerateCertificateAuthorityStub = stub
}

func (fake *GenerateCertificateAuthorityService) GenerateCertificateAuthorityReturns(result1 api.CA, result2 error) {
	fake.generateCertificateAuthorityMutex.Lock()
	defer fake.generateCertificateAuthorityMutex.Unlock()
	fake.GenerateCertificateAuthorityStub = nil
	fake.generateCertificateAuthorityReturns = struct {
		result1 api.CA
		result2 error
	}{result1, result2}
}

func (fake *GenerateCertificateAuthorityService) GenerateCertificateAuthorityReturnsOnCall(i int, result1 api.CA, result2 error) {
	fake.generateCertificateAuthorityMutex.Lock()
	defer fake.generateCertificateAuthorityMutex.Unlock()
	fake.GenerateCertificateAuthorityStub = nil
	if fake.generateCertificateAuthorityReturnsOnCall == nil {
		fake.generateCertificateAuthorityReturnsOnCall = make(map[int]struct {
			result1 api.CA
			result2 error
		})
	}
	fake.generateCertificateAuthorityReturnsOnCall[i] = struct {
		result1 api.CA
		result2 error
	}{result1, result2}
}

func (fake *GenerateCertificateAuthorityService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateCertificateAuthorityMutex.RLock()
	defer fake.generateCertificateAuthorityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *GenerateCertificateAuthorityService) recordInvocation(key string, args []interface{}) {
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
