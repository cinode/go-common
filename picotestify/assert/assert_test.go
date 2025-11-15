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

package assert_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cinode/go-common/picotestify/assert"
)

type testingMock struct {
	helper bool
	error  bool
}

func (t *testingMock) Helper() {
	t.helper = true
}

func (t *testingMock) Error(msgAndArgs ...any) {
	t.error = true
}

func TestAssert(t *testing.T) {
	var testError = errors.New("test-error")

	for _, tt := range []struct {
		pass func(t assert.TestingT)
		fail func(t assert.TestingT)
		name string
	}{
		{
			name: "Equal",
			pass: func(t assert.TestingT) { assert.Equal(t, 1, 1) },
			fail: func(t assert.TestingT) { assert.Equal(t, 1, 2) },
		},
		{
			name: "NotEqual",
			pass: func(t assert.TestingT) { assert.NotEqual(t, 1, 2) },
			fail: func(t assert.TestingT) { assert.NotEqual(t, 1, 1) },
		},
		{
			name: "True",
			pass: func(t assert.TestingT) { assert.True(t, true) },
			fail: func(t assert.TestingT) { assert.True(t, false) },
		},
		{
			name: "False",
			pass: func(t assert.TestingT) { assert.False(t, false) },
			fail: func(t assert.TestingT) { assert.False(t, true) },
		},
		{
			name: "Nil",
			pass: func(t assert.TestingT) { assert.Nil(t, nil) },
			fail: func(t assert.TestingT) { assert.Nil(t, 1) },
		},
		{
			name: "Nil - function",
			pass: func(t assert.TestingT) { assert.Nil(t, (func())(nil)) },
			fail: func(t assert.TestingT) { assert.Nil(t, func() {}) },
		},
		{
			name: "NotNil",
			pass: func(t assert.TestingT) { assert.NotNil(t, 1) },
			fail: func(t assert.TestingT) { assert.NotNil(t, nil) },
		},
		{
			name: "NoError",
			pass: func(t assert.TestingT) { assert.NoError(t, nil) },
			fail: func(t assert.TestingT) { assert.NoError(t, errors.New("error")) },
		},
		{
			name: "ErrorIs",
			pass: func(t assert.TestingT) { assert.ErrorIs(t, fmt.Errorf("error: %w", testError), testError) },
			fail: func(t assert.TestingT) { assert.ErrorIs(t, fmt.Errorf("error: %w", testError), errors.New("error2")) },
		},
		{
			name: "ErrorContains",
			pass: func(t assert.TestingT) { assert.ErrorContains(t, errors.New("test error: check"), "error") },
			fail: func(t assert.TestingT) { assert.ErrorContains(t, errors.New("test error: check"), "error2") },
		},
		{
			name: "ErrorContains - nil",
			pass: func(t assert.TestingT) { assert.ErrorContains(t, errors.New("test error: check"), "error") },
			fail: func(t assert.TestingT) { assert.ErrorContains(t, nil, "error") },
		},
		{
			name: "Empty",
			pass: func(t assert.TestingT) { assert.Empty(t, []int{}) },
			fail: func(t assert.TestingT) { assert.Empty(t, []int{1}) },
		},
		{
			name: "NotEmpty",
			pass: func(t assert.TestingT) { assert.NotEmpty(t, 1) },
			fail: func(t assert.TestingT) { assert.NotEmpty(t, nil) },
		},
		{
			name: "Zero",
			pass: func(t assert.TestingT) { assert.Zero(t, 0) },
			fail: func(t assert.TestingT) { assert.Zero(t, 1) },
		},
		{
			name: "NotZero",
			pass: func(t assert.TestingT) { assert.NotZero(t, 1) },
			fail: func(t assert.TestingT) { assert.NotZero(t, 0) },
		},
		{
			name: "Greater",
			pass: func(t assert.TestingT) { assert.Greater(t, 2, 1) },
			fail: func(t assert.TestingT) { assert.Greater(t, 1, 2) },
		},
		{
			name: "GreaterOrEqual",
			pass: func(t assert.TestingT) { assert.GreaterOrEqual(t, 2, 2) },
			fail: func(t assert.TestingT) { assert.GreaterOrEqual(t, 1, 2) },
		},
		{
			name: "Panics",
			pass: func(t assert.TestingT) { assert.Panics(t, func() { panic("test") }) },
			fail: func(t assert.TestingT) { assert.Panics(t, func() {}) },
		},
		{
			name: "NotPanics",
			pass: func(t assert.TestingT) { assert.NotPanics(t, func() {}) },
			fail: func(t assert.TestingT) { assert.NotPanics(t, func() { panic("test") }) },
		},
		{
			name: "Regexp",
			pass: func(t assert.TestingT) { assert.Regexp(t, `^t.s*t+$`, "test") },
			fail: func(t assert.TestingT) { assert.Regexp(t, `^t.s*t+$`, "a test") },
		},
		{
			name: "Regexp - bad pattern",
			pass: func(t assert.TestingT) { assert.Regexp(t, `test`, "test") },
			fail: func(t assert.TestingT) { assert.Regexp(t, `*test`, "test") },
		},
		{
			name: "Len",
			pass: func(t assert.TestingT) { assert.Len(t, []int{1, 2, 3}, 3) },
			fail: func(t assert.TestingT) { assert.Len(t, []int{1, 2, 3}, 2) },
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("pass", func(t *testing.T) {
				hlp := &testingMock{}
				tt.pass(hlp)
				assert.True(t, hlp.helper)
				assert.False(t, hlp.error)
			})
			t.Run("fail", func(t *testing.T) {
				hlp := &testingMock{}
				tt.fail(hlp)
				assert.True(t, hlp.helper)
				assert.True(t, hlp.error)
			})
		})
	}
}
