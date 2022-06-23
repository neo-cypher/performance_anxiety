package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"storj.io/uplink"
)

const (
	myAccessGrant = "1T4nEymKBctvv2pBXbqhiAqMkirmRZbpQsAsKzNVvg3qFEWsqxj4Y1WHMYbWd4GXAfi16DixF9r1QXKPLkSNSqGSr93g5XJ97Sqznv3PMmNiE6bRbhfSsxcVc4M9Ms18BrzmiiU6ojQ2icTsfdkrF8m5twcqvJegViRRiKef25E3D8oJ2qE1bM7NLCViMj8zBx2sC1zZC5BsdaXy1kcAZK3ULy5Ne8LnrYq74w7oTc8npPemDyUfjBF1CL17v2QGvmaor5oYqHFMtdxAbaz8HwM"
	myBucket      = "performance"
	myObjectKey   = "foo/bar/baz"
	myData        = "one fish two fish red fish blue fish"
)

// UploadAndDownloadData uploads the data to objectKey in
// bucketName, using accessGrant.
func UploadAndDownloadData(ctx context.Context,
	accessGrant, bucketName, objectKey string,
	data []byte) error {

	// Parse the Access Grant.
	access, err := uplink.ParseAccess(accessGrant)
	if err != nil {
		return fmt.Errorf("could not parse access grant: %v", err)
	}

	// Open up the Project we will be working with.
	project, err := uplink.OpenProject(ctx, access)
	if err != nil {
		return fmt.Errorf("could not open project: %v", err)
	}
	defer project.Close()

	// Ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("could not ensure bucket: %v", err)
	}

	// Intitiate the upload of our Object to the specified bucket and key.
	upload, err := project.UploadObject(ctx, bucketName, objectKey, nil)
	if err != nil {
		return fmt.Errorf("could not initiate upload: %v", err)
	}

	// Copy the data to the upload.
	buf := bytes.NewBuffer(data)
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		return fmt.Errorf("could not upload data: %v", err)
	}

	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		return fmt.Errorf("could not commit uploaded object: %v", err)
	}

	// Initiate a download of the same object again
	download, err := project.DownloadObject(ctx, bucketName, objectKey, nil)
	if err != nil {
		return fmt.Errorf("could not open object: %v", err)
	}
	defer download.Close()

	// Read everything from the download stream
	receivedContents, err := ioutil.ReadAll(download)
	if err != nil {
		return fmt.Errorf("could not read data: %v", err)
	}

	// Check that the downloaded data is the same as the uploaded data.
	if !bytes.Equal(receivedContents, data) {
		return fmt.Errorf("got different object back: %q != %q", data, receivedContents)
	}

	return nil
}

func main() {
	err := UploadAndDownloadData(context.Background(),
		myAccessGrant, myBucket, myObjectKey, []byte(myData))
	if err != nil {
		log.Fatalln("error:", err)
	}

	fmt.Println("success!")
}
