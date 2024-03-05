package services

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const dir = "uploads"
const defaultFile = "def.png"

type FileService struct {
	UploadsDir string
}

func NewFileService() *FileService {
	return &FileService{UploadsDir: dir}
}

func (fs *FileService) RetrieveFile(filename string) ([]byte, error) {
	filePath := filepath.Join(dir, filename)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (fs *FileService) SaveFile(header *multipart.FileHeader) (string, error) {
	if header == nil {
		return "", errors.New("header is nil")
	}
	newid := uuid.New().String()
	newfilename := newid + filepath.Ext(header.Filename)
	file, err := header.Open()
	if err != nil {
		return "", err
	}

	filebytes, err := io.ReadAll(file)

	if err != nil {

		return "", err
	}

	filePath := filepath.Join(dir, newfilename)
	err = os.WriteFile(filePath, filebytes, 0644)
	if err != nil {
		return "", err
	}
	return newfilename, nil
}
