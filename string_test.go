// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/essentials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package ess

import (
	"testing"

	"aahframework.org/test.v0/assert"
)

func TestStringEmptyNotEmpty(t *testing.T) {
	assert.Equal(t, false, IsStrEmpty("    Welcome    "))

	assert.Equal(t, true, IsStrEmpty("        "))
}
