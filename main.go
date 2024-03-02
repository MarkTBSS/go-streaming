package main

import (
	"fmt"
	"net/http"
	"time"
)

func streamNumbers(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/event-stream
	w.Header().Set("Content-Type", "text/event-stream")

	// Start streaming numbers infinitely
	for i := 1; i < 20; i++ {
		// Format the number as a string and write it to the response writer
		fmt.Fprintf(w, "data: %d\n\n", i)
		w.(http.Flusher).Flush() // Flush the response writer to ensure data is sent immediately

		// Simulate some processing time
		// This is just to demonstrate continuous streaming, you can remove it if not needed
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Register the handler for the /stream endpoint
	http.HandleFunc("/stream", streamNumbers)

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
