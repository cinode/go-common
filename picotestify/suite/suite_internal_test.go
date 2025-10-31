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

package suite

import (
	"reflect"
	"testing"
)

type sampleSuite struct {
	Suite
}

func (s *sampleSuite) Test1()     {}
func (s *sampleSuite) NotATest2() {}
func (s *sampleSuite) Test3()     {}

func TestListTests(t *testing.T) {
	t.Run("list all tests", func(t *testing.T) {
		tests := []string{}
		for name, testFunc := range listTests(&sampleSuite{}) {
			if testFunc == nil {
				t.Errorf("test function is nil for test %s", name)
				continue
			}
			tests = append(tests, name)
		}

		if !reflect.DeepEqual(tests, []string{"Test1", "Test3"}) {
			t.Errorf("expected tests to be [Test1, Test3], got %v", tests)
		}
	})

	t.Run("get one test only", func(t *testing.T) {
		for name := range listTests(&sampleSuite{}) {
			if name != "Test1" {
				t.Errorf("expected Test1, got %s", name)
				continue
			}

			break
		}
	})
}
