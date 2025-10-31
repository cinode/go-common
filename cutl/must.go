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

package cutl

func PanicIf[T any](condition bool, value T) {
	if condition {
		panic(value)
	}
}

func PanicIfError(err error) {
	PanicIf(err != nil, err)
}

func Must[T any](val T, err error) T {
	PanicIfError(err)
	return val
}

func Must2[T1 any, T2 any](val1 T1, val2 T2, err error) (ret1 T1, ret2 T2) {
	PanicIfError(err)
	return val1, val2
}
