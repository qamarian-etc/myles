package main

/* This component implements an HTTP(S) interface for your mylware.

        HOW IT WORKS
        When this component receives an HTTP request, it does the following:

                1. Checks its routing table to see who to direct the request to;
                2. If it finds a controller to handle the request, it handles the request over to
                        the controller, but if there are no controller to handle the request, error
                        404 is returned;
                3. The controller can then decide what to do with the HTTP request.

                ** Controllers can be created in file "aaaaal.zController.go", and registered in
                        file "aaaaal.zRouter.go".

        DEPENDENCIES
        Comp AAAAAC (Customized output assistant)
        Comp AAAAAG (Configuration data provider)
        Comp AAAAAI (Critical Event Zain)
        Comp AAAAAK (Logger)

	NOTES
	1. Ensure the following conf data are registered with comp AAAAAG:
		AAAAAL.NetAddr
		AAAAAL.NetPort
		AAAAAL.MaxDurationForNetIO
		AAAAAL.HttpMaxReqHeaderSize
		AAAAAL.TLSCertBundle
		AAAAAL.TLSPrivateKey */

import (
        "context"
	"fmt"
        "github.com/gorilla/mux"
	"net/http"
        "os"
	"strconv"
        "time"
)

func init () {
        iRegComp_AAAAAA ("aaaaal", iInit_AAAAAL, iDnit_AAAAAL, []string {"aaaaac", "aaaaag", "aaaaai",
                "aaaaak"})
}

func iInit_AAAAAL () {
	// Fetching network address and port to use. { ...
	netAddr, errA := iScalarData_AAAAAG ("AAAAAL.NetAddr")
	netPort, errB := iScalarData_AAAAAG ("AAAAAL.NetPort")

	if errA != nil {
		output := fmt.Sprintf ("Tried fetching value of 'AAAAAL.NetAddr' from the conf " +
                                "file, but an error occured. [%s] (HTTP Interface)", errA.Error ())
		iOutputT2_AAAAAC ("aaaaal", "err", output)
		os.Exit (1)
	}
	if errB != nil {
		output := fmt.Sprintf ("Tried fetching value of 'AAAAAL.NetPort' from the conf " +
                                "file, but an error occured. [%s] (HTTP Interface)", errB.Error ())
		iOutputT2_AAAAAC ("aaaaal", "err", output)
		os.Exit (1)
	}
	// ... }

	// Fetching max duration allowed for net input/output. { ...
        netIODurationBeforeTimeout, errP := iScalarData_AAAAAG ("AAAAAL.MaxDurationForNetIO")
        if errP != nil {
                output := fmt.Sprintf ("Tried fetching value of 'AAAAAL.MaxDurationForNetIO' from" +
                        " the conf file, but an error occured. [%s] (HTTP Interface)",
                        errP.Error ())
                iOutputT2_AAAAAC ("aaaaal", "err", output)
                os.Exit (1)
        }

        // Casting duration from string to int. { ...
        intNetIOMaxDuration, errH := strconv.Atoi (netIODurationBeforeTimeout)
        if errH != nil {
                output := fmt.Sprintf ("Tried casting value of 'AAAAAL.MaxDurationForNetIO' into " +
                        "an integer, but I couldn't. Are you sure an integer value was provided " +
                        "in the conf file? [%s] (HTTP Interface)", errH.Error ())
                iOutputT2_AAAAAC ("aaaaal", "err", output)
                os.Exit (1)
        }
        // ... }
        // ... }

        // Fetching max size allowed for HTTP request headers. { ...
        httpMaxReqHeaderSize, errQ := iScalarData_AAAAAG ("AAAAAL.HttpMaxReqHeaderSize")
        if errQ != nil {
        	output := fmt.Sprintf ("Tried fetching value of 'AAAAAL.HttpMaxReqHeaderSize' " +
                        "from the conf file, but an error occured. [%s] (HTTP Interface)",
                        errQ.Error ())
                iOutputT2_AAAAAC ("aaaaal", "err", output)
                os.Exit (1)
        }

        // Casting max header size from string to int. { ...
        intReqHeaderSize, errI := strconv.Atoi (httpMaxReqHeaderSize)
        if errI != nil {
        	output := fmt.Sprintf ("Tried casting value of 'AAAAAL.HttpMaxReqHeaderSize' into" +
                        " an integer, but I couldn't. Are you sure an integer value was provided " +
                        "in the conf file? [%s] (HTTP Interface)", errI.Error ())
                iOutputT2_AAAAAC ("aaaaal", "err", output)
                os.Exit (1)
        }
        // ... }
        // ... }

        // Composing HTTP server details to be used. { ...
	dServerInfo_AAAAAL := &http.Server {
                Addr:            netAddr + ":" +  netPort,
                ReadTimeout:     time.Duration (intNetIOMaxDuration) * time.Second,
                WriteTimeout:    time.Duration (intNetIOMaxDuration) * time.Second,
                MaxHeaderBytes:  intReqHeaderSize,
        }

        router := mux.NewRouter ()

        // Registering routing rules with HTTP router.
        for _, rule := range dRoutingRule_AAAAAL {
                router.HandleFunc (rule.route, rule.controller)
        }

        dServerInfo_AAAAAL.Handler = router
        // ... }

        // Determining if HTTP or HTTPS should be be used. { ...

        // Fetching filepath of TLS cert bundle. { ...
        certBundle, errO := iScalarData_AAAAAG ("AAAAAL.TLSCertBundle")
        if errO != nil {
        	output := fmt.Sprintf ("Tried fetching value of 'AAAAAL.TLSCertBundle' from the " +
                        "conf file, but an error occured. [%s] (HTTP Interface)", errO.Error ())
                iOutputT2_AAAAAC ("aaaaal", "err", output)
                os.Exit (1)
        }
        // ... }

        // Fetching filepath of TLS private key. { ...
        privateKey, errT := iScalarData_AAAAAG ("AAAAAL.TLSPrivateKey")
        if errT != nil {
        	output := fmt.Sprintf ("Tried fetching value of 'AAAAAL.TLSPrivateKey' from the " +
                        "conf file, but an error occured. [%s] (HTTP Interface)", errT.Error ())
                iOutputT2_AAAAAC ("aaaaal", "err", output)
                os.Exit (1)
        }
        // ... }

        httpProtocolInUse := "" // Declaration of a variable that'll be used later.

        if certBundle != "" && privateKey != "" {
                httpProtocolInUse = "HTTPS"
        } else {
                httpProtocolInUse = "HTTP"
                iOutputT2_AAAAAC ("aaaaal", "wrn", "HTTP in use, since value of " +
                        "AAAAAL.TLSCertBundle and/or AAAAAL.TLSPrivateKey are not set. (HTTP " +
                        "Interface)")
        }
        // ... }

        // Startup notification.
        iOutputT2_AAAAAC ("aaaaal", "std", fmt.Sprintf ("Net Addr: %s:%s (HTTP Interface)",
		netAddr, netPort))

        // Starting server, as a different routine.
        go func () {
        	errJ := *new (error) // Declaration of a variable that'll be used later.

        	if httpProtocolInUse == "HTTPS" { // HTTPS startup.
                        // Note, this function blocks.
                	errJ = dServerInfo_AAAAAL.ListenAndServeTLS (certBundle, privateKey)
        	} else { // HTTP startup.

	                errJ = dServerInfo_AAAAAL.ListenAndServe () // Note, this function blocks.
        	}

                // By the time execution reachs here, the server must have shutdown.
	        iOutputT2_AAAAAC ("aaaaal", "std", "State: Server has shutdown!")

	        /* If server shutdowns due to an error, a log is recorded, and a critical event
                        zain is notified. */
        	if errJ != nil && errJ != http.ErrServerClosed {
                        errMssg := fmt.Sprintf ("I've shutdown, error occured while doing my" +
                                " job. [%s] (HTTP Interface)", errJ.Error ())
                	iOutputT2_AAAAAC ("aaaaal", "err", errMssg)
                        iBeInformed_AAAAAI ("AAAAAL: " + errMssg)
	                iRecord_AAAAAK ("AAAAAL: " + errMssg)
        	}
        } ()
}

func iDnit_AAAAAL () {
        iShutdown_AAAAAL ()
}

func iShutdown_AAAAAL () {
        // If a panic should occur, it is prevented from affecting other components.
        defer func () {
                recover ()
        } ()

        // Shutting down server gracefully. { ...
        errX := dServerInfo_AAAAAL.Shutdown (context.Background ())
        if errX != nil {
                errMssg := fmt.Sprintf ("AAAAAL: Graceful shutdown failed. [%s] (HTTP Interface)",
                        errX.Error ())
                iRecord_AAAAAK (errMssg)
        }
        // ... }
}

var (
	dServerInfo_AAAAAL *http.Server // Information needed to run the interface's server.
)

type tController_AAAAAL func (http.ResponseWriter, *http.Request)

func (controller tController_AAAAAL) ServeHTTP (res http.ResponseWriter, req *http.Request) {
        controller (res, req)
}

type tRoutingRule_AAAAAL struct {
	route string
	controller tController_AAAAAL
}
