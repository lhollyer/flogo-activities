package hx711

import (
    "github.com/TIBCOSoftware/flogo-lib/core/activity"
//    "github.com/TIBCOSoftware/flogo-lib/logger"
    "fmt"
    "os/exec"
    "log"
    "strconv"
    "strings"
"bytes"
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
/*    cmd := exec.Command("/home/pi/hx711pyFINAL/middle.py")
out, err := cmd.Output()

if err != nil {
	println(err.Error())
	return
}


fmt.Println(string(out))
*/


cmd := exec.Command("python", "/home/pi/hx711pyFINAL/middle.py")
var out bytes.Buffer
var stderr bytes.Buffer
cmd.Stdout = &out
cmd.Stderr = &stderr
erro := cmd.Run()
if erro != nil {
    fmt.Println(fmt.Sprint(erro) + ": " + stderr.String())
    return
}
// fmt.Println("Result: " + out.String())



//    initial := string(out)
initial := out.String()
// fmt.Println(initial)

    s := strings.TrimSpace(initial)
// fmt.Println(s)

    b, err := strconv.Atoi(s)
// fmt.Println(b)

    if err != nil {
                    log.Fatal(err)
    }
    fmt.Println(b)
    context.SetOutput("result", b)



    return true, nil
}

/*package hx711

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
*/
