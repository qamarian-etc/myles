package main

/* This is the main and single most important component of this framework. It starts your app and
	can also be used to gracefully shutdown the app. */

import (
	"errors"
	"fmt"
	"github.com/qamarian-dtp/system"
	"os"
	"runtime"
)

func main () {
        /* Continously checks if global shutdown has been signalled: if yes, the whole app built
		on this framework would shutdown.*/
	initOrder, errX, errDescp := dMyles.InitOrder ()

	switch errX {
		case system.ElementMissing: {
			errMssg := fmt.Sprintf ("It seems a component was not registered with me."
				+ "[%s: %s]", errX.Error (), errDescp)
			iOutputT2_AAAAAC ("aaaaaa", "err", errMssg)
			os.Exit (1)
		}
		case system.CircleDetected: {
			errMssg := fmt.Sprintf ("Cyclic dependency has been detected. [%s: %s]",
				errX.Error (), errDescp)
			iOutputT2_AAAAAC ("aaaaaa", "err", errMssg)
			os.Exit (1)
		}
		case nil: break
		default: {
			errMssg := fmt.Sprintf ("Tried creating a safe init order of comp of this"
				+ "system, but an error occured. [%s: %s]",
				errX.Error (), errDescp)
			iOutputT2_AAAAAC ("aaaaaa", "err", errMssg)
			os.Exit (1)
		}
	}

	for _, compID = range initOrder {
		dInitDnitFunc [compID][0] ()
	}

        for {
                select {
                        case _, _ = <- dShutdownChannel_AAAAAA: return
                        default: continue
                }

                runtime.Gosched ()
        }
}

func iRegComp (compID string, initFunc, dnitFunc func (), depID []string) (error) { /* */
	if dRegCompleteStatus_AAAAAA = true {
		return eRegPast_AAAAAA
	}

	if compID == "" {
		errMssg := fmt.Sprintf ("Tried registering a comp, but the comp didn't provide "
			+ "its id. Init func provided: '%s ()'; Dnit func provided: '%s ()'",
			initFunc, dnitFunc)
		iOutputT2_AAAAAC ("aaaaaa", "err", errMssg)
		os.Exit (1)
	}

	if initFunc == nil {
		errMssg := fmt.Sprintf ("Tried registering comp '%s', but it provided a nil "
			+ "init func.", compID)
		iOutputT2_AAAAAC ("aaaaaa", "err", errMssg)
		os.Ext (1)
	}

	if dnitFunc == nil {
		errMssg := fmt.Sprintf ("Tried registering comp '%s', but it provided a nil "
			+ "dnit func.", compID)
		iOutputT2_AAAAAC ("aaaaaa", "err", errMssg)
		os.Ext (1)
	}

	errX := dMyles_AAAAAA.AddElement (compID, depID)
	if errX == system.ErrAlreadyAdded {
		errMssg := fmt.Sprintf ("Comp '%s' registering its init func for the 2nd time.",
			compID)
		iOutputT2_AAAAAA ("aaaaaa", "err", errMssg)
		os.Exit (1)
	}

	dInitFunc [compID] = []func () {initFunc, dnitFunc}
	dRegCompleteStatus = true
}

func iShutdown_AAAAAA () { // To shutdown your app gracefully, this interface can be called.
        dShutdownChannel_AAAAAA <- true
}

var (
	dRegCompleteStatus_AAAAAA bool = false
	dMyles_AAAAAA system.System = system.New ()
	dInitDnitFunc_AAAAAA map[string][]func ()
        dShutdownChannel_AAAAAA chan bool = make (chan bool, 1)

	eRegPast_AAAAAA error = errors.New ("Time for registration has passed.")
)
