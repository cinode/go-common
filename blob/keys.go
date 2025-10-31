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
	"bytes"
	"crypto/subtle"
)

// Key with cipher type
type Key struct{ key []byte }

func KeyFromBytes(key []byte) *Key { return &Key{key: bytes.Clone(key)} }
func (k *Key) Bytes() []byte       { return bytes.Clone(k.key) }
func (k *Key) Equal(k2 *Key) bool  { return subtle.ConstantTimeCompare(k.key, k2.key) == 1 }

// IV
type IV struct{ iv []byte }

func IVFromBytes(iv []byte) *IV { return &IV{iv: bytes.Clone(iv)} }
func (i *IV) Bytes() []byte     { return bytes.Clone(i.iv) }
func (i *IV) Equal(i2 *IV) bool { return subtle.ConstantTimeCompare(i.iv, i2.iv) == 1 }
