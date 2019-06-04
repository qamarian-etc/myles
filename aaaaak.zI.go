package main

/* This component is a logger.

	When different goroutines simultanously ask it to record logs, it queues all the requests,
	then records them one-by-one, instead of trying to write them to the log file all at once
	(thereby degrading the performance of the app using it).

	In short, this logger is designed not to degrade the performance of apps using it, unlike
		most other loggers.

	USAGE NOTES
	1. When about starting an app whose source code contains this component, do the following:

   		- create a log file, on the computer on which the app would run
   		- set the logfile's path as conf data "AAAAAK.LogfilePath", in the conf file of
   			dependency AAAAAG (remember "AAAAAG" was listed as one of the dependencies
   			needed by this component)

   	2. To tell this component to record a log, use interface "iRecord_AAAAAK ()". */

import (
	"fmt"
	qamarian_Logger "github.com/qamarian-inc/logger"
	"os"
	"strings"
)

func init () {
	iRegComp_AAAAAA ("aaaaak", iInit_AAAAAK, iDnit_AAAAAK, []string {"aaaaac", "aaaaae", "aaaaag",
		"aaaaai"})
}

func iInit_AAAAAK () {
	// Getting the filepath of the log file. { ...
	logfilePath, errT := aaaaakDPConfProvider.ScalarData ("AAAAAK.LogfilePath")

	if errT == dErrNotSet_AAAAAG {
		errMssg := fmt.Sprintf ("My logfile path is not set with comp AAAAAG. (Logger)")
		aaaaakDPOutAssist.Output ("aaaaak", "err", errMssg)
		os.Exit (1)
	} else if errT != nil {
		errMssg := fmt.Sprintf ("Unable to fetch the filepath of my log file. [%s] " +
			"(Logger)", errT.Error ())
		aaaaakDPOutAssist.Output ("aaaaak", "err", errMssg)
		os.Exit (1)
	}
	// ... }

	logfilePath = strings.Trim (logfilePath, " ")

	// Halts app, if an empty string is set as the logfile's path.
	if logfilePath == "" {
		aaaaakDPOutAssist.Output ("aaaaak", "err", "Empty string was provided as my logfile " +
			"path, in comp AAAAAG's conf file. (Logger)")
		os.Exit (1)
	}

	/* The filepath of your app's log file is expected to be onion-formatted, and this section
		does the decoding of the filepath into its genuine form. { ... */
	logfilePath, errV := iDecode_AAAAAE (logfilePath)
	if errV != nil {
		errMssg := fmt.Sprintf ("Unable to decode my onion-formatted logfile path. " +
			"Are you sure you provided me a valid onion filepath? [%s] (Logger)",
			errV.Error ())
		aaaaakDPOutAssist.Output ("aaaaak", "err", errMssg)
		os.Exit (1)
	}
	// ... }

	dLoggingInfo_AAAAAK = &qamarian_Logger.Logging_Info {Log_File: logfilePath} /* Creating
		data needed to run the external logger. */

	// Starting the logger.
	go cAAAAAK ()
}

func cAAAAAK () {
	// If a panic should occur, it is prevented from affecting other routines.
	defer func () {
		recover ()
	} ()

        errM := dLoggingInfo_AAAAAK.Logger (0)

        if errM != nil {
        	errMssg := fmt.Sprintf ("Logger has shutdown due to an error. [%s] (Logger)",
        		errM.Error ())
                aaaaakDPOutAssist.Output ("aaaaak", "wrn", errMssg)
                aaaaakDPCEZain.BeInformed ("AAAAAK: " + errMssg)
        }
}

func iDnit_AAAAAK () {
	iShutdown_AAAAAK ()
}

func iRecord_AAAAAK (new_Log string) (error) { /* Call this interface, to record a new log.

	INPUT
	input 0: The log to be recorded.

	OUTPT
	outpt 0: Any error that occurs during the logging. If logging succeeds, value would
		be nil. If logging fails, value would the error that occured. */

	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	// Recording log.
	errX := dLoggingInfo_AAAAAK.Log (new_Log)

	return errX
}

func iShutdown_AAAAAK () {
	// If a panic should occur, it is prevented from affecting caller of this function.
	defer func () {
		recover ()
	} ()

	dLoggingInfo_AAAAAK.Shutdown ()
}

var (
	dLoggingInfo_AAAAAK *qamarian_Logger.Logging_Info
)
