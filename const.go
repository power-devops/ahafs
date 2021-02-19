// Copyright 2021 Power-Devops.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ahafs

// AHAFS return codes from event producers
const (
	EvProd               = 0    // Most of the event producers return just 0.
	ModFileWrite         = 1000 // The monitored file was written to.
	ModFileUnmount       = 1001 // The file system containing the monitored file was unmounted. This is an unavailable event.
	ModFileMap           = 1002 // A process has mapped a portion of the monitored file for writing.
	ModFileRemove        = 1003 // The monitored file has been removed. This is an unavailable event.
	ModFileRename        = 1004 // The monitored file has been renamed. This is an unavailable event.
	ModFileFclear        = 1005 // A process has issued an fclear for the monitored file.
	ModFileFtrunc        = 1006 // A process has issued an ftrunc for the monitored file.
	ModFileOvermount     = 1007 // The monitored file has been over mounted.
	ModFileAttrUnmount   = 1001 // The filesystem containing the monitored file or directory was unmounted. This is an unavailable event.
	ModFileAttrRemove    = 1003 // The monitored file or directory has been removed. This is an unavailable event.
	ModFileAttrRename    = 1004 // The monitored file or directory has been renamed. This is an unavailable event.
	ModFileAttrOvermount = 1007 // The monitored file or directory has been over mounted. This is an unavailable event.
	ModFileAttrSetacl    = 1008 // The ACLs of the monitored file or directory were modified.
	ModFileAttrSetown    = 1009 // The ownership of the monitored file or directory was modified.
	ModFileAttrSetmode   = 1010 // The mode of the monitored file or directory was modified.
	ModDirCreate         = 1000 // A new file system object has been created under the monitored directory.
	ModDirUnmount        = 1001 // The file system containing the monitored directory has been unmounted. This is an unavailable event.
	ModDirRemove         = 1002 // A file system object within the monitored directory has been removed.
	ModDirRemoveSelf     = 1003 // The monitored directory itself has been removed or renamed. This is an unavailable event.
	UtilFsTheeshHit      = 1000 // The file system being monitored has reached the threshold specified.
	UtilFsUnmount        = 1001 // The file system being monitored has been unmounted. This is an unavailable event.
)
