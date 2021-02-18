// Copyright 2021 Power-Devops.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ahafs

// Monitor describes a single monitor in AHAFS
type Monitor struct {
	Path string // path to monitor
	Fd   int    // file descriptor
}

// Event is a generalized AHAFS event
type Event struct {
	TimeSec   int
	TimeNSec  int
	SeqNumber int
	PID       int
	UID       int
	RUID      int
	GID       int
	Program   string
	RC        int
	Value     int
	Info      string
}
