// Copyright 2021 Power-Devops.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ahafs

const (
	MODFILE_WRITE         = 1000 // The monitored file was written to.
	MODFILE_UNMOUNT       = 1001 // The file system containing the monitored file was unmounted. This is an unavailable event.
	MODFILE_MAP           = 1002 // A process has mapped a portion of the monitored file for writing.
	MODFILE_REMOVE        = 1003 // The monitored file has been removed. This is an unavailable event.
	MODFILE_RENAME        = 1004 // The monitored file has been renamed. This is an unavailable event.
	MODFILE_FCLEAR        = 1005 // A process has issued an fclear for the monitored file.
	MODFILE_FTRUNC        = 1006 // A process has issued an ftrunc for the monitored file.
	MODFILE_OVERMOUNT     = 1007 // The monitored file has been over mounted.
	MODFILEATTR_UNMOUNT   = 1001 // The filesystem containing the monitored file or directory was unmounted. This is an unavailable event.
	MODFILEATTR_REMOVE    = 1003 // The monitored file or directory has been removed. This is an unavailable event.
	MODFILEATTR_RENAME    = 1004 // The monitored file or directory has been renamed. This is an unavailable event.
	MODFILEATTR_OVERMOUNT = 1007 // The monitored file or directory has been over mounted. This is an unavailable event.
	MODFILEATTR_SETACL    = 1008 // The ACLs of the monitored file or directory were modified.
	MODFILEATTR_SETOWN    = 1009 // The ownership of the monitored file or directory was modified.
	MODFILEATTR_SETMODE   = 1010 // The mode of the monitored file or directory was modified.
	MODDIR_CREATE         = 1000 // A new file system object has been created under the monitored directory.
	MODDIR_UNMOUNT        = 1001 // The file system containing the monitored directory has been unmounted. This is an unavailable event.
	MODDIR_REMOVE         = 1002 // A file system object within the monitored directory has been removed.
	MODDIR_REMOVE_SELF    = 1003 // The monitored directory itself has been removed or renamed. This is an unavailable event.
	UTILFS_THRESH_HIT     = 1000 // The file system being monitored has reached the threshold specified.
	UTILFS_UNMOUNT        = 1001 // The file system being monitored has been unmounted. This is an unavailable event.
)
