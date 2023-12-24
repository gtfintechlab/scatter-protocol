package networking

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadFileBytes(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return []byte{}
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var data []byte
	for {
		chunk := make([]byte, 1024)
		n, err := reader.Read(chunk)
		if err != nil {
			break
		}
		data = append(data, chunk[:n]...)
	}

	return data
}

func WriteBytesToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()

	// Write the data to the file.
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err)
	}

	return nil
}

func ZipFolder(folderPath string) (*bytes.Buffer, error) {
	zipBuffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(zipBuffer)
	defer zipWriter.Close()

	err := filepath.Walk(folderPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Modify the header's name to be relative to the folder being zipped
		relPath, _ := filepath.Rel(folderPath, filePath)
		header.Name = relPath

		if info.IsDir() {
			header.Name += "/"
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return zipBuffer, err
}
func UnzipFolder(zipBytes []byte, destPath string) error {
	zipReader, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		fullPath := filepath.Join(destPath, file.Name)

		if file.FileInfo().IsDir() {
			// Create directories if they don't exist
			os.MkdirAll(fullPath, file.Mode())
			continue
		}

		// Create the file and set its permissions
		writer, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer writer.Close()

		// Open the zipped file
		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		// Copy the zipped file's content to the destination file
		_, err = io.Copy(writer, zippedFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func FindFilePathWithExtension(rootDir string, extension string) (string, error) {
	var pthFilePath string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			pthFilePath = path
			return filepath.SkipDir // To skip subdirectories once the file is found
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return pthFilePath, nil
}
