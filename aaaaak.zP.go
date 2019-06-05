package main

/* DEPENDENCIES
   	Comp AAAAAC (Easy-to-read output assistant)
   	Comp AAAAAE (Onion path decoder)
   	Comp AAAAAG (Configuration data provider)
   	Comp AAAAAI (Critical Event Zain) */

var dLSIOutAssist_AAAAAK interface {
	Output (string, string, string, ... string)
} = iNewLSIA_AAAAAC ()

var dLSIOnionDecoder_AAAAAK interface {
	Decode (string) (string, error)
} = iNewLSIA_AAAAAE ()

var dLSIConfProvider_AAAAAK interface {
	ScalarData (string) (string, error)
	SliceData (string) ([]string, error)
	MapData (string) (map[string]string, error)
} = iNewLSIA_AAAAAG ()

var dLSICEZain_AAAAAK interface {
	BeInformed (string)
} = iNewLSIA_AAAAAI ()
