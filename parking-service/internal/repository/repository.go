package repository

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

type Repository interface {
	Save(data string) error
	Load() ([]string, error)
	Delete(plateNumber string) error
}

type repositoryImpl struct {
	dataFile string
	mu       sync.Mutex
}

func NewRepository(dataFile string) Repository {
	return &repositoryImpl{dataFile: dataFile}
}

func (r *repositoryImpl) Save(data string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.OpenFile(r.dataFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data + "\n")
	return err // Handle write errors here
}

func (r *repositoryImpl) Load() ([]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.Open(r.dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err // Handle scan errors here
	}

	return lines, nil
}

func (r *repositoryImpl) Delete(plateNumber string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	newFile, err := os.CreateTemp("", "repository-*.tmp")
	if err != nil {
		return err // Handle temporary file creation error
	}
	defer os.Remove(newFile.Name()) // Clean up temporary file on exit

	writer := bufio.NewWriter(newFile)
	defer writer.Flush() // Ensure data is written before closing

	file, err := os.Open(r.dataFile)
	if err != nil {
		return err // Handle original file opening error
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if fields[0] != plateNumber {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err // Handle write errors during copy
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err // Handle scan errors here
	}

	// Close the writer before reading from newFile
	if err := writer.Flush(); err != nil {
		return err
	}
	if err := newFile.Close(); err != nil {
		return err
	}

	// Read contents from newFile and write to r.dataFile
	data, err := os.ReadFile(newFile.Name())
	if err != nil {
		return err
	}
	if err := os.WriteFile(r.dataFile, data, 0644); err != nil {
		return err
	}

	return nil
}
