// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type StagedDirectorConfigService struct {
	GetStagedDirectorAvailabilityZonesStub        func() (api.AvailabilityZonesOutput, error)
	getStagedDirectorAvailabilityZonesMutex       sync.RWMutex
	getStagedDirectorAvailabilityZonesArgsForCall []struct {
	}
	getStagedDirectorAvailabilityZonesReturns struct {
		result1 api.AvailabilityZonesOutput
		result2 error
	}
	getStagedDirectorAvailabilityZonesReturnsOnCall map[int]struct {
		result1 api.AvailabilityZonesOutput
		result2 error
	}
	GetStagedDirectorIaasConfigurationsStub        func(bool) (map[string][]map[string]interface{}, error)
	getStagedDirectorIaasConfigurationsMutex       sync.RWMutex
	getStagedDirectorIaasConfigurationsArgsForCall []struct {
		arg1 bool
	}
	getStagedDirectorIaasConfigurationsReturns struct {
		result1 map[string][]map[string]interface{}
		result2 error
	}
	getStagedDirectorIaasConfigurationsReturnsOnCall map[int]struct {
		result1 map[string][]map[string]interface{}
		result2 error
	}
	GetStagedDirectorNetworksStub        func() (api.NetworksConfigurationOutput, error)
	getStagedDirectorNetworksMutex       sync.RWMutex
	getStagedDirectorNetworksArgsForCall []struct {
	}
	getStagedDirectorNetworksReturns struct {
		result1 api.NetworksConfigurationOutput
		result2 error
	}
	getStagedDirectorNetworksReturnsOnCall map[int]struct {
		result1 api.NetworksConfigurationOutput
		result2 error
	}
	GetStagedDirectorPropertiesStub        func(bool) (map[string]interface{}, error)
	getStagedDirectorPropertiesMutex       sync.RWMutex
	getStagedDirectorPropertiesArgsForCall []struct {
		arg1 bool
	}
	getStagedDirectorPropertiesReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	getStagedDirectorPropertiesReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	GetStagedProductByNameStub        func(string) (api.StagedProductsFindOutput, error)
	getStagedProductByNameMutex       sync.RWMutex
	getStagedProductByNameArgsForCall []struct {
		arg1 string
	}
	getStagedProductByNameReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	getStagedProductByNameReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	GetStagedProductJobResourceConfigStub        func(string, string) (api.JobProperties, error)
	getStagedProductJobResourceConfigMutex       sync.RWMutex
	getStagedProductJobResourceConfigArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getStagedProductJobResourceConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	getStagedProductJobResourceConfigReturnsOnCall map[int]struct {
		result1 api.JobProperties
		result2 error
	}
	GetStagedProductNetworksAndAZsStub        func(string) (map[string]interface{}, error)
	getStagedProductNetworksAndAZsMutex       sync.RWMutex
	getStagedProductNetworksAndAZsArgsForCall []struct {
		arg1 string
	}
	getStagedProductNetworksAndAZsReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	getStagedProductNetworksAndAZsReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	ListStagedProductJobsStub        func(string) (map[string]string, error)
	listStagedProductJobsMutex       sync.RWMutex
	listStagedProductJobsArgsForCall []struct {
		arg1 string
	}
	listStagedProductJobsReturns struct {
		result1 map[string]string
		result2 error
	}
	listStagedProductJobsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	ListStagedVMExtensionsStub        func() ([]api.VMExtension, error)
	listStagedVMExtensionsMutex       sync.RWMutex
	listStagedVMExtensionsArgsForCall []struct {
	}
	listStagedVMExtensionsReturns struct {
		result1 []api.VMExtension
		result2 error
	}
	listStagedVMExtensionsReturnsOnCall map[int]struct {
		result1 []api.VMExtension
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAvailabilityZones() (api.AvailabilityZonesOutput, error) {
	fake.getStagedDirectorAvailabilityZonesMutex.Lock()
	ret, specificReturn := fake.getStagedDirectorAvailabilityZonesReturnsOnCall[len(fake.getStagedDirectorAvailabilityZonesArgsForCall)]
	fake.getStagedDirectorAvailabilityZonesArgsForCall = append(fake.getStagedDirectorAvailabilityZonesArgsForCall, struct {
	}{})
	fake.recordInvocation("GetStagedDirectorAvailabilityZones", []interface{}{})
	fake.getStagedDirectorAvailabilityZonesMutex.Unlock()
	if fake.GetStagedDirectorAvailabilityZonesStub != nil {
		return fake.GetStagedDirectorAvailabilityZonesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedDirectorAvailabilityZonesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAvailabilityZonesCallCount() int {
	fake.getStagedDirectorAvailabilityZonesMutex.RLock()
	defer fake.getStagedDirectorAvailabilityZonesMutex.RUnlock()
	return len(fake.getStagedDirectorAvailabilityZonesArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAvailabilityZonesCalls(stub func() (api.AvailabilityZonesOutput, error)) {
	fake.getStagedDirectorAvailabilityZonesMutex.Lock()
	defer fake.getStagedDirectorAvailabilityZonesMutex.Unlock()
	fake.GetStagedDirectorAvailabilityZonesStub = stub
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAvailabilityZonesReturns(result1 api.AvailabilityZonesOutput, result2 error) {
	fake.getStagedDirectorAvailabilityZonesMutex.Lock()
	defer fake.getStagedDirectorAvailabilityZonesMutex.Unlock()
	fake.GetStagedDirectorAvailabilityZonesStub = nil
	fake.getStagedDirectorAvailabilityZonesReturns = struct {
		result1 api.AvailabilityZonesOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAvailabilityZonesReturnsOnCall(i int, result1 api.AvailabilityZonesOutput, result2 error) {
	fake.getStagedDirectorAvailabilityZonesMutex.Lock()
	defer fake.getStagedDirectorAvailabilityZonesMutex.Unlock()
	fake.GetStagedDirectorAvailabilityZonesStub = nil
	if fake.getStagedDirectorAvailabilityZonesReturnsOnCall == nil {
		fake.getStagedDirectorAvailabilityZonesReturnsOnCall = make(map[int]struct {
			result1 api.AvailabilityZonesOutput
			result2 error
		})
	}
	fake.getStagedDirectorAvailabilityZonesReturnsOnCall[i] = struct {
		result1 api.AvailabilityZonesOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorIaasConfigurations(arg1 bool) (map[string][]map[string]interface{}, error) {
	fake.getStagedDirectorIaasConfigurationsMutex.Lock()
	ret, specificReturn := fake.getStagedDirectorIaasConfigurationsReturnsOnCall[len(fake.getStagedDirectorIaasConfigurationsArgsForCall)]
	fake.getStagedDirectorIaasConfigurationsArgsForCall = append(fake.getStagedDirectorIaasConfigurationsArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("GetStagedDirectorIaasConfigurations", []interface{}{arg1})
	fake.getStagedDirectorIaasConfigurationsMutex.Unlock()
	if fake.GetStagedDirectorIaasConfigurationsStub != nil {
		return fake.GetStagedDirectorIaasConfigurationsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedDirectorIaasConfigurationsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedDirectorIaasConfigurationsCallCount() int {
	fake.getStagedDirectorIaasConfigurationsMutex.RLock()
	defer fake.getStagedDirectorIaasConfigurationsMutex.RUnlock()
	return len(fake.getStagedDirectorIaasConfigurationsArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedDirectorIaasConfigurationsCalls(stub func(bool) (map[string][]map[string]interface{}, error)) {
	fake.getStagedDirectorIaasConfigurationsMutex.Lock()
	defer fake.getStagedDirectorIaasConfigurationsMutex.Unlock()
	fake.GetStagedDirectorIaasConfigurationsStub = stub
}

func (fake *StagedDirectorConfigService) GetStagedDirectorIaasConfigurationsArgsForCall(i int) bool {
	fake.getStagedDirectorIaasConfigurationsMutex.RLock()
	defer fake.getStagedDirectorIaasConfigurationsMutex.RUnlock()
	argsForCall := fake.getStagedDirectorIaasConfigurationsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedDirectorConfigService) GetStagedDirectorIaasConfigurationsReturns(result1 map[string][]map[string]interface{}, result2 error) {
	fake.getStagedDirectorIaasConfigurationsMutex.Lock()
	defer fake.getStagedDirectorIaasConfigurationsMutex.Unlock()
	fake.GetStagedDirectorIaasConfigurationsStub = nil
	fake.getStagedDirectorIaasConfigurationsReturns = struct {
		result1 map[string][]map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorIaasConfigurationsReturnsOnCall(i int, result1 map[string][]map[string]interface{}, result2 error) {
	fake.getStagedDirectorIaasConfigurationsMutex.Lock()
	defer fake.getStagedDirectorIaasConfigurationsMutex.Unlock()
	fake.GetStagedDirectorIaasConfigurationsStub = nil
	if fake.getStagedDirectorIaasConfigurationsReturnsOnCall == nil {
		fake.getStagedDirectorIaasConfigurationsReturnsOnCall = make(map[int]struct {
			result1 map[string][]map[string]interface{}
			result2 error
		})
	}
	fake.getStagedDirectorIaasConfigurationsReturnsOnCall[i] = struct {
		result1 map[string][]map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorNetworks() (api.NetworksConfigurationOutput, error) {
	fake.getStagedDirectorNetworksMutex.Lock()
	ret, specificReturn := fake.getStagedDirectorNetworksReturnsOnCall[len(fake.getStagedDirectorNetworksArgsForCall)]
	fake.getStagedDirectorNetworksArgsForCall = append(fake.getStagedDirectorNetworksArgsForCall, struct {
	}{})
	fake.recordInvocation("GetStagedDirectorNetworks", []interface{}{})
	fake.getStagedDirectorNetworksMutex.Unlock()
	if fake.GetStagedDirectorNetworksStub != nil {
		return fake.GetStagedDirectorNetworksStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedDirectorNetworksReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedDirectorNetworksCallCount() int {
	fake.getStagedDirectorNetworksMutex.RLock()
	defer fake.getStagedDirectorNetworksMutex.RUnlock()
	return len(fake.getStagedDirectorNetworksArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedDirectorNetworksCalls(stub func() (api.NetworksConfigurationOutput, error)) {
	fake.getStagedDirectorNetworksMutex.Lock()
	defer fake.getStagedDirectorNetworksMutex.Unlock()
	fake.GetStagedDirectorNetworksStub = stub
}

func (fake *StagedDirectorConfigService) GetStagedDirectorNetworksReturns(result1 api.NetworksConfigurationOutput, result2 error) {
	fake.getStagedDirectorNetworksMutex.Lock()
	defer fake.getStagedDirectorNetworksMutex.Unlock()
	fake.GetStagedDirectorNetworksStub = nil
	fake.getStagedDirectorNetworksReturns = struct {
		result1 api.NetworksConfigurationOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorNetworksReturnsOnCall(i int, result1 api.NetworksConfigurationOutput, result2 error) {
	fake.getStagedDirectorNetworksMutex.Lock()
	defer fake.getStagedDirectorNetworksMutex.Unlock()
	fake.GetStagedDirectorNetworksStub = nil
	if fake.getStagedDirectorNetworksReturnsOnCall == nil {
		fake.getStagedDirectorNetworksReturnsOnCall = make(map[int]struct {
			result1 api.NetworksConfigurationOutput
			result2 error
		})
	}
	fake.getStagedDirectorNetworksReturnsOnCall[i] = struct {
		result1 api.NetworksConfigurationOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorProperties(arg1 bool) (map[string]interface{}, error) {
	fake.getStagedDirectorPropertiesMutex.Lock()
	ret, specificReturn := fake.getStagedDirectorPropertiesReturnsOnCall[len(fake.getStagedDirectorPropertiesArgsForCall)]
	fake.getStagedDirectorPropertiesArgsForCall = append(fake.getStagedDirectorPropertiesArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("GetStagedDirectorProperties", []interface{}{arg1})
	fake.getStagedDirectorPropertiesMutex.Unlock()
	if fake.GetStagedDirectorPropertiesStub != nil {
		return fake.GetStagedDirectorPropertiesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedDirectorPropertiesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedDirectorPropertiesCallCount() int {
	fake.getStagedDirectorPropertiesMutex.RLock()
	defer fake.getStagedDirectorPropertiesMutex.RUnlock()
	return len(fake.getStagedDirectorPropertiesArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedDirectorPropertiesCalls(stub func(bool) (map[string]interface{}, error)) {
	fake.getStagedDirectorPropertiesMutex.Lock()
	defer fake.getStagedDirectorPropertiesMutex.Unlock()
	fake.GetStagedDirectorPropertiesStub = stub
}

func (fake *StagedDirectorConfigService) GetStagedDirectorPropertiesArgsForCall(i int) bool {
	fake.getStagedDirectorPropertiesMutex.RLock()
	defer fake.getStagedDirectorPropertiesMutex.RUnlock()
	argsForCall := fake.getStagedDirectorPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedDirectorConfigService) GetStagedDirectorPropertiesReturns(result1 map[string]interface{}, result2 error) {
	fake.getStagedDirectorPropertiesMutex.Lock()
	defer fake.getStagedDirectorPropertiesMutex.Unlock()
	fake.GetStagedDirectorPropertiesStub = nil
	fake.getStagedDirectorPropertiesReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorPropertiesReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.getStagedDirectorPropertiesMutex.Lock()
	defer fake.getStagedDirectorPropertiesMutex.Unlock()
	fake.GetStagedDirectorPropertiesStub = nil
	if fake.getStagedDirectorPropertiesReturnsOnCall == nil {
		fake.getStagedDirectorPropertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.getStagedDirectorPropertiesReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductByName(arg1 string) (api.StagedProductsFindOutput, error) {
	fake.getStagedProductByNameMutex.Lock()
	ret, specificReturn := fake.getStagedProductByNameReturnsOnCall[len(fake.getStagedProductByNameArgsForCall)]
	fake.getStagedProductByNameArgsForCall = append(fake.getStagedProductByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStagedProductByName", []interface{}{arg1})
	fake.getStagedProductByNameMutex.Unlock()
	if fake.GetStagedProductByNameStub != nil {
		return fake.GetStagedProductByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameCalls(stub func(string) (api.StagedProductsFindOutput, error)) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = stub
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	argsForCall := fake.getStagedProductByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	if fake.getStagedProductByNameReturnsOnCall == nil {
		fake.getStagedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.getStagedProductByNameReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductJobResourceConfig(arg1 string, arg2 string) (api.JobProperties, error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	ret, specificReturn := fake.getStagedProductJobResourceConfigReturnsOnCall[len(fake.getStagedProductJobResourceConfigArgsForCall)]
	fake.getStagedProductJobResourceConfigArgsForCall = append(fake.getStagedProductJobResourceConfigArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetStagedProductJobResourceConfig", []interface{}{arg1, arg2})
	fake.getStagedProductJobResourceConfigMutex.Unlock()
	if fake.GetStagedProductJobResourceConfigStub != nil {
		return fake.GetStagedProductJobResourceConfigStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductJobResourceConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedProductJobResourceConfigCallCount() int {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	return len(fake.getStagedProductJobResourceConfigArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedProductJobResourceConfigCalls(stub func(string, string) (api.JobProperties, error)) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	defer fake.getStagedProductJobResourceConfigMutex.Unlock()
	fake.GetStagedProductJobResourceConfigStub = stub
}

func (fake *StagedDirectorConfigService) GetStagedProductJobResourceConfigArgsForCall(i int) (string, string) {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	argsForCall := fake.getStagedProductJobResourceConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *StagedDirectorConfigService) GetStagedProductJobResourceConfigReturns(result1 api.JobProperties, result2 error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	defer fake.getStagedProductJobResourceConfigMutex.Unlock()
	fake.GetStagedProductJobResourceConfigStub = nil
	fake.getStagedProductJobResourceConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductJobResourceConfigReturnsOnCall(i int, result1 api.JobProperties, result2 error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	defer fake.getStagedProductJobResourceConfigMutex.Unlock()
	fake.GetStagedProductJobResourceConfigStub = nil
	if fake.getStagedProductJobResourceConfigReturnsOnCall == nil {
		fake.getStagedProductJobResourceConfigReturnsOnCall = make(map[int]struct {
			result1 api.JobProperties
			result2 error
		})
	}
	fake.getStagedProductJobResourceConfigReturnsOnCall[i] = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductNetworksAndAZs(arg1 string) (map[string]interface{}, error) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	ret, specificReturn := fake.getStagedProductNetworksAndAZsReturnsOnCall[len(fake.getStagedProductNetworksAndAZsArgsForCall)]
	fake.getStagedProductNetworksAndAZsArgsForCall = append(fake.getStagedProductNetworksAndAZsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStagedProductNetworksAndAZs", []interface{}{arg1})
	fake.getStagedProductNetworksAndAZsMutex.Unlock()
	if fake.GetStagedProductNetworksAndAZsStub != nil {
		return fake.GetStagedProductNetworksAndAZsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductNetworksAndAZsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedProductNetworksAndAZsCallCount() int {
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	return len(fake.getStagedProductNetworksAndAZsArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedProductNetworksAndAZsCalls(stub func(string) (map[string]interface{}, error)) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	defer fake.getStagedProductNetworksAndAZsMutex.Unlock()
	fake.GetStagedProductNetworksAndAZsStub = stub
}

func (fake *StagedDirectorConfigService) GetStagedProductNetworksAndAZsArgsForCall(i int) string {
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	argsForCall := fake.getStagedProductNetworksAndAZsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedDirectorConfigService) GetStagedProductNetworksAndAZsReturns(result1 map[string]interface{}, result2 error) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	defer fake.getStagedProductNetworksAndAZsMutex.Unlock()
	fake.GetStagedProductNetworksAndAZsStub = nil
	fake.getStagedProductNetworksAndAZsReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductNetworksAndAZsReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	defer fake.getStagedProductNetworksAndAZsMutex.Unlock()
	fake.GetStagedProductNetworksAndAZsStub = nil
	if fake.getStagedProductNetworksAndAZsReturnsOnCall == nil {
		fake.getStagedProductNetworksAndAZsReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.getStagedProductNetworksAndAZsReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) ListStagedProductJobs(arg1 string) (map[string]string, error) {
	fake.listStagedProductJobsMutex.Lock()
	ret, specificReturn := fake.listStagedProductJobsReturnsOnCall[len(fake.listStagedProductJobsArgsForCall)]
	fake.listStagedProductJobsArgsForCall = append(fake.listStagedProductJobsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListStagedProductJobs", []interface{}{arg1})
	fake.listStagedProductJobsMutex.Unlock()
	if fake.ListStagedProductJobsStub != nil {
		return fake.ListStagedProductJobsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listStagedProductJobsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) ListStagedProductJobsCallCount() int {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	return len(fake.listStagedProductJobsArgsForCall)
}

func (fake *StagedDirectorConfigService) ListStagedProductJobsCalls(stub func(string) (map[string]string, error)) {
	fake.listStagedProductJobsMutex.Lock()
	defer fake.listStagedProductJobsMutex.Unlock()
	fake.ListStagedProductJobsStub = stub
}

func (fake *StagedDirectorConfigService) ListStagedProductJobsArgsForCall(i int) string {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	argsForCall := fake.listStagedProductJobsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedDirectorConfigService) ListStagedProductJobsReturns(result1 map[string]string, result2 error) {
	fake.listStagedProductJobsMutex.Lock()
	defer fake.listStagedProductJobsMutex.Unlock()
	fake.ListStagedProductJobsStub = nil
	fake.listStagedProductJobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) ListStagedProductJobsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.listStagedProductJobsMutex.Lock()
	defer fake.listStagedProductJobsMutex.Unlock()
	fake.ListStagedProductJobsStub = nil
	if fake.listStagedProductJobsReturnsOnCall == nil {
		fake.listStagedProductJobsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.listStagedProductJobsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) ListStagedVMExtensions() ([]api.VMExtension, error) {
	fake.listStagedVMExtensionsMutex.Lock()
	ret, specificReturn := fake.listStagedVMExtensionsReturnsOnCall[len(fake.listStagedVMExtensionsArgsForCall)]
	fake.listStagedVMExtensionsArgsForCall = append(fake.listStagedVMExtensionsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListStagedVMExtensions", []interface{}{})
	fake.listStagedVMExtensionsMutex.Unlock()
	if fake.ListStagedVMExtensionsStub != nil {
		return fake.ListStagedVMExtensionsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listStagedVMExtensionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedDirectorConfigService) ListStagedVMExtensionsCallCount() int {
	fake.listStagedVMExtensionsMutex.RLock()
	defer fake.listStagedVMExtensionsMutex.RUnlock()
	return len(fake.listStagedVMExtensionsArgsForCall)
}

func (fake *StagedDirectorConfigService) ListStagedVMExtensionsCalls(stub func() ([]api.VMExtension, error)) {
	fake.listStagedVMExtensionsMutex.Lock()
	defer fake.listStagedVMExtensionsMutex.Unlock()
	fake.ListStagedVMExtensionsStub = stub
}

func (fake *StagedDirectorConfigService) ListStagedVMExtensionsReturns(result1 []api.VMExtension, result2 error) {
	fake.listStagedVMExtensionsMutex.Lock()
	defer fake.listStagedVMExtensionsMutex.Unlock()
	fake.ListStagedVMExtensionsStub = nil
	fake.listStagedVMExtensionsReturns = struct {
		result1 []api.VMExtension
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) ListStagedVMExtensionsReturnsOnCall(i int, result1 []api.VMExtension, result2 error) {
	fake.listStagedVMExtensionsMutex.Lock()
	defer fake.listStagedVMExtensionsMutex.Unlock()
	fake.ListStagedVMExtensionsStub = nil
	if fake.listStagedVMExtensionsReturnsOnCall == nil {
		fake.listStagedVMExtensionsReturnsOnCall = make(map[int]struct {
			result1 []api.VMExtension
			result2 error
		})
	}
	fake.listStagedVMExtensionsReturnsOnCall[i] = struct {
		result1 []api.VMExtension
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStagedDirectorAvailabilityZonesMutex.RLock()
	defer fake.getStagedDirectorAvailabilityZonesMutex.RUnlock()
	fake.getStagedDirectorIaasConfigurationsMutex.RLock()
	defer fake.getStagedDirectorIaasConfigurationsMutex.RUnlock()
	fake.getStagedDirectorNetworksMutex.RLock()
	defer fake.getStagedDirectorNetworksMutex.RUnlock()
	fake.getStagedDirectorPropertiesMutex.RLock()
	defer fake.getStagedDirectorPropertiesMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	fake.listStagedVMExtensionsMutex.RLock()
	defer fake.listStagedVMExtensionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StagedDirectorConfigService) recordInvocation(key string, args []interface{}) {
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
