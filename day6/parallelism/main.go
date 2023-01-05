package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

var wg sync.WaitGroup
var threadProfile = pprof.Lookup("threadCreate")

func BookTicket(wg *sync.WaitGroup) {

	defer wg.Done()
	// Open file for reading
	file, err := os.Open("ticket.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read up to len(b) bytes from the File
	// Zero bytes written means end of file
	// End of file returns error type io.EOF
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)

}

func ListenMusic(wg *sync.WaitGroup) {

	defer wg.Done()
	// Create output file
	newFile, err := os.Create("golang.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// HTTP GET request devdungeon.com
	url := "https://golang.org/doc/"
	response, err := http.Get(url)
	defer response.Body.Close()

	// Write bytes from HTTP response to file.
	// response.Body satisfies the reader interface.
	// newFile satisfies the writer interface.
	// That allows us to use io.Copy which accepts
	// any type that implements reader and writer interface
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}

func SendMail(wg *sync.WaitGroup) {

	defer wg.Done()
	// Create output file
	newFile, err := os.Create("golang.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// HTTP GET request devdungeon.com
	url := "https://golang.org/doc/"
	response, err := http.Get(url)
	defer response.Body.Close()

	// Write bytes from HTTP response to file.
	// response.Body satisfies the reader interface.
	// newFile satisfies the writer interface.
	// That allows us to use io.Copy which accepts
	// any type that implements reader and writer interface
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}

func OrderSuite(wg *sync.WaitGroup) {
	defer wg.Done()
	// Open file for reading
	file, err := os.Open("ticket.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read up to len(b) bytes from the File
	// Zero bytes written means end of file
	// End of file returns error type io.EOF
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of Processors in use %d", runtime.NumCPU())
}

var listOfTasks = []func(wg *sync.WaitGroup){
	BookTicket, ListenMusic, SendMail, OrderSuite,
}

func main() {
	fmt.Println("Before Calling Routines", threadProfile.Count())
	wg.Add(len(listOfTasks))

	for _, task := range listOfTasks {
		go task(&wg)
	}

	wg.Wait()

	fmt.Println("Returned to main and completed all the tasks", threadProfile.Count())

}
