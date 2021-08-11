package compile

import (
	"matrixone/pkg/sql/protocol"
)

func Transfer(s *Scope) protocol.Scope {
	var ps protocol.Scope

	ps.Ins = s.Ins
	ps.Magic = s.Magic
	if s.Data != nil {
		ps.Data.ID = s.Data.ID
		ps.Data.DB = s.Data.DB
		ps.Data.Refer = s.Data.Refs
		ps.Data.Segs = make([]protocol.Segment, len(s.Data.Segs))
		for i, seg := range s.Data.Segs {
			ps.Data.Segs[i].Id = seg.Id
			ps.Data.Segs[i].GroupId = seg.GroupId
			ps.Data.Segs[i].Version = seg.Version
			ps.Data.Segs[i].IsRemote = seg.IsRemote
			ps.Data.Segs[i].TabletId = seg.TabletId
		}
	}
	ps.Ss = make([]protocol.Scope, len(s.Ss))
	for i := range s.Ss {
		ps.Ss[i] = Transfer(s.Ss[i])
	}
	return ps
}