package dest

import (
	"fmt"
	"time"

	"github.com/corvinusy/mygolangtoys/facade-example/ctx"
	"github.com/corvinusy/mygolangtoys/facade-example/source/facade"
)

// Dest ...
type Dest struct {
	C *ctx.Context
	F facade.SourceFacade
}

// Do ...
func (d *Dest) Do() {
	fmt.Println("dest.Do() started")
	dur := time.Duration(d.C.Value * 1e9)
	fmt.Println("dur =", dur)
	d.F.ShowAddress()
	d.F.ShowContent()
	time.Sleep(dur)
	d.showDestContentAddress()
	fmt.Println("dest.Do() finishing")
}

// showDestContentAddress ...
func (d *Dest) showDestContentAddress() {
	fmt.Printf("source context value address is %d\n", &d.C.Value)
}
