package main

/* DEPENDENCIES
        Comp AAAAAC (Customized output assistant)
        Comp AAAAAG (Configuration data provider)
        Comp AAAAAI (Critical Event Zain)
        Comp AAAAAK (Logger) */

var aaaaalDPOutAssist interface {
	Output (string, string, string, ... string)
} = iNewLSIA_AAAAAC ()

var aaaaalDPConfProvider interface {
	ScalarData (string) (string, error)
	SliceData (string) ([]string, error)
	MapData (string) (map[string]string, error)
} = iNewLSIA_AAAAAG ()

var aaaaalDPCEZain interface {
	BeInformed (string)
} = iNewLSIA_AAAAAI ()

var aaaaalDPLogger interface {
	Record (string) (error)
} = iNewLSIA_AAAAAK ()
