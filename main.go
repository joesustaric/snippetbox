package main

import (
  "fmt"
  "log"
  "net/http"
  "strconv"
)

// Home Handler. Writes a byte slice as the response body
func home(w http.ResponseWriter, r *http.Request) {

  // Ensuer we 404 for any other route not matching "/"
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  w.Write([]byte("Hello from Joes Snippet Box"))
}

// Show snippet handler
func showSnippet(w http.ResponseWriter, r *http.Request) {

  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
// Create snippet handler
func createSnippet(w http.ResponseWriter, r *http.Request) {

  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    // Use the http.Error() function to send a 405 status code and "Method Not
    // Allowed" string as the response body.
    http.Error(w, "Method Not Allowed", 405)
    return
  }

  w.Write([]byte("Create a new snippet..."))
}

func main() {
  // Register handler functions and corresponding URL patterns with the servemux
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  // Use the http.ListenAndServe() function to start a new web server. We pass in
  // two parameters: the TCP network address to listen on (in this case ":4000")
  // and the servemux we just created. If http.ListenAndServe() returns an error
  // we use the log.Fatal() function to log the error message and exit.
  log.Println("Starting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
