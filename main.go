package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"log"
)

func main() {
	http.HandleFunc("/start", startMotion)
	http.HandleFunc("/stop", stopMotion)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startMotion(w http.ResponseWriter, r *http.Request){
	cmd := exec.Command("sudo", "systemctl", "start", "motion")
	err := cmd.Run()
	if err != nil{
		http.Error(w, "Failed to start motion service", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Motion service started")
}

func stopMotion(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sudo", "systemctl", "stop", "motion")
	err := cmd.Run()
	if err != nil {
		http.Error(w, "Failed to stop motion service", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Motion service stopped")
}
