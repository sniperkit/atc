// This file was generated by counterfeiter
package workerfakes

import (
	"io"
	"sync"
	"time"

	"github.com/concourse/atc/worker"
	"github.com/concourse/baggageclaim"
)

type FakeVolume struct {
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns     struct {
		result1 string
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct{}
	pathReturns     struct {
		result1 string
	}
	SetTTLStub        func(time.Duration) error
	setTTLMutex       sync.RWMutex
	setTTLArgsForCall []struct {
		arg1 time.Duration
	}
	setTTLReturns struct {
		result1 error
	}
	SetPropertyStub        func(key string, value string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		key   string
		value string
	}
	setPropertyReturns struct {
		result1 error
	}
	StreamInStub        func(path string, tarStream io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		path      string
		tarStream io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	StreamOutStub        func(path string) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		path string
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	ExpirationStub        func() (time.Duration, time.Time, error)
	expirationMutex       sync.RWMutex
	expirationArgsForCall []struct{}
	expirationReturns     struct {
		result1 time.Duration
		result2 time.Time
		result3 error
	}
	PropertiesStub        func() (baggageclaim.VolumeProperties, error)
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct{}
	propertiesReturns     struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	ReleaseStub        func(*time.Duration)
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
		arg1 *time.Duration
	}
	SizeInBytesStub        func() (int64, error)
	sizeInBytesMutex       sync.RWMutex
	sizeInBytesArgsForCall []struct{}
	sizeInBytesReturns     struct {
		result1 int64
		result2 error
	}
	DestroyStub        func()
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct{}
	invocations        map[string][][]interface{}
	invocationsMutex   sync.RWMutex
}

func (fake *FakeVolume) Handle() string {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	} else {
		return fake.handleReturns.result1
	}
}

func (fake *FakeVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeVolume) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Path() string {
	fake.pathMutex.Lock()
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct{}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	} else {
		return fake.pathReturns.result1
	}
}

func (fake *FakeVolume) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeVolume) PathReturns(result1 string) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) SetTTL(arg1 time.Duration) error {
	fake.setTTLMutex.Lock()
	fake.setTTLArgsForCall = append(fake.setTTLArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.recordInvocation("SetTTL", []interface{}{arg1})
	fake.setTTLMutex.Unlock()
	if fake.SetTTLStub != nil {
		return fake.SetTTLStub(arg1)
	} else {
		return fake.setTTLReturns.result1
	}
}

func (fake *FakeVolume) SetTTLCallCount() int {
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	return len(fake.setTTLArgsForCall)
}

func (fake *FakeVolume) SetTTLArgsForCall(i int) time.Duration {
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	return fake.setTTLArgsForCall[i].arg1
}

func (fake *FakeVolume) SetTTLReturns(result1 error) {
	fake.SetTTLStub = nil
	fake.setTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetProperty(key string, value string) error {
	fake.setPropertyMutex.Lock()
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		key   string
		value string
	}{key, value})
	fake.recordInvocation("SetProperty", []interface{}{key, value})
	fake.setPropertyMutex.Unlock()
	if fake.SetPropertyStub != nil {
		return fake.SetPropertyStub(key, value)
	} else {
		return fake.setPropertyReturns.result1
	}
}

func (fake *FakeVolume) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeVolume) SetPropertyArgsForCall(i int) (string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return fake.setPropertyArgsForCall[i].key, fake.setPropertyArgsForCall[i].value
}

func (fake *FakeVolume) SetPropertyReturns(result1 error) {
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamIn(path string, tarStream io.Reader) error {
	fake.streamInMutex.Lock()
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		path      string
		tarStream io.Reader
	}{path, tarStream})
	fake.recordInvocation("StreamIn", []interface{}{path, tarStream})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(path, tarStream)
	} else {
		return fake.streamInReturns.result1
	}
}

func (fake *FakeVolume) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeVolume) StreamInArgsForCall(i int) (string, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return fake.streamInArgsForCall[i].path, fake.streamInArgsForCall[i].tarStream
}

func (fake *FakeVolume) StreamInReturns(result1 error) {
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamOut(path string) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("StreamOut", []interface{}{path})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(path)
	} else {
		return fake.streamOutReturns.result1, fake.streamOutReturns.result2
	}
}

func (fake *FakeVolume) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeVolume) StreamOutArgsForCall(i int) string {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return fake.streamOutArgsForCall[i].path
}

func (fake *FakeVolume) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) Expiration() (time.Duration, time.Time, error) {
	fake.expirationMutex.Lock()
	fake.expirationArgsForCall = append(fake.expirationArgsForCall, struct{}{})
	fake.recordInvocation("Expiration", []interface{}{})
	fake.expirationMutex.Unlock()
	if fake.ExpirationStub != nil {
		return fake.ExpirationStub()
	} else {
		return fake.expirationReturns.result1, fake.expirationReturns.result2, fake.expirationReturns.result3
	}
}

func (fake *FakeVolume) ExpirationCallCount() int {
	fake.expirationMutex.RLock()
	defer fake.expirationMutex.RUnlock()
	return len(fake.expirationArgsForCall)
}

func (fake *FakeVolume) ExpirationReturns(result1 time.Duration, result2 time.Time, result3 error) {
	fake.ExpirationStub = nil
	fake.expirationReturns = struct {
		result1 time.Duration
		result2 time.Time
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolume) Properties() (baggageclaim.VolumeProperties, error) {
	fake.propertiesMutex.Lock()
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct{}{})
	fake.recordInvocation("Properties", []interface{}{})
	fake.propertiesMutex.Unlock()
	if fake.PropertiesStub != nil {
		return fake.PropertiesStub()
	} else {
		return fake.propertiesReturns.result1, fake.propertiesReturns.result2
	}
}

func (fake *FakeVolume) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *FakeVolume) PropertiesReturns(result1 baggageclaim.VolumeProperties, result2 error) {
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) Release(arg1 *time.Duration) {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
		arg1 *time.Duration
	}{arg1})
	fake.recordInvocation("Release", []interface{}{arg1})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		fake.ReleaseStub(arg1)
	}
}

func (fake *FakeVolume) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeVolume) ReleaseArgsForCall(i int) *time.Duration {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return fake.releaseArgsForCall[i].arg1
}

func (fake *FakeVolume) SizeInBytes() (int64, error) {
	fake.sizeInBytesMutex.Lock()
	fake.sizeInBytesArgsForCall = append(fake.sizeInBytesArgsForCall, struct{}{})
	fake.recordInvocation("SizeInBytes", []interface{}{})
	fake.sizeInBytesMutex.Unlock()
	if fake.SizeInBytesStub != nil {
		return fake.SizeInBytesStub()
	} else {
		return fake.sizeInBytesReturns.result1, fake.sizeInBytesReturns.result2
	}
}

func (fake *FakeVolume) SizeInBytesCallCount() int {
	fake.sizeInBytesMutex.RLock()
	defer fake.sizeInBytesMutex.RUnlock()
	return len(fake.sizeInBytesArgsForCall)
}

func (fake *FakeVolume) SizeInBytesReturns(result1 int64, result2 error) {
	fake.SizeInBytesStub = nil
	fake.sizeInBytesReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) Destroy() {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct{}{})
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		fake.DestroyStub()
	}
}

func (fake *FakeVolume) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.expirationMutex.RLock()
	defer fake.expirationMutex.RUnlock()
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	fake.sizeInBytesMutex.RLock()
	defer fake.sizeInBytesMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVolume) recordInvocation(key string, args []interface{}) {
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

var _ worker.Volume = new(FakeVolume)
