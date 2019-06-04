package main

/* DEPENDENCIES
        Comp AAAAAC (Customized output assistant)
        Comp AAAAAG (Configuration data provider)
        Comp AAAAAI (Critical Event Zain)
        Comp AAAAAK (Logger) */

var dpOutAssist interface {
	Output (string, string, string, ... string)
} = iNewLSIA_AAAAAC ()

var dpConfProvider interface {
	ScalarData (string) (string, error)
	SliceData (string) ([]string, error)
	MapData (string) (map[string]string, error)
} = iNewLSIA_AAAAAG ()

var dpCEZain interface {
	BeInformed (string)
} = iNewLSIA_AAAAAI ()

var dpLogger interface {
	Record (string) (error)
} = iNewLSIA_AAAAAK ()
