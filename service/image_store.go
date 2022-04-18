package service

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"os"
	"sync"
)

// ImageStore is an interface to store laptop images
type ImageStore interface {
	// Save saves a new laptop image to the store
	Save(laptopID string, imageType string, imageData bytes.Buffer) (string, error)
}

type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string
}

// the images store on disk, and the image information store in memory
type DiskImageStore struct {
	mutex       sync.RWMutex
	imageFolder string
	images      map[string]*ImageInfo
}

func (store *DiskImageStore) Save(
	laptopID string,
	imageType string,
	imageData bytes.Buffer,
) (string, error) {
	imageId, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate uuid: %v", err)
	}

	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, imageId.String(), imageType)

	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %v", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %v", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.images[imageId.String()] = &ImageInfo{
		LaptopID: laptopID,
		Type:     imageType,
		Path:     imagePath,
	}

	return imageId.String(), nil
}

func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}
