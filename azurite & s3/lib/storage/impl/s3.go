package impl

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/minio/minio-go/v7"
	"io"
)

type S3Client struct {
	client *minio.Client
}

func (c S3Client) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	trace, err := c.client.BucketExists(ctx, bucketName)
	if err != nil {
		return false, err
	}

	if !trace {
		return false, nil
	}

	return true, nil
}

func (c S3Client) DeleteObject(ctx context.Context, bucketName, objectName string) error {
	exists, err := c.ObjectExists(ctx, bucketName, objectName)
	if err != nil || !exists {
		return err
	}

	return c.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
}

func (c S3Client) ObjectExists(ctx context.Context, bucketName, objectName string) (bool, error) {
	info, err := c.client.GetObjectACL(ctx, bucketName, objectName)
	if err != nil || info == nil {
		return false, err
	}

	return true, nil
}

func (c S3Client) LoadObject(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	info, err := c.client.GetObjectACL(ctx, bucketName, objectName)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = gob.NewEncoder(&buf).Encode(*info)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c S3Client) StoreObject(ctx context.Context, bucketName, objectName string, reader io.Reader, size int64) error {
	_, err := c.client.PutObject(ctx, bucketName, objectName, reader, size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
