package storage

import (
	"context"
	"io"
)

type Storage interface {
	// BucketExists will check for the existence of bucket named bucketName.
	BucketExists(ctx context.Context, bucketName string) (bool, error)

	// Delete will delete the file from the storage.
	DeleteObject(ctx context.Context, bucketName, objectName string) error

	// ObjectExists will check for the existence of object named objectName in a bucket named bucketName.
	ObjectExists(ctx context.Context, bucketName, objectName string) (bool, error)

	// Load will load the blob format object from the database.
	LoadObject(ctx context.Context, bucketName, objectName string) ([]byte, error)

	// Put will store or update already existence object named objectName in a bucket named bucketName.
	StoreObject(ctx context.Context, bucketName, objectName string, reader io.Reader, size int64) error
}
