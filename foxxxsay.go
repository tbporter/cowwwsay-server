// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
    "net/http"
    "log"
    "math/rand"
)

var whatthefoxsays = [...]string{
"Ring-ding-ding-ding-dingeringeding!",
"Gering-ding-ding-ding-dingeringeding!",
"Wa-pa-pa-pa-pa-pa-pow!",
"Hatee-hatee-hatee-ho!",
"Joff-tchoff-tchoffo-tchoffo-tchoff!",
"Tchoff-tchoff-tchoffo-tchoffo-tchoff!",
"Jacha-chacha-chacha-chow!",
"Chacha-chacha-chacha-chow!",
"Fraka-kaka-kaka-kaka-kow!",
"A-hee-ahee ha-hee!",
"A-oo-oo-oo-ooo!",
"Woo-oo-oo-ooo!",
"Wa-wa-way-do Wub-wid-bid-dum-way-do Wa-wa-way-do",
"Bay-budabud-dum-bam",
"Mama-dum-day-do",
"Abay-ba-da bum-bum bay-do",
}

const foxcount = int32(len(whatthefoxsays))


func cowHandler(w http.ResponseWriter, r *http.Request){
    fmt.Println("handling the foxes")
    
    if _, e := w.Write(([]byte)(whatthefoxsays[rand.Int31n(foxcount)])); e != nil {
        log.Fatal(e)
    }
}

func main() {
    fmt.Println("Started")
    http.HandleFunc("/say", cowHandler)//makeHandler(cowHandler))
    http.ListenAndServe(":80", nil)
}
