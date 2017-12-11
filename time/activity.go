package time

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"time"
	"fmt"
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
	t := time.Now()
	fmt.Println(t.Format("20060102150405"))
	context.SetOutput("result", t.Format("20060102150405"))
	return true, nil
}
