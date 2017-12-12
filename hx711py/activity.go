package hx711

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
//	"github.com/TIBCOSoftware/flogo-lib/logger"
	"fmt"
	"os/exec"
	"log"
	"strconv"
	"strings"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	out, err := exec.Command("/home/pi/hx711pyFINAL/middle.py").Output()

	initial := string(out)
	s := strings.TrimSpace(initial)
	b, err := strconv.Atoi(s)

	if err != nil {
					log.Fatal(err)
	}
	fmt.Println(b)
	context.SetOutput("result", b)

	return true, nil
}
