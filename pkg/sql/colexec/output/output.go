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

package output

import (
	"bytes"
	"matrixone/pkg/container/batch"
	"matrixone/pkg/vm/process"
)

func String(arg interface{}, buf *bytes.Buffer) {
	buf.WriteString("sql output")
}

func Prepare(_ *process.Process, _ interface{}) error {
	return nil
}

func Call(proc *process.Process, arg interface{}) (bool, error) {
	ap := arg.(*Argument)
	if bat := proc.Reg.InputBatch; bat != nil && len(bat.Zs) > 0 {
		if len(ap.Attrs) > 0 {
			batch.Reorder(bat, ap.Attrs)
		}
		if err := ap.Func(ap.Data, bat); err != nil {
			batch.Clean(bat, proc.Mp)
			return true, err
		}
		batch.Clean(bat, proc.Mp)
	}
	return false, nil
}
