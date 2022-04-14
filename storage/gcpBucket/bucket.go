package gcpBucket

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func init() {
    log.SetPrefix("gcpBucket: ")
    log.SetFlags(0)
}

func Write(filename string, content string) {
    ctx, storageClient := client()
    defer storageClient.Close()

    bkt := storageClient.Bucket(os.Getenv("GCP_BUCKET_NAME"))
    obj := bkt.Object(filename)

    w := obj.NewWriter(ctx)
    if _, err := fmt.Fprint(w, string(content)); err != nil {
        log.Fatal(err)
    }
    if err := w.Close(); err != nil {
        log.Fatal(err)
    }
}

func Read(filename string) []byte {
    ctx, storageClient := client()
    defer storageClient.Close()

    bkt := storageClient.Bucket(os.Getenv("GCP_BUCKET_NAME"))
    obj := bkt.Object(filename)

    r, err := obj.NewReader(ctx)
    if err != nil {
        log.Fatal(err)
    }
    data, err := io.ReadAll(r);
    if err != nil {
        log.Fatal(err)
    }

    return data
}

func Exists(filename string) bool {
    ctx, storageClient := client()
    defer storageClient.Close()

    _, err := storageClient.Bucket(os.Getenv("GCP_BUCKET_NAME")).Object(filename).Attrs(ctx)
    if err == storage.ErrObjectNotExist {
        return false
    }
    if err != nil {
        return false
    }

    return true
}

func client() (context.Context, *storage.Client) {
    ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

    return ctx, storageClient
}
