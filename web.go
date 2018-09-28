package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"os/exec"
)

func main(){
	namespace := mux.NewRouter().StrictSlash(true)
	namespace.HandleFunc("/", isAlive)
	namespace.HandleFunc("/action", action)
	namespace.HandleFunc("/action/reboot", aReboot)
	namespace.HandleFunc("/action/shutdown", aShutdown)
	namespace.HandleFunc("/action/update", aUpdate)

	log.Fatal(http.ListenAndServe(":6660", namespace))
}

//Namespace handlers
func isAlive (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "imAliveAndServing")
}

func action (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "success, %q\n", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "available endpoints:\n")
	fmt.Fprintf(w, "reboot\n")
	fmt.Fprintf(w, "shutdown\n")
	fmt.Fprintf(w, "update\n")
}

func aReboot (w http.ResponseWriter, r *http.Request) {
	cmdOut := command(string("shutdown"), []string{"-r" , "+1"})
	fmt.Fprintf(w, "initiated reboot, %q\n", html.EscapeString(r.URL.Path))

	if cmdOut == nil{
		fmt.Fprintf(w, "success\n")
	} else {
		fmt.Fprintf(w, "failed with %v\n", cmdOut)
	}
}

func aShutdown (w http.ResponseWriter, r *http.Request){
	cmdOut := command(string("shutdown"), []string{"-h" , "+1"})
	fmt.Fprintf(w, "initiated shutdown\n")

	if cmdOut == nil {
		fmt.Fprintf(w, "success\n")
	}else {
		fmt.Fprintf(w, "failed with %v\n", cmdOut)
	}
}

func aUpdate (w http.ResponseWriter, r *http.Request){
	cmdOut := command(string("apt-get"), []string{"upgrade" , "-y"})
	fmt.Fprintf(w, "initiated update\n")

	if cmdOut == nil{
		fmt.Fprintf(w, "success\n")
	}else {
		fmt.Fprintf(w, "failed with %v\n", cmdOut)
	}
}

//Commands
func command(cmdName string, args []string) error {
	cmd := exec.Command(cmdName, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Start()

	if err != nil {
		return fmt.Errorf("failed starting %s %v: %s: %v", cmdName, args, stderr.String(), err)
	}

	return nil
}