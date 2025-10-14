package utils

import (
	"strings"

	"github.com/g3techlabs/revit-api/config"
)

var cloudFrontUrl string = config.Get("AWS_CLOUDFRONT_URL")

func MountCloudFrontUrl(objectKey string) *string {
	if !strings.HasSuffix(cloudFrontUrl, "/") {
		cloudFrontUrl += "/"
	}
	mountedUrl := cloudFrontUrl + objectKey
	return &mountedUrl
}
