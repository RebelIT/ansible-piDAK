package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
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


	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	//	fmt.Fprintf(w, "Success, %q", html.EscapeString(r.URL.Path))
	//})
	log.Fatal(http.ListenAndServe(":8081",nil))
}

//Namespace handlers
func isAlive (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "imAliveAndServing")
}

func action (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "available endpoints:")
	fmt.Fprintln(w, "reboot")
	fmt.Fprintln(w, "shutdown")
	fmt.Fprintln(w, "update")
}

func aReboot (w http.ResponseWriter, r *http.Request) {
	cmdOut := command(string("shutdown"), []string{"-r" , "+1"})
	fmt.Fprintln(w, "initiated reboot")

	if cmdOut == nil{
		fmt.Fprintln(w, "success")
	} else {
		fmt.Fprintln(w, "failed with %v", cmdOut)
	}
}

func aShutdown (w http.ResponseWriter, r *http.Request){
	cmdOut := command(string("shutdown"), []string{"-h" , "+1"})
	fmt.Fprintln(w, "initiated shutdown")

	if cmdOut == nil {
		fmt.Fprintln(w, "success")
	}else {
		fmt.Fprintln(w, "failed with %v", cmdOut)
	}
}

func aUpdate (w http.ResponseWriter, r *http.Request){
	cmdOut := command(string("apt-get"), []string{"upgrade" , "-y"})
	fmt.Fprintln(w, "initiated update")

	if cmdOut == nil{
		fmt.Fprintln(w, "success")
	}else {
		fmt.Fprintln(w, "failed with %v", cmdOut)
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
	//err = cmd.Wait()
	//if err != nil {
	//	return fmt.Errorf("failed waiting %s %v: %s: %v", cmdName, args, stderr.String(), err)
	//}
	return nil
}