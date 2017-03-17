package source

import (
	"fmt"

	"github.com/corvinusy/mygolangtoys/facade-example/ctx"
	"github.com/corvinusy/mygolangtoys/facade-example/dest"
)

// Source export
type Source struct {
	ctx *ctx.Context
}

// New contructor
func New(v int) *Source {
	s := new(Source)
	s.ctx = new(ctx.Context)
	s.ctx.Value = v
	return s
}

// Run ...
func (s *Source) Run() {
	dest := dest.Dest{C: s.ctx, F: s}
	dest.Do()
}

// ShowAddress ...
func (s *Source) ShowAddress() {
	fmt.Printf("source address is %d\n", &s)
}

// ShowContent ...
func (s *Source) ShowContent() {
	fmt.Printf("source context value address is %d\n", &s.ctx.Value)
}
