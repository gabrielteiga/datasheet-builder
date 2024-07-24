package services

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileService struct {
	Path string
}

type Schema struct {
	Table string    `json: "table"`
	Rows  []Message `json: "rows"`
}

type Message struct {
	Nome     string `json: "Nome"`
	Email    string `json: "Email"`
	Telefone string `json: "Telefone"`
}

func (f *FileService) GetFiles(path string) []string {
	filesFromDir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading files from directory: ", err.Error())
	}

	var files []string
	for _, file := range filesFromDir {
		files = append(files, file.Name())
	}

	return files
}

func (f *FileService) FormatFiles(path string) {
	var m Schema

	byteFile, err := f.readJson(path)
	if err == nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	err = json.Unmarshal(byteFile, &m)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		os.Exit(1)
	}

	// only for debugging purposes

}

func (f *FileService) readJson(path string) ([]byte, error) {
	fmt.Println(path)
	rawFile, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return rawFile, nil
}
