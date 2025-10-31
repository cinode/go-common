/*
Copyright © 2025 Bartłomiej Święcki (byo)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package suite_test

import (
	"testing"

	"github.com/cinode/go-common/picotestify/suite"
)

type sampleSuite struct {
	suite.Suite

	test1T *testing.T
	test2T *testing.T
	test3T *testing.T

	setupSuiteCalls    int
	tearDownSuiteCalls int

	setupTestCalls    int
	tearDownTestCalls int
}

func (s *sampleSuite) SetupSuite()    { s.setupSuiteCalls++ }
func (s *sampleSuite) TearDownSuite() { s.tearDownSuiteCalls++ }
func (s *sampleSuite) SetupTest()     { s.setupTestCalls++ }
func (s *sampleSuite) TearDownTest()  { s.tearDownTestCalls++ }

func (s *sampleSuite) Test1()     { s.test1T = s.T() }
func (s *sampleSuite) NotATest2() { s.test2T = s.T() }
func (s *sampleSuite) Test3()     { s.test3T = s.T() }

func TestSuite(t *testing.T) {
	s := sampleSuite{}

	suite.Run(t, &s)

	if s.test1T == nil {
		t.Errorf("Test1 was not called")
	}

	if s.test2T != nil {
		t.Errorf("NotATest2 was called")
	}

	if s.test3T == nil {
		t.Errorf("Test3 was not called")
	}

	if s.setupSuiteCalls != 1 {
		t.Errorf("SetupSuite was called %d times, expected 1", s.setupSuiteCalls)
	}

	if s.tearDownSuiteCalls != 1 {
		t.Errorf("TearDownSuite was called %d times, expected 1", s.tearDownSuiteCalls)
	}

	if s.setupTestCalls != 2 {
		t.Errorf("SetupTest was called %d times, expected 2", s.setupTestCalls)
	}

	if s.tearDownTestCalls != 2 {
		t.Errorf("TearDownTest was called %d times, expected 2", s.tearDownTestCalls)
	}
}
