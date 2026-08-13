package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tawesoft/golib/v2/dialog"
)

// Source - https://stackoverflow.com/a/33853856
// Posted by Pablo Jomer, modified by community. See post 'Timeline' for change history
// Retrieved 2026-08-13, License - CC BY-SA 4.0

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	dialog.Alert("Project set up!")
}
