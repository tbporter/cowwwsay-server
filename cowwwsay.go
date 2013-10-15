package main

import (
    "io"
    "os/exec"
    "fmt"
    "net/http"
    "log"
)

func cowHandler(w http.ResponseWriter, r *http.Request){
    fmt.Println("handling the cows")

    sayText := r.FormValue("say")
    cmd := exec.Command("cowsay", sayText)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    
    io.Copy(w,stdout)

    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/say", cowHandler)//makeHandler(cowHandler))
    http.ListenAndServe(":80", nil)
}
