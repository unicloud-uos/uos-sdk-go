package lib

import (
	"os"
)

func (sc *S3Client) GenTestObjectUrl(bucketName, objectKey string) string {
	return "http://" + sc.Client.Config.Endpoint + string(os.PathSeparator) + bucketName + string(os.PathSeparator) + objectKey
}

// Generate 5M part data
func GenMinimalPart() []byte {
	return RandBytes(5 << 20)
}
