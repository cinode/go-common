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
	"errors"

	"github.com/cinode/go-common/base58"
)

var (
	ErrInvalidBlobName = errors.New("invalid blob name")
)

// Name is used to identify blobs.
// Internally it is a single array of bytes that represents
// both the type of the blob and internal hash used to create that blob.
// The type of the blob is not stored directly. Instead it is mixed
// with the hash of the blob to make sure that all bytes in the blob name
// are randomly distributed.
type Name struct {
	bn []byte
}

// NameFromHashAndType generates the name of a blob from some hash (e.g. sha256 of blob's content)
// and given blob type
func NameFromHashAndType(hash []byte, t Type) (*Name, error) {
	if len(hash) == 0 || len(hash) > 0x7E {
		return nil, ErrInvalidBlobName
	}

	bn := make([]byte, len(hash)+1)

	copy(bn[1:], hash)

	scrambledTypeByte := t.t
	for _, b := range hash {
		scrambledTypeByte ^= b
	}
	bn[0] = scrambledTypeByte

	return &Name{bn: bn}, nil
}

// NameFromString decodes base58-encoded string into blob name
func NameFromString(s string) (*Name, error) {
	decoded, err := base58.Decode(s)
	if err != nil {
		return nil, ErrInvalidBlobName
	}
	return NameFromBytes(decoded)
}

func NameFromBytes(n []byte) (*Name, error) {
	if len(n) == 0 || len(n) > 0x7F {
		return nil, ErrInvalidBlobName
	}
	return &Name{bn: bytes.Clone(n)}, nil
}

// Returns base58-encoded blob name
func (b *Name) String() string {
	return base58.Encode(b.bn)
}

// Extracts hash from blob name
func (b *Name) Hash() []byte {
	return b.bn[1:]
}

// Extracts blob type from the name
func (b *Name) Type() Type {
	ret := byte(0)
	for _, by := range b.bn {
		ret ^= by
	}
	return Type{t: ret}
}

func (b *Name) Bytes() []byte {
	return bytes.Clone(b.bn)
}

func (b *Name) Equal(b2 *Name) bool {
	return subtle.ConstantTimeCompare(b.bn, b2.bn) == 1
}
