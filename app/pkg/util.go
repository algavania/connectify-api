package pkg

import (
	"mime/multipart"
	"net/http"
	"strings"
)

func CheckType(fileHeader *multipart.FileHeader) string {
	file, err := fileHeader.Open()
	if err != nil {
		return ""
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return ""
	}

	mimeType := http.DetectContentType(buffer)
	return mimeType
}

func IsImageFile(fileHeader *multipart.FileHeader) bool {

	mimeType := CheckType(fileHeader)

	// Check if the MIME type corresponds to an image format
	return strings.HasPrefix(mimeType, "image/")
}

func IsVideoFile(fileHeader *multipart.FileHeader) bool {

	mimeType := CheckType(fileHeader)

	// Check if the MIME type corresponds to an image format
	return strings.HasPrefix(mimeType, "image/")
}
