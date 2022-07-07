// Copyright 2021 Power-Devops.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build !aix
// +build !aix

package ahafs

import "fmt"

// NewMonitor creates a new monitor in AHAFS filesystem.
// name is the name of the monitor, for example /aha/fs/modDir.monFactory/tmp.mon
// params are the monitoring parameters. Default is CHANGED=YES;WAIT_TYPE=WAIT_IN_SELECT;INFO_LVL=2
func NewMonitor(name, params string) (*Monitor, error) {
	return nil, fmt.Errorf("not AIX")
}

// NewFileMonitor creates a new file monitor in AHAFS filesystem.
// File monitor can be created only if the file already exists.
func NewFileMonitor(name string) (*Monitor, error) {
	return nil, fmt.Errorf("not AIX")
}

// NewFileAttrMonitor creates a new file attribute monitor in AHAFS filesystem.
// File attribute monitor can be created only if the file already exists.
func NewFileAttrMonitor(name string) (*Monitor, error) {
	return nil, fmt.Errorf("not AIX")
}

// NewDirMonitor creates a new directory monitor in AHAFS filesystem.
// Directory monitor can be created only if the directory already exists.
func NewDirMonitor(name string) (*Monitor, error) {
	return nil, fmt.Errorf("not AIX")
}

// Close closes the file descriptor and ends monitoring of the object.
func (m *Monitor) Close() error {
	return fmt.Errorf("not AIX")
}

// Watch monitors for new events in the defined monitor and send events into the channel
func (m *Monitor) Watch(c chan Event) error {
	return fmt.Errorf("not AIX")
}
