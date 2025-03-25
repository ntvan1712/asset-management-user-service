package error_app

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	duplicateKeyErrorCode       = "duplicate-key-error"
	documentNotFoundErrorCode   = "document-not-found-error"
	invalidAccessTokenErrorCode = "invalid-access-token-error"
	expiredAccessTokenErrorCode = "expired-access-token-error"
	wrongPasswordErrorCode      = "wrong-password-error"
	badRequestErrorCode         = "bad-request-error"
	unauthorizedErrorCode       = "unauthorized-error"
	permissionDeniedErrorCode   = "permission-denied-error"
	documentFormatErrorCode     = "document-format-error"
	premiumRequiredErrorCode    = "premium-required-error"
	userNotFoundErrorCode       = "user-not-found-error"
	createTempFileErrorCode     = "create-temp-file-error"
	bodyExceedsLimitErrCode     = "body-exceeds-limit-error"
)

var ErrBadRequest = errors.New(badRequestErrCode)
var ErrDocumentNotFound = errors.New(documentNotFoundErrorCode)
var ErrInvalidAccessToken = errors.New(invalidAccessTokenErrorCode)
var ErrExpiredAccessToken = errors.New(expiredAccessTokenErrorCode)
var ErrWrongPassword = errors.New(wrongPasswordErrorCode)
var ErrDuplicateKey = errors.New(duplicateKeyErrorCode)
var ErrUnauthorized = errors.New(unauthorizedErrorCode)
var ErrDocumentFormat = errors.New(documentFormatErrorCode)
var ErrPermissionDenied = errors.New(permissionDeniedErrorCode)
var ErrPremiumRequired = errors.New(premiumRequiredErrorCode)
var ErrUserNotFound = errors.New(userNotFoundErrorCode)
var ErrCreateTempFile = errors.New(createTempFileErrorCode)
var ErrBodyExceedsLimit = errors.New(bodyExceedsLimitErrCode)

func ErrCodeFromGRpcError(err error) error {
	grpcStatus, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch grpcStatus.Code() {
	case codes.InvalidArgument:
		return ErrBadRequest
	case codes.Unauthenticated:
		return ErrUnauthorized
	case codes.NotFound:
		return ErrDocumentNotFound
	default:
		return err
	}
}
