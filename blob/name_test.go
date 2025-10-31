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

package blob

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/cinode/go-common/picotestify/assert"
	"github.com/cinode/go-common/picotestify/require"
)

func TestBlobName(t *testing.T) {
	for _, h := range [][]byte{
		{0}, {1}, {2}, {0xFE}, {0xFF},
		{0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6},
		sha256.New().Sum(nil),
	} {
		for _, bt := range []Type{
			{t: 0x00},
			{t: 0x01},
			{t: 0x02},
			{t: 0xFE},
			{t: 0xFF},
		} {
			t.Run(fmt.Sprintf("%v:%v", bt, h), func(t *testing.T) {
				bn, err := NameFromHashAndType(h, bt)
				assert.NoError(t, err)
				assert.NotEmpty(t, bn)
				assert.Greater(t, len(bn.bn), len(h))
				assert.Equal(t, h, bn.Hash())
				assert.Equal(t, bt, bn.Type())

				s := bn.String()
				bn2, err := NameFromString(s)
				require.NoError(t, err)
				require.Equal(t, bn, bn2)
				require.True(t, bn.Equal(bn2))

				b := bn.Bytes()
				bn3, err := NameFromBytes(b)
				require.NoError(t, err)
				require.Equal(t, bn, bn3)
				require.True(t, bn.Equal(bn3))
			})
		}
	}

	_, err := NameFromString("!@#")
	require.ErrorIs(t, err, ErrInvalidBlobName)

	_, err = NameFromString("")
	require.ErrorIs(t, err, ErrInvalidBlobName)

	_, err = NameFromHashAndType(nil, Type{t: 0x00})
	require.ErrorIs(t, err, ErrInvalidBlobName)
}
