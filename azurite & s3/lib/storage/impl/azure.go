package impl

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/errors"
	"io"
	"strings"
)

const defaultBlobSize = 1024 * 1024 * 1

type AzureClient struct {
	*service.Client
}

func (a AzureClient) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	containerClient := a.NewContainerClient(bucketName)

	_, err := containerClient.Create(ctx, nil)
	if err != nil {
		if strings.Contains(err.Error(), "409 The specified container already exists") {
			return true, nil
		}

		return false, err
	}

	_, err = containerClient.Delete(ctx, nil)
	if err != nil {
		return false, err
	}

	return false, nil
}

func (a AzureClient) DeleteObject(ctx context.Context, bucketName, objectName string) error {
	exists, err := a.BucketExists(ctx, bucketName)
	if !exists {
		return errors.New(fmt.Sprintf("bucket %s does not exist", bucketName))
	}

	exists, err = a.ObjectExists(ctx, bucketName, objectName)
	if !exists {
		return errors.New(fmt.Sprintf("object %s does not exist in the %s", objectName, bucketName))
	}

	containerClient := a.NewContainerClient(bucketName)
	blobClient := containerClient.NewBlockBlobClient(objectName)

	_, err = blobClient.Delete(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func (a AzureClient) ObjectExists(ctx context.Context, bucketName, objectName string) (bool, error) {
	containerClient := a.NewContainerClient(bucketName)
	blobClient := containerClient.NewBlockBlobClient(objectName)

	streamResp, err := blobClient.DownloadStream(ctx, nil)
	if err != nil {
		return false, err
	}

	_, err = streamResp.Body.Read(nil)
	if err != nil {
		if err == io.EOF {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (a AzureClient) LoadObject(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	exists, err := a.BucketExists(ctx, bucketName)
	if !exists {
		return nil, errors.New(fmt.Sprintf("bucket %s does not exist", bucketName))
	}

	exists, err = a.ObjectExists(ctx, bucketName, objectName)
	if !exists {
		return nil, errors.New(fmt.Sprintf("object %s does not exist in the %s", objectName, bucketName))
	}

	containerClient := a.NewContainerClient(bucketName)
	blobClient := containerClient.NewBlockBlobClient(objectName)

	respStream, err := blobClient.DownloadStream(ctx, nil)
	if err != nil {
		return nil, err
	}

	content := make([]byte, *(respStream.ContentLength))
	_, err = respStream.Body.Read(content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (a AzureClient) StoreObject(ctx context.Context, bucketName, objectName string, reader io.Reader, size int64) error {
	containerClient := a.NewContainerClient(bucketName)
	exists, err := a.BucketExists(ctx, bucketName)
	if !exists {
		_, err = containerClient.Create(ctx, nil)
		if err != nil {
			return err
		}
	}

	blobClient := containerClient.NewBlockBlobClient(objectName)

	if size <= 0 {
		size = defaultBlobSize
	}

	_, err = blobClient.UploadStream(ctx, reader, &azblob.UploadStreamOptions{
		BlockSize: size,
	})

	return err
}
