package main

import (
	"fmt"
	"github.com/nicolaifsf/go-speak"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
	speech.SetWitKey("/**Your Wit API Key here**/") //Wit API Key MUST be set before calling any other Wit.AI functions
	speech.SendWitVoice("test.wav")
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
}

func browseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
}
