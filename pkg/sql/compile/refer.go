package compile

import (
	"matrixone/pkg/sql/colexec/extend"
	"matrixone/pkg/sql/op"
	"matrixone/pkg/sql/op/innerJoin"
	"matrixone/pkg/sql/op/relation"
)

func IncRef(e extend.Extend, mp map[string]uint64) {
	switch v := e.(type) {
	case *extend.Attribute:
		mp[v.Name]++
	case *extend.UnaryExtend:
		IncRef(v.E, mp)
	case *extend.BinaryExtend:
		IncRef(v.Left, mp)
		IncRef(v.Right, mp)
	}
}

func IsSource(o op.OP) bool {
	switch o.(type) {
	case *relation.Relation:
		return true
	case *innerJoin.Join:
		return true
	}
	return false
}