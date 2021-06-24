package dataio

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"matrixone/pkg/vm/engine/aoe/storage/common"
	dio "matrixone/pkg/vm/engine/aoe/storage/dataio"
	"matrixone/pkg/vm/engine/aoe/storage/layout/base"
	"sync"
)

var (
	DupSegError = errors.New("duplicate seg")
)

type FileType uint8

const (
	UnsortedSegFile FileType = iota
	SortedSegFile
)

var DefaultFsMgr base.IManager
var MockFsMgr base.IManager

func init() {
	DefaultFsMgr = NewManager(dio.WRITER_FACTORY.Dirname, false)
	MockFsMgr = NewManager(dio.WRITER_FACTORY.Dirname, true)
}

type Manager struct {
	sync.RWMutex
	UnsortedFiles map[common.ID]base.ISegmentFile
	SortedFiles   map[common.ID]base.ISegmentFile
	Dir           string
	Mock          bool
}

func NewManager(dir string, mock bool) *Manager {
	return &Manager{
		UnsortedFiles: make(map[common.ID]base.ISegmentFile),
		SortedFiles:   make(map[common.ID]base.ISegmentFile),
		Dir:           dir,
		Mock:          mock,
	}
}

func (mgr *Manager) RegisterUnsortedFiles(id common.ID) (base.ISegmentFile, error) {
	var usf base.ISegmentFile
	if mgr.Mock {
		usf = NewMockSegmentFile(mgr.Dir, UnsortedSegFile, id)
	} else {
		usf = NewUnsortedSegmentFile(mgr.Dir, id)
	}
	mgr.Lock()
	defer mgr.Unlock()
	_, ok := mgr.UnsortedFiles[id]
	if ok {
		usf.Close()
		return nil, DupSegError
	}
	mgr.UnsortedFiles[id] = usf
	return usf, nil
}

func (mgr *Manager) RegisterSortedFiles(id common.ID) (base.ISegmentFile, error) {
	var sf base.ISegmentFile
	if mgr.Mock {
		sf = NewMockSegmentFile(mgr.Dir, UnsortedSegFile, id)
	} else {
		sf = NewSortedSegmentFile(mgr.Dir, id)
	}
	mgr.Lock()
	defer mgr.Unlock()
	_, ok := mgr.UnsortedFiles[id]
	if ok {
		sf.Close()
		return nil, DupSegError
	}
	_, ok = mgr.SortedFiles[id]
	if ok {
		sf.Close()
		return nil, DupSegError
	}
	mgr.SortedFiles[id] = sf
	return sf, nil
}

func (mgr *Manager) UpgradeFile(id common.ID) base.ISegmentFile {
	var sf base.ISegmentFile
	if mgr.Mock {
		sf = NewMockSegmentFile(mgr.Dir, UnsortedSegFile, id)
	} else {
		sf = NewSortedSegmentFile(mgr.Dir, id)
	}
	mgr.Lock()
	staleFile, ok := mgr.UnsortedFiles[id]
	if !ok {
		log.Info(mgr.stringNoLock())
		panic(fmt.Sprintf("upgrade file %s not found", id.SegmentString()))
	}
	defer staleFile.Close()
	delete(mgr.UnsortedFiles, id)
	_, ok = mgr.SortedFiles[id]
	if ok {
		panic(fmt.Sprintf("duplicate file %s", id.SegmentString()))
	}
	mgr.SortedFiles[id] = sf
	mgr.Unlock()
	return sf
}

func (mgr *Manager) GetUnsortedFile(id common.ID) base.ISegmentFile {
	mgr.RLock()
	defer mgr.RUnlock()
	f, ok := mgr.UnsortedFiles[id]
	if !ok {
		return nil
	}
	return f
}

func (mgr *Manager) GetSortedFile(id common.ID) base.ISegmentFile {
	mgr.RLock()
	defer mgr.RUnlock()
	f, ok := mgr.SortedFiles[id]
	if !ok {
		return nil
	}
	return f
}

func (mgr *Manager) String() string {
	mgr.RLock()
	defer mgr.RUnlock()
	return mgr.stringNoLock()
}

func (mgr *Manager) stringNoLock() string {
	s := fmt.Sprintf("<Manager:%s>[%v]: Unsorted[%d], Sorted[%d]\n", mgr.Dir, mgr.Mock, len(mgr.UnsortedFiles), len(mgr.SortedFiles))
	if len(mgr.UnsortedFiles) > 0 {
		for k, _ := range mgr.UnsortedFiles {
			s = fmt.Sprintf("%s %s", s, k.SegmentString())
		}
		s = fmt.Sprintf("%s\n", s)
	}
	if len(mgr.SortedFiles) > 0 {
		for k, _ := range mgr.SortedFiles {
			s = fmt.Sprintf("%s %s", s, k.SegmentString())
		}
	}
	return s
}

func (mgr *Manager) Close() error {
	var err error
	for _, f := range mgr.UnsortedFiles {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}
	for _, f := range mgr.SortedFiles {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}
	return nil
}