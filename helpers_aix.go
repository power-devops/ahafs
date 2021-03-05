// Copyright 2021 Power-Devops.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ahafs

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func ahaCreat(path string, mode uint32) (fd int, err error) {
	return unix.Open(path, unix.O_CREAT|unix.O_RDWR|unix.O_TRUNC, mode)
}

func isAhaMounted() bool {
	st := unix.Stat_t{}
	err := unix.Stat("/aha", &st)
	if err != nil {
		return false
	}
	if st.Flag != 1 {
		return false
	}
	return true
}

func getvalue(s string) string {
	re := regexp.MustCompile("(.*)=(.*)")
	m := re.FindStringSubmatch(s)
	if len(m) == 3 {
		return m[2]
	}
	return ""
}

func buf2evt(buf []byte) Event {
	e := Event{Quit: false}
	info := false
	scanner := bufio.NewScanner(bytes.NewReader(buf))
	for scanner.Scan() {
		s := scanner.Text()
		if info && s != "END_EVPROD_INFO" {
			if e.Info == "" {
				e.Info = s
			} else {
				e.Info = e.Info + "\n" + s
			}
			continue
		}
		switch {
		case s == "BEGIN_EVENT_INFO":
			// do nothing, skip it
		case s == "END_EVENT_INFO":
			// last line of the event
			break
		case strings.HasPrefix(s, "TIME_tvsec="):
			e.TimeSec, _ = strconv.Atoi(getvalue(s))
		case strings.HasPrefix(s, "TIME_tvnsec="):
			e.TimeNSec, _ = strconv.Atoi(getvalue(s))
		case strings.HasPrefix(s, "SEQUENCE_NUM="):
			e.SeqNumber, _ = strconv.Atoi(getvalue(s))
		case strings.HasPrefix(s, "PID="):
			e.PID, _ = strconv.Atoi(getvalue(s))
		case strings.HasPrefix(s, "UID="):
			e.UID, _ = strconv.Atoi(getvalue(s))
		case strings.HasPrefix(s, "UID_LOGIN="):
			e.RUID, _ = strconv.Atoi(getvalue(s))
		case strings.HasPrefix(s, "GID="):
			e.GID, _ = strconv.Atoi(getvalue(s))
		case strings.HasPrefix(s, "PROG_NAME="):
			e.Program = getvalue(s)
		case strings.HasPrefix(s, "RC_FROM_EVPROD="):
			e.RC, _ = strconv.Atoi(getvalue(s))
		case strings.HasPrefix(s, "CURRENT_VALUE="):
			e.Value, _ = strconv.Atoi(getvalue(s))
		case s == "BEGIN_EVPROD_INFO":
			// begin of the info. all the following lines should go into e.Info
			info = true
		case s == "END_EVPROD_INFO":
			// end of the info
			info = false
		}
	}
	return e
}
