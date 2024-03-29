package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

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
	Drive string `json:"drive"`
	Path string `json:"path"`
	FullPath string `json:"fullPath"`
	Copied bool `json:"copied"`
}

func (a *App) ListImages(source, target string) []Image {
  images := []Image{}

  // Walk source directory and add images to the array with copied=false
  err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    if !info.IsDir() && isImageFile(path) {
      var drive string
      if strings.Contains(path, ":") {
        drive = strings.Split(path, ":")[0]
      }
      var cleanPath string
      if drive != "" {
        cleanPath = strings.Replace(path, drive + ":", "", 1)
      } else {
        cleanPath = path
      }
      images = append(images, Image{
        Name:     info.Name(),
        Path:     cleanPath,
        Drive:    drive,
        FullPath: drive + ":" + cleanPath,
				Copied:   false,
      })
    }
    return nil
  })

  

  if err != nil {
    return nil
  }

  imageList := naturalSort(images)

  return imageList
}

func naturalSort(images []Image) []Image {
	re := regexp.MustCompile(`(\d+)`)
	returnImages := make([]Image, len(images))
	copy(returnImages, images)

	sort.Slice(returnImages, func(i, j int) bool {
			a, b := returnImages[i].Name, returnImages[j].Name
			aMatches, bMatches := re.FindAllString(a, -1), re.FindAllString(b, -1)
			for i := 0; i < len(aMatches) && i < len(bMatches); i++ {
					aNum, _ := strconv.Atoi(aMatches[i])
					bNum, _ := strconv.Atoi(bMatches[i])
					if aNum != bNum {
							return aNum < bNum
					}
					a = re.ReplaceAllString(a, "")
					b = re.ReplaceAllString(b, "")
			}
			return a < b
	})

	return returnImages
}

func SortImages(images []Image) []Image {
  // Sort the images by name, treating numeric sequences as numbers
  sort.Slice(images, func(i, j int) bool {
    name1 := images[i].Name
    name2 := images[j].Name
    for len(name1) > 0 && len(name2) > 0 {
      rune1, size1 := utf8.DecodeRuneInString(name1)
      rune2, size2 := utf8.DecodeRuneInString(name2)
      if unicode.IsDigit(rune1) && unicode.IsDigit(rune2) {
        // Compare numeric sequences as numbers
        num1, _ := strconv.Atoi(strings.TrimLeft(name1[:size1], "0"))
        num2, _ := strconv.Atoi(strings.TrimLeft(name2[:size2], "0"))
        if num1 != num2 {
          return num1 < num2
        }
        // If numeric sequences are equal, continue comparing strings
        name1 = name1[size1:]
        name2 = name2[size2:]
      } else if rune1 != rune2 {
        // Compare non-numeric runes
        return rune1 < rune2
      } else {
        // If runes are equal, continue comparing strings
        name1 = name1[size1:]
        name2 = name2[size2:]
      }
    }
    // If one name is a prefix of the other, the shorter name comes first
    return len(name1) < len(name2)
  })

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

func (a *App) MoveFile(srcPath, destPath string) error {
	fmt.Println("MoveFile", srcPath, destPath)

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
		fmt.Println("MoveFile", err)
		return err
	}

	// Delete the source file
	if err := os.Remove(srcPath); err != nil {
		fmt.Println("MoveFile", err)
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

func (a *App) SavePage(path string, page int) {
	// Load existing data from file
	data, err := loadDataFromFile("data.json")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	// Add new key-value pair
	data[path] = page

	// Save updated data to file
	err = saveDataToFile("data.json", data)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	fmt.Printf("Saved page %d for path %s\n", page, path)
}

func (a *App) GetPage(path string) (int, error) {
	// Load data from file
	data, err := loadDataFromFile("data.json")
	if err != nil {
		return 0, err
	}

	// Lookup page number by path
	page, ok := data[path]
	if !ok {
		return 0, fmt.Errorf("page not found for path %s", path)
	}

	return page, nil
}

func loadDataFromFile(fileName string) (map[string]int, error) {
	data := make(map[string]int)

	// Open file for reading
	file, err := os.Open(fileName)
	if err != nil {
		// Return empty data if file doesn't exist yet
		if os.IsNotExist(err) {
			return data, nil
		}
		return nil, err
	}
	defer file.Close()

	// Read file contents
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func saveDataToFile(fileName string, data map[string]int) error {
	// Marshal data to JSON
	contents, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write JSON data to file
	err = ioutil.WriteFile(fileName, contents, 0644)
	if err != nil {
		return err
	}

	return nil
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