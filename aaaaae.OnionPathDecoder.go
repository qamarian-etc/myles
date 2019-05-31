package main

/* This component decodes an onion filepath into its real form. If the decoded filepath is a symlink,
	the symlink will be further evaluated into its extremely-real form.

	To use this component, simply call its interface "iDecode_AAAAAE ()". */

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init () {
	iRegComp ("aaaaae", iInit_AAAAAE, iDnit_AAAAAE, nil)
}

func iInit_AAAAAE () {}

func iDnit_AAAAAE () {}

func iDecode_AAAAAE (onionPath string) (path string, err error) { /* This interface decodes an onion
	filepath, into its genuine form.

	In addition to decoding onion-formatted filepaths, if an onion path decodes to a symbolic
		link, this function further evaluates the symbolic link into its extremely-real
		form.

	EXPLANATION
	If an onion filepath (lets say "./file.ext") decodes to "/pathA/file.ext", and this
		filepath (/pathA/file.ext) is a symlink which points to "/pathB/file.ext"
		(another symlink), and "/pathB/file.ext" further points to "/pathC/file.ext", which
		is a real filepath, then the input of "./file.ext" would result into the output of
		"/pathC/file.ext".

	INPUT
	input 0: The onion filepath to be decoded.

	OUTPT
	outpt 0: The decoded form of input 0. On successful decoding, value would be the decoded
		form of input 0. On failed decoding, value would be an empty string.

	outpt 1: Any error that occurs during the decoding. On successful decoding, value would be
		nil. On failed decoding, value would the error that occured. */

	// If a panic should occur, it is prevented from affecting other components. { ...
	err = errors.New ("aaaaae: iDecode_AAAAAE () paniced.") /* This error will be returned if a
		panic should occur. */

	defer func () {
		panicReason := recover ()

		if panicReason != nil {
			return
		}
	} ()
	// ... }

	// If the filepath starts with "*/", "*/" is replaced with the actual directory of the
		/* program's file (the directory where the executable file of the program is
		located). */
	if strings.Index (onionPath, "*/") == 0 {
		programFileDir, errX := os.Executable ()
		if errX != nil {
			errMssg := fmt.Sprintf ("aaaaae: Tried fetching the dir of this app's " +
				"executable file, but I couldn't. [%s]", errX.Error ())
			return "", errors.New (errMssg)
		}

		pathSeparator := fmt.Sprintf ("%c", os.PathSeparator)
		onionPath = strings.Replace (onionPath, "*/",
			(filepath.Dir (programFileDir) + pathSeparator), 1)

	} else if strings.Index (onionPath, "./") == 0 {
	// If the filepath starts with "./", "./" is replaced with the present working directory.

		presentWorkingDir, errY := os.Getwd ()
		if errY != nil {
			errMssg := fmt.Sprintf ("aaaaae: Tried fetching the working dir of this" +
				" app, but I couldn't. [%s]", errY.Error ())
			return "", errors.New (errMssg)
		}

		pathSeparator := fmt.Sprintf ("%c", os.PathSeparator)
		onionPath = strings.Replace (onionPath, "./", (presentWorkingDir + pathSeparator),
			1)
	}

	/* In case the real form of the onion filepath is a symlink, the symlink will be evaluated
		into its extreme real form. */
	actualFilepath, errZ := filepath.EvalSymlinks (onionPath)
	if errZ != nil {
		errMssg := fmt.Sprintf ("aaaaae: Tried evaluating '%s' into a genuine filepath " +
			"(just in case its a symlink), but I couldn't. Are you sure file '%s' " +
			"exists? [%s]", onionPath, onionPath, errZ.Error ())
		return "", errors.New (errMssg)
	}

	return actualFilepath, nil
}
