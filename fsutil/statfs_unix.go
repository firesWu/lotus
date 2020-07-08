package fsutil

import (
	"syscall"

	"golang.org/x/xerrors"
)

func Statfs(path string) (FsStat, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return FsStat{}, xerrors.Errorf("statfs: %w", err)
	}

	return FsStat{
		Capacity:  int64(stat.Blocks) * stat.Bsize,
		Available: int64(stat.Bavail) * stat.Bsize,
	}, nil
}
