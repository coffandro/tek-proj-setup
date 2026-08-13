package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

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
	err := downloadFile(".gitignore",
		"https://github.com/coffandro/tek-proj-setup/releases/download/files/default.gitignore")
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	err = downloadFile(".gitattributes ",
		"https://github.com/coffandro/tek-proj-setup/releases/download/files/default.gitattributes")
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	dialog.Alert("Project files downloaded! [1/5]")

	err = downloadFile("git_installer.exe ",
		"https://github.com/git-for-windows/git/releases/download/v2.55.0.windows.4/Git-2.55.0.4-64-bit.exe")
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	err = downloadFile("git_lfs_installer.exe ",
		"https://github.com/git-lfs/git-lfs/releases/download/v3.7.1/git-lfs-windows-v3.7.1.exe")
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	dialog.Alert("Git temp files downloaded! [2/5]")

	dialog.Alert("We are about to download and install git, please follow the instructions in the installer(s), leave everything as default.")

	err = exec.Command("./git_installer.exe").Run()
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	err = exec.Command("./git_lfs_installer.exe").Run()
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	dialog.Alert("Git installed! [3/5]")

	err = os.Remove("git_installer.exe")
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	err = os.Remove("git_lfs_installer.exe")
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	dialog.Alert("Cleaned up git temp files! [4/5]")

	err = exec.Command("git.exe", "init").Run()
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	err = exec.Command("git.exe", "lfs", "install").Run()
	if err != nil {
		dialog.Error(err.Error())
		panic(err.Error())
	}

	dialog.Alert("Project set up! [5/5]")
}
