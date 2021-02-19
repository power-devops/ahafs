// Copyright 2021 Power-Devops.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ahafs

import (
	"fmt"
	"time"
)

func (e Event) String() string {
	return fmt.Sprintf(`Timestamp (seconds): %d
Timestamp (nano-seconds): %d
Timestamp (human readable): %s
Sequence number: %d
Process ID: %d
User ID: %d
Login User ID: %d
Group ID: %d
Program name: %s
Current value: %d
Return code from event producer: %d
Information from event producer: %s
`, e.TimeSec, e.TimeNSec, time.Unix(int64(e.TimeSec), int64(e.TimeNSec)), e.SeqNumber, e.PID, e.UID, e.RUID, e.GID, e.Program, e.Value, e.RC, e.Info)
}

// Time returns time of the event a Go structure time.Time
func (e Event) Time() time.Time {
	return time.Unix(int64(e.TimeSec), int64(e.TimeNSec))
}
