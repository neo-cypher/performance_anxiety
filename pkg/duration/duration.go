// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package duration

import (
	"time"

	"github.com/zeebo/errs"
)

var durationErr = errs.Class("Duration")

func Duration(invocationTime time.Time, name string) {
	elapsed := time.Since(invocationTime)
	durationErr.New("%s took %s to execute.", name, elapsed)
} // end of Duration
