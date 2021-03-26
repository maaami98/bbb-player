package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

var epoch = time.Unix(0, 0).Format(time.RFC1123)

var noCacheHeaders = map[string]string{
    "Expires":         epoch,
    "Cache-Control":   "no-cache, private, max-age=0",
    "Pragma":          "no-cache",
    "X-Accel-Expires": "0",
}

var etagHeaders = []string{
    "ETag",
    "If-Modified-Since",
    "If-Match",
    "If-None-Match",
    "If-Range",
    "If-Unmodified-Since",
}

func NoCache(h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        // Delete any ETag headers that may have been set
        for _, v := range etagHeaders {
            if r.Header.Get(v) != "" {
                r.Header.Del(v)
            }
        }

        // Set our NoCache headers
        for k, v := range noCacheHeaders {
            w.Header().Set(k, v)
        }

        h.ServeHTTP(w, r)
    }

    return http.HandlerFunc(fn)
}
func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/public", 301)
}

func main() {
	//fileServer := http.FileServer(http.Dir("./public/")) // New code
	                   // New code
	http.Handle("/public/",NoCache(http.StripPrefix("/public", http.FileServer(http.Dir("./public/")))   ) )
	http.HandleFunc("/", redirect  ) 
	fmt.Println("---------------------------")
	fmt.Println("Dokuz Eylul University    ")
	fmt.Println("BigBlueButton Recording Viewer")
	fmt.Println("Author: github.com/maaami98")

	fmt.Println("---------------------------")
	fmt.Println("In your modern web browser open:")
	fmt.Println("http://localhost:8080")
	fmt.Println("Press CTRL+C when done.")
	fmt.Println("----------------------")
	openbrowser("http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
