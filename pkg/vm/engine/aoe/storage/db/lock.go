package db

import (
	log "github.com/sirupsen/logrus"
	"io"
	e "matrixone/pkg/vm/engine/aoe/storage"
	"os"
	"syscall"
)

const (
	LockName string = "AOE"
)

func createDBLock(dir string) (io.Closer, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
	}
	fname := e.MakeFilename(dir, e.FTLock, LockName, false)
	f, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	flockT := syscall.Flock_t{
		Type:   syscall.F_WRLCK,
		Whence: io.SeekStart,
		Start:  0,
		Len:    0,
		Pid:    int32(os.Getpid()),
	}
	if err := syscall.FcntlFlock(f.Fd(), syscall.F_SETLK, &flockT); err != nil {
		log.Errorf("error locking file: %s", err)
		f.Close()
		return nil, err
	}
	return f, nil
}