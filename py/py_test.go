// Copyright 2018 The go-python Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package py_test

import (
	"testing"

	"github.com/goplus/pyg/pytest"
)

func TestPy(t *testing.T) {
	pytest.RunTests(t, "tests")
}
