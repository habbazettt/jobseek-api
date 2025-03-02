package config

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var CLD *cloudinary.Cloudinary

func SetupCloudinary() error {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return fmt.Errorf("gagal menginisialisasi Cloudinary: %v", err)
	}
	CLD = cld
	fmt.Println("Cloudinary berhasil diinisialisasi")
	return nil
}

func UploadImage(file multipart.File) (string, error) {
	uploadResult, err := CLD.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: "avatars",
	})
	if err != nil {
		return "", fmt.Errorf("gagal mengunggah gambar: %v", err)
	}
	return uploadResult.SecureURL, nil
}
