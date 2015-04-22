// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/web/getbuilds"
)

type FakeBuildsDB struct {
	GetAllBuildsStub        func() ([]db.Build, error)
	getAllBuildsMutex       sync.RWMutex
	getAllBuildsArgsForCall []struct{}
	getAllBuildsReturns     struct {
		result1 []db.Build
		result2 error
	}
}

func (fake *FakeBuildsDB) GetAllBuilds() ([]db.Build, error) {
	fake.getAllBuildsMutex.Lock()
	fake.getAllBuildsArgsForCall = append(fake.getAllBuildsArgsForCall, struct{}{})
	fake.getAllBuildsMutex.Unlock()
	if fake.GetAllBuildsStub != nil {
		return fake.GetAllBuildsStub()
	} else {
		return fake.getAllBuildsReturns.result1, fake.getAllBuildsReturns.result2
	}
}

func (fake *FakeBuildsDB) GetAllBuildsCallCount() int {
	fake.getAllBuildsMutex.RLock()
	defer fake.getAllBuildsMutex.RUnlock()
	return len(fake.getAllBuildsArgsForCall)
}

func (fake *FakeBuildsDB) GetAllBuildsReturns(result1 []db.Build, result2 error) {
	fake.GetAllBuildsStub = nil
	fake.getAllBuildsReturns = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

var _ getbuilds.BuildsDB = new(FakeBuildsDB)
