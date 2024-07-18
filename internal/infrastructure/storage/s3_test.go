package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type mockDeleteObjectAPI func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)

func (m mockDeleteObjectAPI) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	return m(ctx, params, optFns...)
}

func TestS3_DeleteFile(t *testing.T) {
	mockClient := func(t *testing.T) S3DeleteObjectAPI {
		return mockDeleteObjectAPI(func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
			t.Helper()
			if params.Bucket == nil {
				t.Fatal("expect bucket to not be nil")
			}
			if e, a := "fooBucket", *params.Bucket; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
			if params.Key == nil {
				t.Fatal("expect key to not be nil")
			}

			if *params.Key == "invalidKey" {
				return nil, errors.New("Client error")
			}

			if e, a := "barKey", *params.Key; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}

			return &s3.DeleteObjectOutput{}, nil
		})
	}

	type fields struct {
		client func(t *testing.T) S3DeleteObjectAPI
	}
	type args struct {
		bucket string
		key    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"DeleteSuccessfully", fields{client: mockClient}, args{"fooBucket", "barKey"}, false},
		{"ClientError", fields{client: mockClient}, args{"fooBucket", "invalidKey"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &S3{
				client: tt.fields.client(t),
			}
			if err := s.DeleteFile(tt.args.bucket, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("S3.DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
