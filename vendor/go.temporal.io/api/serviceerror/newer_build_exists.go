// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package serviceerror

import (
	"fmt"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// NewerBuildExists is returned to a poll request from a build that has been superceded by
	// a newer build in versioning metadata.
	NewerBuildExists struct {
		Message        string
		DefaultBuildID string
		st             *status.Status
	}
)

// NewNewerBuildExists returns new NewerBuildExists error.
func NewNewerBuildExists(defaultBuildID string) error {
	return &NewerBuildExists{
		Message:        fmt.Sprintf("Task queue has a newer compatible build: %q", defaultBuildID),
		DefaultBuildID: defaultBuildID,
	}
}

// Error returns string message.
func (e *NewerBuildExists) Error() string {
	return e.Message
}

func (e *NewerBuildExists) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.OutOfRange, e.Message)
	st, _ = st.WithDetails(
		&errordetails.NewerBuildExistsFailure{
			DefaultBuildId: e.DefaultBuildID,
		},
	)
	return st
}

func newNewerBuildExists(st *status.Status, errDetails *errordetails.NewerBuildExistsFailure) error {
	return &NewerBuildExists{
		Message:        st.Message(),
		DefaultBuildID: errDetails.GetDefaultBuildId(),
		st:             st,
	}
}
