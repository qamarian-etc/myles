package main

/* This component makes possible the use of configuration file, in your app. The component uses the
	YAML syntax for its configuration file.

   	USAGE NOTE
   	1. Ensure the conf file specified in "dCONFILE_AAAAAG", is available on the machine where
		the app would run.

   	2. Any data set in your conf file, can be fetched using any of interfaces:
   		iScalarData_AAAAAG (),
   		iSliceData_AAAAAG (), and
   		iMapData_AAAAAG ().

   		The interface to use among the three, would depend on what kind of conf data you're
			trying to fetch.

   	3. This component comes with a default configuration file known as "aaaaag.dConf.yml".

   	4. Ensure constant "dCONFILE_AAAAAG" is set to its appropriate value.

   	5. If this mylcom is unable to start successfully, it halts your app/service.
*/

import (
	"errors"
	"fmt"
        viper_Interface "github.com/qamarian-inc/viper"
        "github.com/spf13/viper"
        "os"
)

const dCONFILE_AAAAAG string = "./aaaaag.zzConf.yml" /* You can modify this constant if your app's
	configuration file wouldn't be file "aaaaag.dConf.yml" or if file "aaaaag.dConf.yml"
	wouldn't be present in the app's current working directory. Filepath format supported:
	onion.*/

func init () {
	iRegComp_AAAAAA ("aaaaag", iInit_AAAAAG, iDnit_AAAAAG, []string {"aaaaac", "aaaaae"})
}

func iInit_AAAAAG () { // The initialization basically means caching the configuration file.
	// Decoding the configuration filepath from its onion form into its real form.
	conFilepath, errD := aaaaagDPOnionDecoder.Decode (dCONFILE_AAAAAG)
	if errD != nil {
		errMssg := fmt.Sprintf ("Could not decode my configuration file's filepath " +
			"into its genuine form. [%s] (Conf Data Provider)", errD.Error ())
		aaaaagDPOutAssist.Output ("aaaaag", "err", errMssg)
		os.Exit (1)
        }

        // Loading the configuration file.
	conf, errX := viper_Interface.New_Viper (conFilepath, "yaml")
	if errX != nil {
		errMssg := fmt.Sprintf ("Could not load my configuration file. [%s] (Conf " +
			"Data Provider)", errX.Error ())
		aaaaagDPOutAssist.Output ("aaaaag", "err", errMssg)
		os.Exit (1)
        }

        /* Configuration data is made gloabl, so as to become available to other functions in
        	this component. */
        dConf_AAAAAG = conf

        // Indicating that this component is now ready to provide its services.
        dAvailStatus_AAAAAG = true
}

func iDnit_AAAAAG () {}

func iScalarData_AAAAAG (dataName string) (string, error) { /* This interface helps fetch the value
	of a scalar data, from the configuration file.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data. Value would be an empty string if any error should occur
		during the fetch.

	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On
		failed fetch, value would the error that occured. If the data is not set, value
		would be error "dErrNotSet_AAAAAG". If the component has just started up, and it is
		yet to be available, value would be error "dErrNotAvail_AAAAAG". */

	if dAvailStatus_AAAAAG == false {
		return "", dErrNotAvail_AAAAAG
	}

	if ! dConf_AAAAAG.IsSet (dataName) {
		return "", dErrNotSet_AAAAAG
	}

	return dConf_AAAAAG.GetString (dataName), nil
}

func iSliceData_AAAAAG (dataName string) (data []string, err error) { /* This interface helps fetch the
	value of an array data, from the configuration file.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data. Value would be an empty slice if any error should occur during
		the fetch.

	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On
		failed fetch, value would the error that occured. If the data is not set, value would
		be error "dErrNotSet_AAAAAG". If the component has just started up, and it is yet to
		be available, value would be error "dErrNotAvail_AAAAAG". */

	if dAvailStatus_AAAAAG == false {
		return data, dErrNotAvail_AAAAAG
	}

	if ! dConf_AAAAAG.IsSet (dataName) {
		return []string {}, dErrNotSet_AAAAAG
	}

	return dConf_AAAAAG.GetStringSlice (dataName), nil
}

func iMapData_AAAAAG (dataName string) (data map[string]string, err error) { /* This interface helps
	fetch the value of a hash map data, from the configuration file.

	INPUT
	input 0: The name of the data to be fetched.

	OUTPT
	outpt 0: The value of the data. Value would be an empty map if any error should occur during the
		fetch.

	outpt 1: Any error that occurs during the fetch. On successful fetch, value would be nil. On
		failed fetch, value would the error that occured. If the data is not set, value would
		be error "dErrNotSet_AAAAAG". If the component has just started up, and it is yet to
		be available, value would be error "dErrNotAvail_AAAAAG". */

	if dAvailStatus_AAAAAG == false {
		return data, dErrNotAvail_AAAAAG
	}
	if ! dConf_AAAAAG.IsSet (dataName) {
		return data, dErrNotSet_AAAAAG
	}
	return dConf_AAAAAG.GetStringMapString (dataName), nil
}

var (
	dConf_AAAAAG *viper.Viper // Configuration data cache. It would be initialized by init ().
	dAvailStatus_AAAAAG bool = false /* This data signifies if this component is availble to
		provide its service or not. */

	dErrNotAvail_AAAAAG error = errors.New ("This component is not available yet.")
	dErrNotSet_AAAAAG error = errors.New ("The configuration data requested is not set.")
)
