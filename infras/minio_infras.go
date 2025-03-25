package infras

import (
	"context"
	"sync"
	"user_service/app_config"
	"user_service/common/logger"
	app_utils "user_service/common/utils"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioProvider *MinioProvider
var initMinioOnce sync.Once

type MinioProvider struct {
	MinioClient *minio.Client
	AssetBucket string
}

func GetMinioProvider() *MinioProvider {
	initMinioOnce.Do(initMinioProvider)
	return minioProvider
}

func initMinioProvider() {
	config := app_config.GetAppConfig().MinioConfig

	// Initialize minio client object.
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: false, // http
	})
	if err != nil {
		logger.Fatal("[MinioInfras] init error", err)
	}
	minioProvider = &MinioProvider{
		MinioClient: minioClient,
		AssetBucket: config.AssetBucket,
	}

	logger.Info("[MinioInfras] Init MinioProvider")
}

func TestUpload() {
	objectName := "asset_labels/4.jpg"
	filePath := "4.jpg"

	info, err := GetMinioProvider().MinioClient.FPutObject(
		context.Background(), GetMinioProvider().AssetBucket, objectName, filePath,
		minio.PutObjectOptions{ContentType: app_utils.GetContentType(filePath)})
	if err != nil {
		logger.Fatal("[Upload]", err)
	}
	logger.Info("[UploadInfo]", info)
}

func TestDelete() {
	objectName := "4.jpg"
	// filePath := "4.jpg"

	err := GetMinioProvider().MinioClient.RemoveObject(
		context.Background(), GetMinioProvider().AssetBucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		logger.Fatal("[Upload]", err)
	}
	// logger.Info("[UploadInfo]", info)
}
