package helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var datadir = ""

// getHomeDir returns path to user directory
// windows: C:\users\cyjan
// linux:   /home/cyjan
func getHomeDir() string {
	var home = ""
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	} else {
		home = os.Getenv("HOME")
	}
	return home
}

// SetDataDir - Set path for storing data
func SetDataDir(newdatadir string) {
	datadir = newdatadir
}

// GetDataDir - Get string of path that should be used to store data
func GetDataDir() string {
	if datadir != "" {
		return datadir
	}
	var dir = ""
	if runtime.GOOS == "windows" {
		dir = getHomeDir() + "\\LibJWgo"
	} else {
		dir = getHomeDir() + "/LibJWgo"
	}
	return dir
}

// Mkdir - make a directory, recursive
func Mkdir(dirName string) error {
	err := os.MkdirAll(dirName, 0770)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

// PrintMemUsage - print memory usage to console
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\r", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Unzip - read file and unzip it to somewhere.
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0770)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, 0770)
		} else {
			os.MkdirAll(filepath.Dir(path), 0770)
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0770)
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					fmt.Println(err)
					return
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// GetFileContentType - returns file type
func GetFileContentType(out *os.File) string {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return ""
	}
	return GetBufferType(buffer)
}

// GetBufferType - Get type of buffer
func GetBufferType(buffer []byte) string {
	// Use the net/http package's handy DetectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType
}
