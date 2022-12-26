package xybotsim

import (
	"testing"
)

func TestCommandDrop(t *testing.T) {
	d := North
	r := newRobot(0, 0, 1)
	for i := 0; i < cap(r.commandQueue); i++ {
		r.EnqueueCommand(d)
	}
	if r.EnqueueCommand(d) == nil {
		t.Fatalf("EnqueueCommand() drop command silently.")
	}
}
