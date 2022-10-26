package s3_utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	aws_aws "github.com/aws/aws-sdk-go/aws"
	aws_credentials "github.com/aws/aws-sdk-go/aws/credentials"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	aws_s3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	aws_s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cheggaaa/pb/v3"
	"github.com/sikalabs/slu/utils/vault_s3_utils"
)

func DeleteBucketWithObjects(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
) error {
	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			access_key,
			secret_key,
			"",
		),
	}
	if region != "" {
		awsConfig.Region = aws_aws.String(region)
	}
	if endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(endpoint)
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return err
	}

	svc := aws_s3.New(session)

	err = svc.ListObjectsPages(&aws_s3.ListObjectsInput{
		Bucket:  aws_aws.String(bucket_name),
		MaxKeys: aws_aws.Int64(1000),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			_, err = svc.DeleteObject(&aws_s3.DeleteObjectInput{
				Bucket: aws_aws.String(bucket_name),
				Key:    obj.Key,
			})
			if err != nil {
				log.Println(err)
			}
		}
		return true
	})
	if err != nil {
		log.Println(err)
		return err
	}

	svc.DeleteBucket(&aws_s3.DeleteBucketInput{
		Bucket: aws_aws.String(bucket_name),
	})

	return nil
}

func Upload(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	f io.ReadSeeker,
) error {
	return baseUpload(
		access_key,
		secret_key,
		region,
		endpoint,
		bucket_name,
		key,
		f,
		5,
		1,
	)
}

func DownloadToFile(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	localFilePath string,
) error {

	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			access_key,
			secret_key,
			"",
		),
	}
	if region != "" {
		awsConfig.Region = aws_aws.String(region)
	}
	if endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(endpoint)
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return err
	}

	file, err := os.Create(localFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(session)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket_name),
			Key:    aws.String(key),
		})
	if err != nil {
		return err
	}
	return nil
}

func baseUpload(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	f io.ReadSeeker,
	partSizeMB int,
	concurrency int,
) error {
	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			access_key,
			secret_key,
			"",
		),
	}
	if region != "" {
		awsConfig.Region = aws_aws.String(region)
	}
	if endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(endpoint)
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return err
	}
	uploader := aws_s3manager.NewUploader(session, func(u *aws_s3manager.Uploader) {
		u.PartSize = int64(partSizeMB) * 1024 * 1024 // The minimum/default allowed part size is 5MB
		u.Concurrency = concurrency                  // default is 5
	})

	size, _ := f.Seek(0, os.SEEK_END)
	f.Seek(0, 0)

	bar := pb.Full.Start64(size)

	// create proxy reader
	barReader := bar.NewProxyReader(f)

	_, err = uploader.Upload(&aws_s3manager.UploadInput{
		Bucket: aws_aws.String(bucket_name),
		ACL:    aws_aws.String("private"),
		Key:    aws_aws.String(key),
		Body:   barReader,
	})
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}

func GetObjectPresignUrl(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	ttl time.Duration,
) (string, error) {
	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			access_key,
			secret_key,
			"",
		),
	}
	if region != "" {
		awsConfig.Region = aws_aws.String(region)
	}
	if endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(endpoint)
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return "", err
	}
	svc := s3.New(session)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws_aws.String(bucket_name),
		Key:    aws_aws.String(key),
	})
	urlStr, err := req.Presign(ttl)

	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func RemoveObjectsByAge(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	age time.Duration,
) error {
	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			access_key,
			secret_key,
			"",
		),
	}
	if region != "" {
		awsConfig.Region = aws_aws.String(region)
	}
	if endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(endpoint)
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return err
	}

	svc := aws_s3.New(session)
	err = svc.ListObjectsPages(&aws_s3.ListObjectsInput{
		Bucket:  aws_aws.String(bucket_name),
		MaxKeys: aws_aws.Int64(1000),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			if time.Since(*obj.LastModified) > age {
				fmt.Println("removing", *obj.Key)
				_, err = svc.DeleteObject(&aws_s3.DeleteObjectInput{
					Bucket: aws_aws.String(bucket_name),
					Key:    obj.Key,
				})
				if err != nil {
					log.Println(err)
				}
			} else {
				fmt.Println("keeping", *obj.Key)
			}
		}
		return true
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetS3SecretsFromVaultOrEnvOrDie(vaultPath string) (
	string, string, string, string, string,
) {
	accessKeyVault, secretKeyVault, regionVault,
		endpointVault, bucketNameVault, _ := vault_s3_utils.GetS3Secrets("secret/data/slu/upload")

	// Access Key
	var accessKey string
	accessKeyEnv := os.Getenv("SLU_UPLOAD_ACCESS_KEY")
	if accessKeyVault != "" {
		accessKey = accessKeyVault
	}
	if accessKeyEnv != "" {
		accessKey = accessKeyEnv
	}
	if accessKey == "" {
		log.Fatalln("SLU_UPLOAD_ACCESS_KEY is empty")
	}

	// Secret Key
	var secretKey string
	secretKeyEnv := os.Getenv("SLU_UPLOAD_ACCESS_KEY")
	if secretKeyVault != "" {
		secretKey = secretKeyVault
	}
	if accessKeyEnv != "" {
		secretKey = secretKeyEnv
	}
	if accessKey == "" {
		log.Fatalln("SLU_UPLOAD_SECRET_KEY is empty")
	}

	// Region
	var region string
	regionEnv := os.Getenv("SLU_UPLOAD_REGION")
	if regionVault != "" {
		region = regionVault
	}
	if regionEnv != "" {
		region = regionEnv
	}

	// Endpoint
	var endpoint string
	endpointEnv := os.Getenv("SLU_UPLOAD_ENDPOINT")
	if endpointVault != "" {
		endpoint = endpointVault
	}
	if endpointEnv != "" {
		endpoint = endpointEnv
	}

	// Region, Endpoint Validation
	if region == "" && endpoint == "" {
		log.Fatalln("SLU_UPLOAD_REGION and SLU_UPLOAD_ENDPOINT are empty")
	}

	// Secret Key
	var bucketName string
	bucketNameEnv := os.Getenv("SLU_UPLOAD_BUCKET_NAME")
	if bucketNameVault != "" {
		bucketName = bucketNameVault
	}
	if bucketNameEnv != "" {
		bucketName = bucketNameEnv
	}
	if bucketName == "" {
		log.Fatalln("SLU_UPLOAD_BUCKET_NAME is empty")
	}

	return accessKey, secretKey, region, endpoint, bucketName
}
