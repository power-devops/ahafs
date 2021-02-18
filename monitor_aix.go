// Copyright 2021 Power-Devops.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ahafs

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

func NewMonitor(name, params string) (*Monitor, error) {
	if !is_aha_mounted() {
		return nil, fmt.Errorf("AHAFS is not mounted")
	}
	if params == "" {
		params = "CHANGED=YES;WAIT_TYPE=WAIT_IN_SELECT;INFO_LVL=2"
	}
	name = filepath.Clean(name)
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return nil, err
	}
	fd, err := aha_creat(name, 0666)
	if err != nil {
		return nil, err
	}
	_, err = unix.Write(fd, []byte(params))
	if err != nil {
		return nil, err
	}
	m := &Monitor{
		Path: name,
		Fd:   fd,
	}
	return m, nil
}

func (m *Monitor) Watch(c chan Event) error {
	readfds := &unix.FdSet{}
	readfds.Set(m.Fd)
	nfd := m.Fd + 1
	for {
		n, err := unix.Select(nfd, readfds, nil, nil, nil)
		if err != nil {
			return err
		}
		if n > 0 {
			buf := make([]byte, 4096)
			_, err = unix.Pread(m.Fd, buf, 0)
			if err != nil {
				return err
			}
			c <- buf2evt(buf)
		}
	}
}
