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

package cutl_test

import (
	"errors"
	"testing"

	"github.com/cinode/go-common/cutl"
	"github.com/cinode/go-common/picotestify/require"
)

func TestPanicIf(t *testing.T) {
	require.NotPanics(t, func() {
		cutl.PanicIf(false, "no panic")
	})

	require.Panics(t, func() {
		cutl.PanicIf(true, "panic")
	})
}

func TestPanicIfError(t *testing.T) {
	require.NotPanics(t, func() {
		cutl.PanicIfError(nil)
	})

	require.Panics(t, func() {
		cutl.PanicIfError(errors.New("error"))
	})
}

func TestMust(t *testing.T) {
	require.NotPanics(t, func() {
		require.Equal(t, 1234, cutl.Must(1234, nil))
	})

	require.Panics(t, func() {
		require.Equal(t, 1234, cutl.Must(1234, errors.New("error")))
	})
}

func TestMust2(t *testing.T) {
	require.NotPanics(t, func() {
		v1, v2 := cutl.Must2(1234, 5678, nil)

		require.Equal(t, 1234, v1)
		require.Equal(t, 5678, v2)
	})

	require.Panics(t, func() {
		v1, v2 := cutl.Must2(1234, 5678, errors.New("error"))
		require.Equal(t, 1234, v1)
		require.Equal(t, 5678, v2)
	})
}
