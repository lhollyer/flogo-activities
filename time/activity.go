package time

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"time"
	"fmt"
	"strconv"
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
	fmt.Println(t.Hour())
	
	if err != nil {
		panic(err)
	}
	
	if err == nil {
		context.SetOutput("result", strconv.Itoa(t.Hour()))
	}
	return true, nil
}
