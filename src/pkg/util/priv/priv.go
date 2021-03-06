// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package priv

import (
	"os"
	"syscall"
)

// Escalate escalates thread privileges
func Escalate() error {
	uid := os.Getuid()
	if err := syscall.Setresuid(uid, 0, uid); err != nil {
		return err
	}
	return nil
}

// Drop drops thread privileges
func Drop() error {
	uid := os.Getuid()
	if err := syscall.Setresuid(uid, uid, 0); err != nil {
		return err
	}
	return nil
}
