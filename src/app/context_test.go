package app

import (
	"testing"
)

// just fo coverage
func TestContext_BuildConfigDefault(t *testing.T) {

	ctx := NewContext(nil)
	*ctx.Vars() = Vars{}
	_ = ctx.Vars()
}
