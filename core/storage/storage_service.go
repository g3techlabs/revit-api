package storage

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/g3techlabs/revit-api/config"
	customErrors "github.com/g3techlabs/revit-api/core/storage/errors"
	"github.com/g3techlabs/revit-api/response/generics"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/sirupsen/logrus"
)

type StorageService interface {
	PresignPutObjectURL(objectKey string, contentType string) (string, error)
	DoesObjectExist(objectKey string) error
	DeleteObject(objectKey string) error
}

type S3Service struct {
	Client               *s3.Client
	PresignClient        *s3.PresignClient
	Context              context.Context
	bucket               string
	presignPutExpiration int
	presginGetExpiration int
	log                  *logrus.Logger
}

func NewS3Service(client *s3.Client, presignClient *s3.PresignClient, c context.Context) StorageService {
	s3Client := &S3Service{
		Client:               client,
		PresignClient:        presignClient,
		Context:              c,
		bucket:               config.Get("AWS_BUCKET_NAME"),
		presignPutExpiration: config.GetIntVariable("PRESIGNED_PUT_URL_EXPIRATION"),
		presginGetExpiration: config.GetIntVariable("PRESIGNED_GET_URL_EXPIRATION"),
	}
	s3Client.initBucket(s3Client.bucket)
	return s3Client
}

func (s *S3Service) initBucket(bucket string) {
	_, err := s.Client.CreateBucket(s.Context, &s3.CreateBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		var bne *types.BucketAlreadyExists
		if errors.As(err, &bne) {
			utils.Log.Info("Bucket already exists. Skipping creation...")
		}
		utils.Log.Warnf("Unable to create the bucket: %v", err)
	}
	utils.Log.Infof("Created bucket %q.\n", bucket)
}

func (s *S3Service) PresignPutObjectURL(objectKey string, contentType string) (string, error) {
	req, err := s.PresignClient.PresignPutObject(s.Context, &s3.PutObjectInput{
		Bucket:      &s.bucket,
		Key:         &objectKey,
		ContentType: &contentType,
	}, s3.WithPresignExpires(time.Minute*time.Duration(s.presignPutExpiration)))
	if err != nil {
		return "", err
	}

	return req.URL, nil
}

func (s *S3Service) DoesObjectExist(objectKey string) error {
	_, err := s.Client.HeadObject(s.Context, &s3.HeadObjectInput{
		Bucket: &s.bucket,
		Key:    &objectKey,
	})

	if err == nil {
		return nil
	}

	var responseError *http.ResponseError
	if errors.As(err, &responseError) {
		switch responseError.HTTPStatusCode() {
		case 404:
			s.log.Errorf("Object not found in S3: %s", err.Error())
			return customErrors.ObjectNotFound()
		case 403:
			s.log.Errorf("Error of permission in HeadObject operation: %s", err.Error())
			return generics.InternalError()
		default:
			s.log.Errorf("Error checking if object exists in S3: %s", err.Error())
			return generics.InternalError()
		}
	}
	s.log.Errorf("Error checking if object exists in S3: %s", err.Error())
	return generics.InternalError()
}

func (s *S3Service) DeleteObject(objectKey string) error {
	_, err := s.Client.DeleteObject(s.Context, &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &objectKey,
	})
	if err == nil {
		return nil
	}

	var responseError *http.ResponseError
	if errors.As(err, &responseError) {
		switch responseError.HTTPStatusCode() {
		case 404:
			s.log.Errorf("Object not found in S3: %s", err.Error())
			return generics.InternalError()
		case 403:
			s.log.Errorf("Error of permission in HeadObject operation: %s", err.Error())
			return generics.InternalError()
		default:
			s.log.Errorf("Error checking if object exists in S3: %s", err.Error())
			return generics.InternalError()
		}
	}

	s.log.Errorf("Error deleting object in S3: %s", err.Error())
	return generics.InternalError()
}
