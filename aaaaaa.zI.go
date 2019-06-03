package main

/* This is the main and the most important component of this framework. It starts your app and can
	also be used to gracefully shut it down.

	In addition, this component also help coordinate the initialization of the mylcoms of your
		cloud app or service.

	DEPENDENCY
		Comp AAAAAC (Easy-to-read Output) */

import (
	"errors"
	"fmt"
	"github.com/qamarian-etc/slices"
	"github.com/qamarian-dtp/system"
	"os"
	"runtime"
	"strings"
)

func main () {/* This component does three things: it coordinates the init of other components in
	your app, it keeps your app running until shutdown has been signalled, and it coordinates
	the dnit of other components in your app. */

	dRegCompleteStatus_AAAAAA = true

        // Coordinattion of init. { ...
	initOrder, errX, errDescp := dMyles_AAAAAA.InitOrder ()

	switch errX {
		case system.ErrElementMissing: {
			errMssg := fmt.Sprintf ("It seems a component was not registered with " +
			"me. [%s: %s]", errX.Error (), errDescp)
			dpOutAssit.Output ("aaaaaa", "err", errMssg)
			os.Exit (1)
		}
		case system.ErrCircleDetected: {
			errMssg := fmt.Sprintf ("Cyclic dependency has been detected. [%s: %s]" +
				errX.Error (), errDescp)
			dpOutAssit.Output ("aaaaaa", "err", errMssg)
			os.Exit (1)
		}
		case nil: break
		default: {
			errMssg := fmt.Sprintf ("Tried creating a safe init order of comp of " +
				"this system, but an error occured. [%s: %s]",
				errX.Error (), errDescp)
			dpOutAssit.Output ("aaaaaa", "err", errMssg)
			os.Exit (1)
		}
	}

	// Asking your app's mylcoms to init, one-after-the-other.
	for _, compID := range initOrder {
		dInitDnitFunc_AAAAAA [compID][0] ()
	}
	// ... }

	dpOutAssit.Output ("aaaaaa", "std", "Cloud app/service now running!")

	// Keeps your app running until shutdown has been signalled.
        for {
                select {
                        case _, _ = <- dShutdownChannel_AAAAAA: return
                        default: continue
                }
                runtime.Gosched ()
        }
}

func iRegComp_AAAAAA (compID string, initFunc, dnitFunc func (), depID []string) (error) { /* To
	register a component, this interface can be called. If a registration fails when your app
	has just started, this interface halts the app. If a registration fails quite long after
	your app has started, this interface returns an error.

	INPUT
	input 0: The ID of the component trying to register itself.
	input 1: The init func of the comp. Nil not allowed.
	input 2: The dnit func of the comp. Nil not allowed.
	input 3: An array of the IDs of all comps depended on.

	OUTPT
	outpt 0: If registration succeeds, value would be nil error. Otherwise, value will be an
		actual error. Possible errors include: eRegPast_AAAAAA. */

	if dRegCompleteStatus_AAAAAA == true {
		return dErrRegPast_AAAAAA
	}
	if compID == "" {
		errMssg := fmt.Sprintf ("Tried registering a comp, but the comp failed to pr" +
		"ovide its id.")
		dpOutAssit.Output ("aaaaaa", "err", errMssg)
		os.Exit (1)
	}
	if initFunc == nil {
		errMssg := fmt.Sprintf ("Tried registering comp '%s', but it provided a nil " +
			"init func.", compID)
		dpOutAssit.Output ("aaaaaa", "err", errMssg)
		os.Exit (1)
	}
	if dnitFunc == nil {
		errMssg := fmt.Sprintf ("Tried registering comp '%s', but it provided a nil " +
			"dnit func.", compID)
		dpOutAssit.Output ("aaaaaa", "err", errMssg)
		os.Exit (1)
	}

	// Converting IDs to lower case. Just to make registration case-insensitive. { ...
	compID = strings.ToLower (compID)
	idClone := []string {}
	for _, id := range depID {
		idClone = append (idClone, strings.ToLower (id))
	}
	depID = idClone
	// ... }

	errX := dMyles_AAAAAA.AddElement (compID, depID)
	if errX == system.ErrAlreadyAdded {
		errMssg := fmt.Sprintf ("Comp '%s' registering itself for the 2nd time.", compID)
		dpOutAssit.Output ("aaaaaa", "err", errMssg)
		os.Exit (1)
	}
	dInitDnitFunc_AAAAAA [compID] = []func () {initFunc, dnitFunc}
	return nil
}

func iShutdown_AAAAAA () { /* To shutdown your app gracefully, this interface can be called.
	However, before shutdown starts, all mylcoms registered would be asked to dnit. The dnit
	would be done in the reverse order of init. This way dependencies would dnit last. */

	initOrder, _, _ := dMyles_AAAAAA.InitOrder ()
	dnitOrder := slices.RevStringSlice (initOrder)
	for _, compID := range dnitOrder {
		dInitDnitFunc_AAAAAA [compID][1] ()
	}
        dShutdownChannel_AAAAAA <- true
}

var (
	dRegCompleteStatus_AAAAAA bool = false
	dMyles_AAAAAA *system.System = system.New ()
	dInitDnitFunc_AAAAAA map[string][]func () = map[string][]func () {}
        dShutdownChannel_AAAAAA chan bool = make (chan bool, 1)
	dErrRegPast_AAAAAA error = errors.New ("Registration time has passed.")
)
