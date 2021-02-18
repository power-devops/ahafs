// Copyright 2021 Power-Devops.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build !aix

package ahafs

import "fmt"

// NewMonitor creates a new monitor in AHAFS filesystem.
// name is the name of the monitor, for example /aha/fs/modDir.monFactory/tmp.mon
// params are the monitoring parameters. Default is CHANGED=YES;WAIT_TYPE=WAIT_IN_SELECT;INFO_LVL=2
func NewMonitor(name, params string) (*Monitor, error) {
	return nil, fmt.Errorf("not AIX")
}

// Watch monitors for new events in the defined monitor and send events into the channel
func (m *Monitor) Watch(c chan Event) error {
	return fmt.Errorf("not AIX")
}
