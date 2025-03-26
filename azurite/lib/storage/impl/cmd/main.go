package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"os"
	"strings"
)

const connectionString = "DefaultEndpointsProtocol=http;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;BlobEndpoint=http://127.0.0.1:10000/devstoreaccount1;QueueEndpoint=http://127.0.0.1:10001/devstoreaccount1;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;"

func main() {
	conStr := strings.TrimSpace(connectionString)

	client, err := azblob.NewClientFromConnectionString(conStr, nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ctx, f := context.WithCancel(ctx)
	defer f()

	serviceClient := client.ServiceClient()

	_, err = serviceClient.DeleteContainer(ctx, "container", nil)
	if err != nil {
		panic(err)
	}

	containerClient := serviceClient.NewContainerClient("container")
	_, err = containerClient.Create(ctx, nil)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("../../../../data/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	blobClient := containerClient.NewBlockBlobClient("blob.txt")

	_, err = blobClient.UploadBuffer(ctx, []byte("buffer"), nil)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 100)
	_, err = blobClient.DownloadBuffer(ctx, buf, nil)
	fmt.Println(string(buf))
}
