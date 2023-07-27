package networking

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadFileBytes(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
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

func ZipFolder(folderPath string, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(folderPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create the relative path inside the zip file
		relPath, err := filepath.Rel(folderPath, filePath)
		if err != nil {
			return err
		}

		// Add file or folder to the zip archive
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If it's a regular file, write its contents to the zip
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

	return err
}

func UnzipFolder(zipFilePath, targetDir string) error {
	zipReader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		// Construct the full file path based on the target directory
		targetPath := filepath.Join(targetDir, file.Name)

		// If the file is a directory, create it
		if file.FileInfo().IsDir() {
			os.MkdirAll(targetPath, file.Mode())
			continue
		}

		// Otherwise, create the file and write its contents
		err = os.MkdirAll(filepath.Dir(targetPath), 0755)
		if err != nil {
			return err
		}

		outFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		inFile, err := file.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}
	}

	return nil
}
