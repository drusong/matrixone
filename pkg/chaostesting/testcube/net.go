// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"strconv"
	"sync/atomic"
)

var nextPort int64

func randPort() string {
	return strconv.FormatInt(
		10000+atomic.AddInt64(&nextPort, 1)%50000,
		10,
	)
}

type ListenHost string

func (_ Def) ListenHost() ListenHost {
	return "localhost"
}
