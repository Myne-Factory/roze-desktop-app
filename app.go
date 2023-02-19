package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DiagFolderPath(title string) string {

	result, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: title,
	})

	if err != nil {
		return ""
	}

	return result
}

type Image struct{
	Name string `json:"name"`
	Path string `json:"path"`
}

func (a *App) ListImages(path string) []Image {
	images := []Image{}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isImageFile(path) {
			images = append(images, Image{
				Name: info.Name(),
				Path: path,
			})
		}
		return nil
	})

	if err != nil {
		return nil
	}

	return images
}

func (a *App) CopyFile(srcPath, destPath string) error {
	fmt.Println("CopyFile", srcPath, destPath)
	// Open the source file for reading
	srcFile, err := os.Open(srcPath)
	if err != nil {
			return err
	}
	defer srcFile.Close()

	// Check if the destination directory exists
	destDir := filepath.Dir(destPath)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
			// Create the destination directory
			if err := os.MkdirAll(destDir, 0755); err != nil {
					return err
			}
	}

	// Create the destination file for writing
	destFile, err := os.Create(destPath)
	if err != nil {
		
		return err
	}
	defer destFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("CopyFile", err)
		return err
	}

	return nil
}

func (a *App) DeleteFile(path string) error {
	fmt.Println("DeleteFile", path)
	err := os.Remove(path)
	if err != nil {
			return err
	}
	return nil
}

func (a *App) FileExists(dir, fileName string) (bool, error) {
	// Construct the path to the file
	filePath := filepath.Join(dir, fileName)

	// Get information about the file
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
			// The file doesn't exist
			return false, nil
	} else if err != nil {
			// An error occurred while getting file info
			return false, err
	} else {
			// The file exists
			return true, nil
	}
}

func (a *App) GetBase64Data(path string) (string) {
	// Open the file for reading
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	// Create a buffer to hold the encoded data
	buf := bytes.NewBuffer(nil)

	// Create an encoder for writing Base64 data to the buffer
	encoder := base64.NewEncoder(base64.StdEncoding, buf)

	// Read and encode the file in chunks
	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return ""
		}
		if n == 0 {
			break
		}

		_, err = encoder.Write(buffer[:n])
		if err != nil {
			return ""
		}
	}

	// Close the encoder to flush any remaining data to the buffer
	err = encoder.Close()
	if err != nil {
		return ""
	}

	// Return the encoded data as a string
	return buf.String()
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".webp" || ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif"
}