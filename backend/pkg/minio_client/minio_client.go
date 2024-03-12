package minio_client

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kavkaco/Kavka-Core/utils/random"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tahadostifam/MusicStreamingApp/config"
)

const RandomFileNameLength = 25

type (
	MinioClient struct{ minioClient *minio.Client }
	File        struct {
		Name string
		Size int64
	}
)

func NewMinioClient(endpoint, accessKeyID, secretAccessKey string) *MinioClient {
	creds := credentials.NewStaticV4(accessKeyID, secretAccessKey, "")
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: creds,
	})
	if err != nil {
		panic(err)
	}

	return &MinioClient{minioClient}
}

func (s *MinioClient) UploadFile(bucketName string, filePath string) (*File, error) {
	// Collect objectName, contentType and filePath
	fileInfo, statErr := os.Stat(filePath)
	if statErr != nil {
		return nil, statErr
	}

	objectName := random.GenerateRandomFileName(RandomFileNameLength)
	contentType := filepath.Ext(filePath)

	// Upload the file
	_, err := s.minioClient.FPutObject(context.Background(), bucketName,
		objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return nil, err
	}

	return &File{
		Name: objectName,
		Size: fileInfo.Size(),
	}, err
}

func (s *MinioClient) DeleteFile(bucketName string, objectName string) error {
	// Delete the file
	opts := minio.RemoveObjectOptions{GovernanceBypass: true}

	err := s.minioClient.RemoveObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		return err
	}

	return nil
}

// DownloadFile function gets the file from minio-bucket and saves it in the directory you give to it.
func (s *MinioClient) DownloadFile(saveDirectory, bucketName, objectName string) (localFileInTmp string, err error) {
	opts := minio.GetObjectOptions{}
	obj, err := s.minioClient.GetObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		return "", err
	}
	defer obj.Close()

	randomFileName := random.GenerateRandomFileName(RandomFileNameLength)
	filePath := fmt.Sprintf("%s%s/%s", config.ProjectRootPath, saveDirectory, randomFileName)
	localFile, createErr := os.Create(filePath)
	if createErr != nil {
		return "", createErr
	}
	defer localFile.Close()

	stat, statErr := obj.Stat()
	if statErr != nil {
		return "", statErr
	}

	if _, copyErr := io.CopyN(localFile, obj, stat.Size); copyErr != nil {
		return "", copyErr
	}

	return filePath, nil
}
