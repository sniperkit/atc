// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/resource"
	"github.com/pivotal-golang/lager"
)

type FakeFetchSourceProviderFactory struct {
	NewFetchSourceProviderStub        func(logger lager.Logger, session resource.Session, tags atc.Tags, teamName string, resourceTypes atc.ResourceTypes, cacheIdentifier resource.CacheIdentifier, resourceOptions resource.ResourceOptions, containerCreator resource.FetchContainerCreator) resource.FetchSourceProvider
	newFetchSourceProviderMutex       sync.RWMutex
	newFetchSourceProviderArgsForCall []struct {
		logger           lager.Logger
		session          resource.Session
		tags             atc.Tags
		teamName         string
		resourceTypes    atc.ResourceTypes
		cacheIdentifier  resource.CacheIdentifier
		resourceOptions  resource.ResourceOptions
		containerCreator resource.FetchContainerCreator
	}
	newFetchSourceProviderReturns struct {
		result1 resource.FetchSourceProvider
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProvider(logger lager.Logger, session resource.Session, tags atc.Tags, teamName string, resourceTypes atc.ResourceTypes, cacheIdentifier resource.CacheIdentifier, resourceOptions resource.ResourceOptions, containerCreator resource.FetchContainerCreator) resource.FetchSourceProvider {
	fake.newFetchSourceProviderMutex.Lock()
	fake.newFetchSourceProviderArgsForCall = append(fake.newFetchSourceProviderArgsForCall, struct {
		logger           lager.Logger
		session          resource.Session
		tags             atc.Tags
		teamName         string
		resourceTypes    atc.ResourceTypes
		cacheIdentifier  resource.CacheIdentifier
		resourceOptions  resource.ResourceOptions
		containerCreator resource.FetchContainerCreator
	}{logger, session, tags, teamName, resourceTypes, cacheIdentifier, resourceOptions, containerCreator})
	fake.recordInvocation("NewFetchSourceProvider", []interface{}{logger, session, tags, teamName, resourceTypes, cacheIdentifier, resourceOptions, containerCreator})
	fake.newFetchSourceProviderMutex.Unlock()
	if fake.NewFetchSourceProviderStub != nil {
		return fake.NewFetchSourceProviderStub(logger, session, tags, teamName, resourceTypes, cacheIdentifier, resourceOptions, containerCreator)
	} else {
		return fake.newFetchSourceProviderReturns.result1
	}
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProviderCallCount() int {
	fake.newFetchSourceProviderMutex.RLock()
	defer fake.newFetchSourceProviderMutex.RUnlock()
	return len(fake.newFetchSourceProviderArgsForCall)
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProviderArgsForCall(i int) (lager.Logger, resource.Session, atc.Tags, string, atc.ResourceTypes, resource.CacheIdentifier, resource.ResourceOptions, resource.FetchContainerCreator) {
	fake.newFetchSourceProviderMutex.RLock()
	defer fake.newFetchSourceProviderMutex.RUnlock()
	return fake.newFetchSourceProviderArgsForCall[i].logger, fake.newFetchSourceProviderArgsForCall[i].session, fake.newFetchSourceProviderArgsForCall[i].tags, fake.newFetchSourceProviderArgsForCall[i].teamName, fake.newFetchSourceProviderArgsForCall[i].resourceTypes, fake.newFetchSourceProviderArgsForCall[i].cacheIdentifier, fake.newFetchSourceProviderArgsForCall[i].resourceOptions, fake.newFetchSourceProviderArgsForCall[i].containerCreator
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProviderReturns(result1 resource.FetchSourceProvider) {
	fake.NewFetchSourceProviderStub = nil
	fake.newFetchSourceProviderReturns = struct {
		result1 resource.FetchSourceProvider
	}{result1}
}

func (fake *FakeFetchSourceProviderFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newFetchSourceProviderMutex.RLock()
	defer fake.newFetchSourceProviderMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFetchSourceProviderFactory) recordInvocation(key string, args []interface{}) {
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

var _ resource.FetchSourceProviderFactory = new(FakeFetchSourceProviderFactory)
