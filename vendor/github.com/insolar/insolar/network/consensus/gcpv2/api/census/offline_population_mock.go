package census

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/network/consensus/common/endpoints"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/profiles"
)

// OfflinePopulationMock implements OfflinePopulation
type OfflinePopulationMock struct {
	t minimock.Tester

	funcFindRegisteredProfile          func(identity endpoints.Inbound) (h1 profiles.Host)
	inspectFuncFindRegisteredProfile   func(identity endpoints.Inbound)
	afterFindRegisteredProfileCounter  uint64
	beforeFindRegisteredProfileCounter uint64
	FindRegisteredProfileMock          mOfflinePopulationMockFindRegisteredProfile
}

// NewOfflinePopulationMock returns a mock for OfflinePopulation
func NewOfflinePopulationMock(t minimock.Tester) *OfflinePopulationMock {
	m := &OfflinePopulationMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FindRegisteredProfileMock = mOfflinePopulationMockFindRegisteredProfile{mock: m}
	m.FindRegisteredProfileMock.callArgs = []*OfflinePopulationMockFindRegisteredProfileParams{}

	return m
}

type mOfflinePopulationMockFindRegisteredProfile struct {
	mock               *OfflinePopulationMock
	defaultExpectation *OfflinePopulationMockFindRegisteredProfileExpectation
	expectations       []*OfflinePopulationMockFindRegisteredProfileExpectation

	callArgs []*OfflinePopulationMockFindRegisteredProfileParams
	mutex    sync.RWMutex
}

// OfflinePopulationMockFindRegisteredProfileExpectation specifies expectation struct of the OfflinePopulation.FindRegisteredProfile
type OfflinePopulationMockFindRegisteredProfileExpectation struct {
	mock    *OfflinePopulationMock
	params  *OfflinePopulationMockFindRegisteredProfileParams
	results *OfflinePopulationMockFindRegisteredProfileResults
	Counter uint64
}

// OfflinePopulationMockFindRegisteredProfileParams contains parameters of the OfflinePopulation.FindRegisteredProfile
type OfflinePopulationMockFindRegisteredProfileParams struct {
	identity endpoints.Inbound
}

// OfflinePopulationMockFindRegisteredProfileResults contains results of the OfflinePopulation.FindRegisteredProfile
type OfflinePopulationMockFindRegisteredProfileResults struct {
	h1 profiles.Host
}

// Expect sets up expected params for OfflinePopulation.FindRegisteredProfile
func (mmFindRegisteredProfile *mOfflinePopulationMockFindRegisteredProfile) Expect(identity endpoints.Inbound) *mOfflinePopulationMockFindRegisteredProfile {
	if mmFindRegisteredProfile.mock.funcFindRegisteredProfile != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("OfflinePopulationMock.FindRegisteredProfile mock is already set by Set")
	}

	if mmFindRegisteredProfile.defaultExpectation == nil {
		mmFindRegisteredProfile.defaultExpectation = &OfflinePopulationMockFindRegisteredProfileExpectation{}
	}

	mmFindRegisteredProfile.defaultExpectation.params = &OfflinePopulationMockFindRegisteredProfileParams{identity}
	for _, e := range mmFindRegisteredProfile.expectations {
		if minimock.Equal(e.params, mmFindRegisteredProfile.defaultExpectation.params) {
			mmFindRegisteredProfile.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFindRegisteredProfile.defaultExpectation.params)
		}
	}

	return mmFindRegisteredProfile
}

// Inspect accepts an inspector function that has same arguments as the OfflinePopulation.FindRegisteredProfile
func (mmFindRegisteredProfile *mOfflinePopulationMockFindRegisteredProfile) Inspect(f func(identity endpoints.Inbound)) *mOfflinePopulationMockFindRegisteredProfile {
	if mmFindRegisteredProfile.mock.inspectFuncFindRegisteredProfile != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("Inspect function is already set for OfflinePopulationMock.FindRegisteredProfile")
	}

	mmFindRegisteredProfile.mock.inspectFuncFindRegisteredProfile = f

	return mmFindRegisteredProfile
}

// Return sets up results that will be returned by OfflinePopulation.FindRegisteredProfile
func (mmFindRegisteredProfile *mOfflinePopulationMockFindRegisteredProfile) Return(h1 profiles.Host) *OfflinePopulationMock {
	if mmFindRegisteredProfile.mock.funcFindRegisteredProfile != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("OfflinePopulationMock.FindRegisteredProfile mock is already set by Set")
	}

	if mmFindRegisteredProfile.defaultExpectation == nil {
		mmFindRegisteredProfile.defaultExpectation = &OfflinePopulationMockFindRegisteredProfileExpectation{mock: mmFindRegisteredProfile.mock}
	}
	mmFindRegisteredProfile.defaultExpectation.results = &OfflinePopulationMockFindRegisteredProfileResults{h1}
	return mmFindRegisteredProfile.mock
}

//Set uses given function f to mock the OfflinePopulation.FindRegisteredProfile method
func (mmFindRegisteredProfile *mOfflinePopulationMockFindRegisteredProfile) Set(f func(identity endpoints.Inbound) (h1 profiles.Host)) *OfflinePopulationMock {
	if mmFindRegisteredProfile.defaultExpectation != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("Default expectation is already set for the OfflinePopulation.FindRegisteredProfile method")
	}

	if len(mmFindRegisteredProfile.expectations) > 0 {
		mmFindRegisteredProfile.mock.t.Fatalf("Some expectations are already set for the OfflinePopulation.FindRegisteredProfile method")
	}

	mmFindRegisteredProfile.mock.funcFindRegisteredProfile = f
	return mmFindRegisteredProfile.mock
}

// When sets expectation for the OfflinePopulation.FindRegisteredProfile which will trigger the result defined by the following
// Then helper
func (mmFindRegisteredProfile *mOfflinePopulationMockFindRegisteredProfile) When(identity endpoints.Inbound) *OfflinePopulationMockFindRegisteredProfileExpectation {
	if mmFindRegisteredProfile.mock.funcFindRegisteredProfile != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("OfflinePopulationMock.FindRegisteredProfile mock is already set by Set")
	}

	expectation := &OfflinePopulationMockFindRegisteredProfileExpectation{
		mock:   mmFindRegisteredProfile.mock,
		params: &OfflinePopulationMockFindRegisteredProfileParams{identity},
	}
	mmFindRegisteredProfile.expectations = append(mmFindRegisteredProfile.expectations, expectation)
	return expectation
}

// Then sets up OfflinePopulation.FindRegisteredProfile return parameters for the expectation previously defined by the When method
func (e *OfflinePopulationMockFindRegisteredProfileExpectation) Then(h1 profiles.Host) *OfflinePopulationMock {
	e.results = &OfflinePopulationMockFindRegisteredProfileResults{h1}
	return e.mock
}

// FindRegisteredProfile implements OfflinePopulation
func (mmFindRegisteredProfile *OfflinePopulationMock) FindRegisteredProfile(identity endpoints.Inbound) (h1 profiles.Host) {
	mm_atomic.AddUint64(&mmFindRegisteredProfile.beforeFindRegisteredProfileCounter, 1)
	defer mm_atomic.AddUint64(&mmFindRegisteredProfile.afterFindRegisteredProfileCounter, 1)

	if mmFindRegisteredProfile.inspectFuncFindRegisteredProfile != nil {
		mmFindRegisteredProfile.inspectFuncFindRegisteredProfile(identity)
	}

	mm_params := &OfflinePopulationMockFindRegisteredProfileParams{identity}

	// Record call args
	mmFindRegisteredProfile.FindRegisteredProfileMock.mutex.Lock()
	mmFindRegisteredProfile.FindRegisteredProfileMock.callArgs = append(mmFindRegisteredProfile.FindRegisteredProfileMock.callArgs, mm_params)
	mmFindRegisteredProfile.FindRegisteredProfileMock.mutex.Unlock()

	for _, e := range mmFindRegisteredProfile.FindRegisteredProfileMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.h1
		}
	}

	if mmFindRegisteredProfile.FindRegisteredProfileMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFindRegisteredProfile.FindRegisteredProfileMock.defaultExpectation.Counter, 1)
		mm_want := mmFindRegisteredProfile.FindRegisteredProfileMock.defaultExpectation.params
		mm_got := OfflinePopulationMockFindRegisteredProfileParams{identity}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFindRegisteredProfile.t.Errorf("OfflinePopulationMock.FindRegisteredProfile got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFindRegisteredProfile.FindRegisteredProfileMock.defaultExpectation.results
		if mm_results == nil {
			mmFindRegisteredProfile.t.Fatal("No results are set for the OfflinePopulationMock.FindRegisteredProfile")
		}
		return (*mm_results).h1
	}
	if mmFindRegisteredProfile.funcFindRegisteredProfile != nil {
		return mmFindRegisteredProfile.funcFindRegisteredProfile(identity)
	}
	mmFindRegisteredProfile.t.Fatalf("Unexpected call to OfflinePopulationMock.FindRegisteredProfile. %v", identity)
	return
}

// FindRegisteredProfileAfterCounter returns a count of finished OfflinePopulationMock.FindRegisteredProfile invocations
func (mmFindRegisteredProfile *OfflinePopulationMock) FindRegisteredProfileAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindRegisteredProfile.afterFindRegisteredProfileCounter)
}

// FindRegisteredProfileBeforeCounter returns a count of OfflinePopulationMock.FindRegisteredProfile invocations
func (mmFindRegisteredProfile *OfflinePopulationMock) FindRegisteredProfileBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindRegisteredProfile.beforeFindRegisteredProfileCounter)
}

// Calls returns a list of arguments used in each call to OfflinePopulationMock.FindRegisteredProfile.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFindRegisteredProfile *mOfflinePopulationMockFindRegisteredProfile) Calls() []*OfflinePopulationMockFindRegisteredProfileParams {
	mmFindRegisteredProfile.mutex.RLock()

	argCopy := make([]*OfflinePopulationMockFindRegisteredProfileParams, len(mmFindRegisteredProfile.callArgs))
	copy(argCopy, mmFindRegisteredProfile.callArgs)

	mmFindRegisteredProfile.mutex.RUnlock()

	return argCopy
}

// MinimockFindRegisteredProfileDone returns true if the count of the FindRegisteredProfile invocations corresponds
// the number of defined expectations
func (m *OfflinePopulationMock) MinimockFindRegisteredProfileDone() bool {
	for _, e := range m.FindRegisteredProfileMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindRegisteredProfileMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFindRegisteredProfileCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindRegisteredProfile != nil && mm_atomic.LoadUint64(&m.afterFindRegisteredProfileCounter) < 1 {
		return false
	}
	return true
}

// MinimockFindRegisteredProfileInspect logs each unmet expectation
func (m *OfflinePopulationMock) MinimockFindRegisteredProfileInspect() {
	for _, e := range m.FindRegisteredProfileMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OfflinePopulationMock.FindRegisteredProfile with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindRegisteredProfileMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFindRegisteredProfileCounter) < 1 {
		if m.FindRegisteredProfileMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OfflinePopulationMock.FindRegisteredProfile")
		} else {
			m.t.Errorf("Expected call to OfflinePopulationMock.FindRegisteredProfile with params: %#v", *m.FindRegisteredProfileMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindRegisteredProfile != nil && mm_atomic.LoadUint64(&m.afterFindRegisteredProfileCounter) < 1 {
		m.t.Error("Expected call to OfflinePopulationMock.FindRegisteredProfile")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *OfflinePopulationMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockFindRegisteredProfileInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *OfflinePopulationMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *OfflinePopulationMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockFindRegisteredProfileDone()
}