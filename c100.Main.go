package main

import (
        "context"
	"fmt"
        "github.com/gorilla/mux"
	"net/http"
        "os"
	"runtime"
	"strconv"
        "time"
)

func main () { // This function helps the software start its server. All services added to variable 'software_Service' (in file ze30.Global.go), will be available via the server.

        software_Name, errW := conf_Data_Provider ("Software_Name")
	output (fmt.Sprintf ("Ware '%s' is starting up (with framework: Amanda)...", software_Name))

        // Initializing services requiring initialization at startup time.
        for _, service := range services_Init {
                service ()
        }

	// If a panic should occur, the panic is logged. { ...
        defer func () {
                panic_Reason := recover ()
                if panic_Reason != nil {
                        RECORD_LOG ("main () paniced")
                }
                os.Exit (1)
        } ()
        // ... }

	// Using the value of 'GOMAXPROCS' specified in the configuration file. { ...
	go_Max_Procs, errX := conf_Data_Provider ("Go_Max_Procs")
	if errX != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'Go_Max_Procs', from the configuration file '%s': main ()", CONF_FILE, errX.Error ()))
		os.Exit (1)
	}

	go_Max_Procs_As_Int, errY := strconv.Atoi (go_Max_Procs)
	if errY != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Converting value of 'Go_Max_Procs', from string to integer: main ()", CONF_FILE, errY.Error ()))
		os.Exit (1)
	}

	runtime.GOMAXPROCS (go_Max_Procs_As_Int)
	// ... }

	// Fetching network address and port of software's server. { ...
	software_Addr, errA := conf_Data_Provider ("Service_Addr.Addr")
	software_Port, errB := conf_Data_Provider ("Service_Addr.Port")

	if errA != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'Service_Addr.Addr', from the configuration file '%s': main ()", CONF_FILE, errA.Error ()))
		os.Exit (1)
	}
	if errB != nil {
		output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'Service_Addr.Port', from the configuration file '%s': main ()", CONF_FILE, errB.Error ()))
		os.Exit (1)
	}
	// ... }
	
	// Creating server. { ...
        net_IO_Duration_Before_Timeout, errP := conf_Data_Provider ("Server_Info.Net_IO_Duration_Before_Timeout")
                // Error handling.
                if errP != nil {
                        output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'Server_Info.Net_IO_Duration_Before_Timeout', from the configuration file '%s': main ()", errP.Error (), CONF_FILE))
                        os.Exit (1)
                }

        net_IO_As_Int, errH := strconv.Atoi (net_IO_Duration_Before_Timeout)
                // Error handling.
                if errH != nil {
                        output (fmt.Sprintf ("Startup Error: %s ---> \n Value of 'Server_Info.Net_IO_Duration_Before_Timeout', in the configuration file '%s', is not a number: main ()", errH.Error (), CONF_FILE))
                        os.Exit (1)
                }

        http_Max_Request_Header_Size, errQ := conf_Data_Provider ("Server_Info.HTTP_Max_Request_Header_Size")
                // Error handling.
                if errQ != nil {
                        output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'Server_Info.HTTP_Max_Request_Header_Size', from the configuration file '%s': main ()", errQ.Error (), CONF_FILE))
                        os.Exit (1)
                }

        request_Header_Size_As_Int, errI := strconv.Atoi (http_Max_Request_Header_Size)
                // Error handling.
                if errI != nil {
                        output (fmt.Sprintf ("Startup Error: %s ---> \n Value of 'Server_Info.Net_IO_Duration_Before_Timeout', in the configuration file '%s', is not a number: main ()", errI.Error (), CONF_FILE))
                        os.Exit (1)
                }

	main___server_Info := &http.Server {	
                Addr:            software_Addr + ":" + software_Port,
                ReadTimeout:     net_IO_As_Int * time.Second,
                WriteTimeout:    net_IO_As_Int * time.Second,
                MaxHeaderBytes:  request_Header_Size_As_Int,
        }

        routerX := mux.NewRouter ()
        for _, service := range software_Service {
                routerX.HandleFunc (service.Service_Path, service.Provider)
        }
        main___server_Info.Handler = routerX
        // ... }

        // Determining if HTTP or HTTPS should be be used.
        cert_Bundle, errO := conf_Data_Provider ("TLS_Res.Cert_Bundle")
                // Error handling.
                if errO != nil {
                        output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'TLS_Res.Cert_Bundle', from the configuration file '%s': main ()", errO.Error (), CONF_FILE))
                        os.Exit (1)
                }

        private_Key, errT := conf_Data_Provider ("TLS_Res.Private_Key")
                // Error handling.
                if errT != nil {
                        output (fmt.Sprintf ("Startup Error: %s ---> \n Fetching value of 'TLS_Res.Private_Key', from the configuration file '%s': main ()", CONF_FILE, errT.Error ()))
                        os.Exit (1)
                }

        http_Protocol_In_Use := ""

        if cert_Bundle != "" && private_Key != "" {
                http_Protocol_In_Use = "HTTPS"
        } else {
                http_Protocol_In_Use = "HTTP"
                output ("Note: Your app will be using HTTP (not HTTPS) since the filepath of one or both of the certificate bundle and private key are not set.")
        }
        // ... }

        // Startup notification.
        output (fmt.Sprintf ("App will start up now! NETWORK ADDRESS: %s:%s (%s)", software_Addr, software_Port, http_Protocol_In_Use))

        // Starting server. { ...
        errJ := *new (error)

        if http_Protocol_In_Use == "HTTPS" { // HTTPS startup.

                errJ = main___server_Info.ListenAndServeTLS (cert_Bundle, private_Key) // Note, this function blocks.
        } else { // HTTP startup.

                errJ = main___server_Info.ListenAndServe () // Note, this function blocks.
        }
        // ... }
                
        output ("State: Server has shutdown!")

        // If server shutdowns due to an error, a log is recorded, and the critical event action is executed.
        if errJ != nil && errJ != http.ErrServerClosed {
                output (fmt.Sprintf ("Runtime Error: %s ---> \n Running server: main ()", errJ.Error ()))

                RECORD_LOG (fmt.Sprintf ("%s ---> \n Running server: main ()", errJ.Error ()))

                CRITICAL_EVENT_ACTION (fmt.Sprintf ("Server has shutdown due to an error: %s", errJ.Error ()))

                // Gracefully shutting down server.
                main___Shutdown_Servers ()
        }

        // If server stopped because it was asked to, main () waits for os.Exit () to be called in main___Shutdown_Servers (), just to ensure graceful shutdown.
        if errJ == http.ErrServerClosed {
                for {
                        runtime.Gosched ()
                }
        }
}

func main___Shutdown_Servers () { // This function is an interface of main (); it tells main () to shutdown the server of this software. To gracefully shutdown this software, this function can be called.

        // If a panic should occur during shutdown, the panic is logged. { ...
        defer func () {
                panic_Reason := recover ()
                if panic_Reason != nil {
                        RECORD_LOG ("Paniced occured ---> \n Shuting down server gracefully: main___Shutdown_Servers ()")
                }
                os.Exit (1)
        } ()
        // ... }

        // Shutting down server gracefully. { ...
        errX := main___server_Info.Shutdown (context.Background ())
        if errX != nil {
                error_Message := fmt.Sprintf ("%s ---> \n Gracefully shutting down service server main___Shutdown_Servers ()", errX.Error ())
                RECORD_LOG (error_Message)
        }
        // ... }

        // Shutting down other components that should be explicitly shutdown. { ...
        SHUTDOWN_LOGGER ()
        // ... }

        // Finally halting the software.
        os.Exit (0)
}

var main___server_Info *http.Server // Information needed to run the software's server.
