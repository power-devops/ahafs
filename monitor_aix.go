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

func NewFileMonitor(name string) (*Monitor, error) {
	if _, err := os.Stat(name); err != nil {
		return nil, err
	}
	name = filepath.Join("/aha/fs/modFile.monFactory", filepath.Clean(name)+".mon")
	return NewMonitor(name, "")
}

func NewFileAttrMonitor(name string) (*Monitor, error) {
	if _, err := os.Stat(name); err != nil {
		return nil, err
	}
	name = filepath.Join("/aha/fs/modFileAttr.monFactory", filepath.Clean(name)+".mon")
	return NewMonitor(name, "")
}

func NewDirMonitor(name string) (*Monitor, error) {
	if fi, err := os.Stat(name); err != nil {
		return nil, err
	} else if !fi.IsDir() {
		return nil, fmt.Errorf("not a directory")
	}
	name = filepath.Join("/aha/fs/modDir.monFactory", filepath.Clean(name)+".mon")
	return NewMonitor(name, "")
}

func NewMonitor(name, params string) (*Monitor, error) {
	if !isAhaMounted() {
		return nil, fmt.Errorf("AHAFS is not mounted")
	}
	if params == "" {
		params = "CHANGED=YES;WAIT_TYPE=WAIT_IN_SELECT;INFO_LVL=2"
	}
	name = filepath.Clean(name)
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return nil, err
	}
	fd, err := ahaCreat(name, 0666)
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

func (m *Monitor) Close() error {
	return unix.Close(m.Fd)
}

func safesend(c chan Event, value Event) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	c <- value
	return false
}

func (m *Monitor) Watch(c chan Event) error {
	readfds := &unix.FdSet{}
	readfds.Set(m.Fd)
	nfd := m.Fd + 1
	for {
		n, err := unix.Select(nfd, readfds, nil, nil, nil)
		if err != nil {
            closed := safesend(c, Event{Quit: true, Error: err})
            if !closed {
			    close(c)
            }
			return err
		}
		if n > 0 {
			buf := make([]byte, 4096)
			_, err = unix.Pread(m.Fd, buf, 0)
			if err != nil {
                closed := safesend(c, Event{Quit: true, Error: err})
                if !closed {
				    close(c)
                }
				return err
			}
			safesend(c, buf2evt(buf))
		}
	}
}
