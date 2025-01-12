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

package compile

import (
	"unsafe"

	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/hashtable"
	"github.com/matrixorigin/matrixone/pkg/container/types"
)

const (
	UnitLimit = 256
)

func constructViews(bats []*batch.Batch, fvars []string) {
	for i, fvar := range fvars {
		constructView(bats[i], fvar)
	}
}

func constructView(bat *batch.Batch, fvar string) {
	ht := &hashtable.Int64HashMap{}
	ht.Init()
	hashes := make([]uint64, UnitLimit)
	values := make([]uint64, UnitLimit)
	keys := make([]uint64, UnitLimit)
	vec := batch.GetVector(bat, fvar)
	switch vec.Typ.Oid {
	case types.T_int8:
		vs := vec.Col.([]int8)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_int16:
		vs := vec.Col.([]int16)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_int32:
		vs := vec.Col.([]int32)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_date:
		vs := vec.Col.([]types.Date)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_int64:
		vs := vec.Col.([]int64)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_datetime:
		vs := vec.Col.([]types.Datetime)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_uint8:
		vs := vec.Col.([]uint8)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_uint16:
		vs := vec.Col.([]uint16)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_uint32:
		vs := vec.Col.([]uint32)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_uint64:
		vs := vec.Col.([]uint64)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_float32:
		vs := vec.Col.([]float32)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_float64:
		vs := vec.Col.([]float64)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			{
				for k := 0; k < n; k++ {
					keys[k] = uint64(vs[int(i)+k])
				}
			}
			hashes[0] = 0
			ht.InsertBatch(n, hashes, unsafe.Pointer(&keys[0]), values)
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	case types.T_char, types.T_varchar:
		ht := &hashtable.StringHashMap{}
		ht.Init()
		var strKeys [UnitLimit][]byte
		var strKeys16 [UnitLimit][16]byte
		var zStrKeys16 [UnitLimit][16]byte
		var states [UnitLimit][3]uint64
		vs := vec.Col.(*types.Bytes)
		count := int64(len(bat.Zs))
		for i := int64(0); i < count; i += UnitLimit {
			n := int(count - i)
			if n > UnitLimit {
				n = UnitLimit
			}
			var padded int
			{
				for k := 0; k < n; k++ {
					if vs.Lengths[i+int64(k)] < 16 {
						copy(strKeys16[padded][:], vs.Get(i+int64(k)))
						strKeys[k] = strKeys16[padded][:]
						padded++
					} else {
						strKeys[k] = vs.Get(i + int64(k))
					}
				}
			}
			ht.InsertStringBatch(states[:], strKeys[:n], values)
			copy(strKeys16[:padded], zStrKeys16[:padded])
		}
		if len(bat.Zs) == int(ht.Cardinality()) {
			bat.Ht = ht
			return
		}
	}
	bat.Ht = nil
}
