package main

/* This component can be used to output data to the console, in an easy-to-read format. The old
	interface "iOutput_AAAAAC ()" has now been deprecated in favour of the new interface
	"iOutputT2_AAAAAC ()" */

import (
	"fmt"
	"github.com/mgutz/ansi"
)

func init () {
	iRegComp_AAAAAA ("aaaaac", iInit_AAAAAC, iDnit_AAAAAC, nil)
}

func iInit_AAAAAC () {}

func iDnit_AAAAAC () {}

func iOutputT2_AAAAAC (compID, oType, output string, moreOutput ... string) { /* This interface
	can be used to output easy-to-read data. Unlike interface "iOutput_AAAAAC" which can be
	used to make one out at a time, this interface can be used to make multiple outputs at
	a time.

	INPUT
	input 0: The ID of the component trying make the output.

	input 1: The output type. Possible values are "std", "wrn", and "err". "wrn" stands for
		warning; warning outputs are outputs that indicate a warning message. "err"
		stands for error; error outputs are outputs that indicate an error. "std"
		stands for standard; standard outputs are normal outputs which are neither
		errors nor warnings. If this value is not any of "std", "wrn", and "err", value
		would be interpreted as "std".

	input 2: The output to be made.

	input 3: Additional outputs. Additional outputs would each be displayed on distinct
		lines. inputs 0 and 1 applies to all other additional outputs. */

	sOutput_AAAAAC (compID, oType, output)

	for _, output := range moreOutput {
		sOutput_AAAAAC (compID, oType, output)
	}
}

func iOutput_AAAAAC (output string) { /* This interface outputs easy-to-read outputs, and it has
	now been deprecated in favour of "iOutputT2_AAAAAC".

	INPUT
	input 0: The string to be outputted. */

	sOutput_AAAAAC ("mylcom", "std", output)
}

func sOutput_AAAAAC (compID, oType, output string) {
	if oType != "err" && oType != "wrn" {
		oType = "std"
	}

	colourCode := map[string]string {"std": "28+b", "wrn": "226+b", "err": "196+b"}
	code := colourCode [oType]

	fmt.Println (fmt.Sprintf ("\n    $ (%s): %s", ansi.Color (compID, code), output))
}
