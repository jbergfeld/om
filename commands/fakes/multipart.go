// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/formcontent"
)

type Multipart struct {
	FinalizeStub        func() formcontent.ContentSubmission
	finalizeMutex       sync.RWMutex
	finalizeArgsForCall []struct{}
	finalizeReturns     struct {
		result1 formcontent.ContentSubmission
	}
	finalizeReturnsOnCall map[int]struct {
		result1 formcontent.ContentSubmission
	}
	ResetStub          func()
	resetMutex         sync.RWMutex
	resetArgsForCall   []struct{}
	AddFileStub        func(key, path string) error
	addFileMutex       sync.RWMutex
	addFileArgsForCall []struct {
		key  string
		path string
	}
	addFileReturns struct {
		result1 error
	}
	addFileReturnsOnCall map[int]struct {
		result1 error
	}
	AddFieldStub        func(key, value string) error
	addFieldMutex       sync.RWMutex
	addFieldArgsForCall []struct {
		key   string
		value string
	}
	addFieldReturns struct {
		result1 error
	}
	addFieldReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Multipart) Finalize() formcontent.ContentSubmission {
	fake.finalizeMutex.Lock()
	ret, specificReturn := fake.finalizeReturnsOnCall[len(fake.finalizeArgsForCall)]
	fake.finalizeArgsForCall = append(fake.finalizeArgsForCall, struct{}{})
	fake.recordInvocation("Finalize", []interface{}{})
	fake.finalizeMutex.Unlock()
	if fake.FinalizeStub != nil {
		return fake.FinalizeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.finalizeReturns.result1
}

func (fake *Multipart) FinalizeCallCount() int {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return len(fake.finalizeArgsForCall)
}

func (fake *Multipart) FinalizeReturns(result1 formcontent.ContentSubmission) {
	fake.FinalizeStub = nil
	fake.finalizeReturns = struct {
		result1 formcontent.ContentSubmission
	}{result1}
}

func (fake *Multipart) FinalizeReturnsOnCall(i int, result1 formcontent.ContentSubmission) {
	fake.FinalizeStub = nil
	if fake.finalizeReturnsOnCall == nil {
		fake.finalizeReturnsOnCall = make(map[int]struct {
			result1 formcontent.ContentSubmission
		})
	}
	fake.finalizeReturnsOnCall[i] = struct {
		result1 formcontent.ContentSubmission
	}{result1}
}

func (fake *Multipart) Reset() {
	fake.resetMutex.Lock()
	fake.resetArgsForCall = append(fake.resetArgsForCall, struct{}{})
	fake.recordInvocation("Reset", []interface{}{})
	fake.resetMutex.Unlock()
	if fake.ResetStub != nil {
		fake.ResetStub()
	}
}

func (fake *Multipart) ResetCallCount() int {
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return len(fake.resetArgsForCall)
}

func (fake *Multipart) AddFile(key string, path string) error {
	fake.addFileMutex.Lock()
	ret, specificReturn := fake.addFileReturnsOnCall[len(fake.addFileArgsForCall)]
	fake.addFileArgsForCall = append(fake.addFileArgsForCall, struct {
		key  string
		path string
	}{key, path})
	fake.recordInvocation("AddFile", []interface{}{key, path})
	fake.addFileMutex.Unlock()
	if fake.AddFileStub != nil {
		return fake.AddFileStub(key, path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addFileReturns.result1
}

func (fake *Multipart) AddFileCallCount() int {
	fake.addFileMutex.RLock()
	defer fake.addFileMutex.RUnlock()
	return len(fake.addFileArgsForCall)
}

func (fake *Multipart) AddFileArgsForCall(i int) (string, string) {
	fake.addFileMutex.RLock()
	defer fake.addFileMutex.RUnlock()
	return fake.addFileArgsForCall[i].key, fake.addFileArgsForCall[i].path
}

func (fake *Multipart) AddFileReturns(result1 error) {
	fake.AddFileStub = nil
	fake.addFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *Multipart) AddFileReturnsOnCall(i int, result1 error) {
	fake.AddFileStub = nil
	if fake.addFileReturnsOnCall == nil {
		fake.addFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Multipart) AddField(key string, value string) error {
	fake.addFieldMutex.Lock()
	ret, specificReturn := fake.addFieldReturnsOnCall[len(fake.addFieldArgsForCall)]
	fake.addFieldArgsForCall = append(fake.addFieldArgsForCall, struct {
		key   string
		value string
	}{key, value})
	fake.recordInvocation("AddField", []interface{}{key, value})
	fake.addFieldMutex.Unlock()
	if fake.AddFieldStub != nil {
		return fake.AddFieldStub(key, value)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addFieldReturns.result1
}

func (fake *Multipart) AddFieldCallCount() int {
	fake.addFieldMutex.RLock()
	defer fake.addFieldMutex.RUnlock()
	return len(fake.addFieldArgsForCall)
}

func (fake *Multipart) AddFieldArgsForCall(i int) (string, string) {
	fake.addFieldMutex.RLock()
	defer fake.addFieldMutex.RUnlock()
	return fake.addFieldArgsForCall[i].key, fake.addFieldArgsForCall[i].value
}

func (fake *Multipart) AddFieldReturns(result1 error) {
	fake.AddFieldStub = nil
	fake.addFieldReturns = struct {
		result1 error
	}{result1}
}

func (fake *Multipart) AddFieldReturnsOnCall(i int, result1 error) {
	fake.AddFieldStub = nil
	if fake.addFieldReturnsOnCall == nil {
		fake.addFieldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFieldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Multipart) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	fake.addFileMutex.RLock()
	defer fake.addFileMutex.RUnlock()
	fake.addFieldMutex.RLock()
	defer fake.addFieldMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Multipart) recordInvocation(key string, args []interface{}) {
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
