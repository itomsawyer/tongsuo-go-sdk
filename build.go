// Copyright (C) 2017. See AUTHORS.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !static
// +build !static

package tongsuogo

// #cgo linux CFLAGS: -Wno-deprecated-declarations -I/opt/tongsuo/include
// #cgo linux LDFLAGS: -L/opt/tongsuo/lib -lssl -lcrypto
// #cgo darwin CFLAGS: -I/opt/tongsuo/include -Wno-deprecated-declarations
// #cgo darwin LDFLAGS: -L/opt/tongsuo/lib -lssl -lcrypto
// #cgo windows CFLAGS: -DWIN32_LEAN_AND_MEAN
// #cgo windows pkg-config: libssl libcrypto
import "C"
