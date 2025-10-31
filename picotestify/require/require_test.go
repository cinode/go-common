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

package require_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cinode/go-common/picotestify/require"
)

type testingMock struct {
	helperCalled  bool
	errorCalled   bool
	failNowCalled bool
}

func (t *testingMock) Helper() {
	t.helperCalled = true
}

func (t *testingMock) Error(msgAndArgs ...any) {
	t.errorCalled = true
}

func (t *testingMock) FailNow() {
	t.failNowCalled = true
}

func TestRequire(t *testing.T) {
	var testError = errors.New("test-error")

	for _, tt := range []struct {
		pass func(t require.TestingT)
		fail func(t require.TestingT)
		name string
	}{
		{
			name: "Equal",
			pass: func(t require.TestingT) { require.Equal(t, 1, 1) },
			fail: func(t require.TestingT) { require.Equal(t, 1, 2) },
		},
		{
			name: "NotEqual",
			pass: func(t require.TestingT) { require.NotEqual(t, 1, 2) },
			fail: func(t require.TestingT) { require.NotEqual(t, 1, 1) },
		},
		{
			name: "True",
			pass: func(t require.TestingT) { require.True(t, true) },
			fail: func(t require.TestingT) { require.True(t, false) },
		},
		{
			name: "False",
			pass: func(t require.TestingT) { require.False(t, false) },
			fail: func(t require.TestingT) { require.False(t, true) },
		},
		{
			name: "Nil",
			pass: func(t require.TestingT) { require.Nil(t, nil) },
			fail: func(t require.TestingT) { require.Nil(t, 1) },
		},
		{
			name: "Nil - function",
			pass: func(t require.TestingT) { require.Nil(t, (func())(nil)) },
			fail: func(t require.TestingT) { require.Nil(t, func() {}) },
		},
		{
			name: "NotNil",
			pass: func(t require.TestingT) { require.NotNil(t, 1) },
			fail: func(t require.TestingT) { require.NotNil(t, nil) },
		},
		{
			name: "NoError",
			pass: func(t require.TestingT) { require.NoError(t, nil) },
			fail: func(t require.TestingT) { require.NoError(t, errors.New("error")) },
		},
		{
			name: "ErrorIs",
			pass: func(t require.TestingT) { require.ErrorIs(t, fmt.Errorf("error: %w", testError), testError) },
			fail: func(t require.TestingT) { require.ErrorIs(t, fmt.Errorf("error: %w", testError), errors.New("error2")) },
		},
		{
			name: "ErrorContains",
			pass: func(t require.TestingT) { require.ErrorContains(t, errors.New("test error: check"), "error") },
			fail: func(t require.TestingT) { require.ErrorContains(t, errors.New("test error: check"), "error2") },
		},
		{
			name: "Empty",
			pass: func(t require.TestingT) { require.Empty(t, []int{}) },
			fail: func(t require.TestingT) { require.Empty(t, []int{1}) },
		},
		{
			name: "NotEmpty",
			pass: func(t require.TestingT) { require.NotEmpty(t, 1) },
			fail: func(t require.TestingT) { require.NotEmpty(t, nil) },
		},
		{
			name: "Zero",
			pass: func(t require.TestingT) { require.Zero(t, 0) },
			fail: func(t require.TestingT) { require.Zero(t, 1) },
		},
		{
			name: "NotZero",
			pass: func(t require.TestingT) { require.NotZero(t, 1) },
			fail: func(t require.TestingT) { require.NotZero(t, 0) },
		},
		{
			name: "Greater",
			pass: func(t require.TestingT) { require.Greater(t, 2, 1) },
			fail: func(t require.TestingT) { require.Greater(t, 1, 2) },
		},
		{
			name: "GreaterOrEqual",
			pass: func(t require.TestingT) { require.GreaterOrEqual(t, 2, 2) },
			fail: func(t require.TestingT) { require.GreaterOrEqual(t, 1, 2) },
		},
		{
			name: "Panics",
			pass: func(t require.TestingT) { require.Panics(t, func() { panic("test") }) },
			fail: func(t require.TestingT) { require.Panics(t, func() {}) },
		},
		{
			name: "NotPanics",
			pass: func(t require.TestingT) { require.NotPanics(t, func() {}) },
			fail: func(t require.TestingT) { require.NotPanics(t, func() { panic("test") }) },
		},
		{
			name: "Regexp",
			pass: func(t require.TestingT) { require.Regexp(t, `^t.s*t+$`, "test") },
			fail: func(t require.TestingT) { require.Regexp(t, `^t.s*t+$`, "a test") },
		},
		{
			name: "Regexp - bad pattern",
			pass: func(t require.TestingT) { require.Regexp(t, `test`, "test") },
			fail: func(t require.TestingT) { require.Regexp(t, `*test`, "test") },
		},
		{
			name: "Len",
			pass: func(t require.TestingT) { require.Len(t, []int{1, 2, 3}, 3) },
			fail: func(t require.TestingT) { require.Len(t, []int{1, 2, 3}, 2) },
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("pass", func(t *testing.T) {
				hlp := &testingMock{}
				tt.pass(hlp)
				require.True(t, hlp.helperCalled)
				require.False(t, hlp.errorCalled)
				require.False(t, hlp.failNowCalled)
			})
			t.Run("fail", func(t *testing.T) {
				hlp := &testingMock{}
				tt.fail(hlp)
				require.True(t, hlp.helperCalled)
				require.True(t, hlp.errorCalled)
				require.True(t, hlp.failNowCalled)
			})
		})
	}
}
