package bos

import (
	"strconv"
	"testing"
	"time"

	"baidubce/test"
)

var bosClient = DefaultClient

func TestGetBucketLocation(t *testing.T) {
	bucketNamePrefix := "baidubce-sdk-go-test-for-get-bucket-location-"
	method := "GetBucketLocation"

	around(t, method, bucketNamePrefix, func(bucketName string) {
		expected := "bj"
		location, _ := bosClient.GetBucketLocation(bucketName, nil)

		if location.LocationConstraint != expected {
			t.Error(test.Format(method, location.LocationConstraint, expected))
		}
	})
}

func TestListBuckets(t *testing.T) {
	_, err := bosClient.ListBuckets(nil)

	if err != nil {
		t.Error(test.Format("ListBuckets", err.Error(), "nil"))
	}
}

func TestCreateBucket(t *testing.T) {
	bucketNamePrefix := "baidubce-sdk-go-test-for-create-bucket-"
	method := "CreateBucket"

	around(t, method, bucketNamePrefix, nil)
}

func TestDoesBucketExist(t *testing.T) {
	bucketNamePrefix := "baidubce-sdk-go-test-for-does-bucket-exist-"
	method := "DoesBucketExist"

	around(t, method, bucketNamePrefix, func(bucketName string) {
		expected := true
		exists, err := bosClient.DoesBucketExist(bucketName, nil)

		if err != nil {
			t.Error(test.Format(method, err.Error(), strconv.FormatBool(expected)))
		} else if exists != expected {
			t.Error(test.Format(method, strconv.FormatBool(exists), strconv.FormatBool(expected)))
		}
	})

}

func TestDeleteBucket(t *testing.T) {
	bucketNamePrefix := "baidubce-sdk-go-test-for-delete-bucket-"
	method := "DeleteBucket"

	around(t, method, bucketNamePrefix, nil)
}

func TestSetBucketPrivate(t *testing.T) {
	bucketNamePrefix := "baidubce-sdk-go-test-for-set-bucket-private-"
	method := "SetBucketPrivate"

	around(t, method, bucketNamePrefix, func(bucketName string) {
		err := bosClient.SetBucketPrivate(bucketName, nil)
		if err != nil {
			t.Error(test.Format(method, err.Error(), "nil"))
		}
	})
}

func TestSetBucketPublicRead(t *testing.T) {
	bucketNamePrefix := "baidubce-sdk-go-test-for-set-bucket-public-read-"
	method := "SetBucketPublicRead"

	around(t, method, bucketNamePrefix, func(bucketName string) {
		err := bosClient.SetBucketPublicRead(bucketName, nil)
		if err != nil {
			t.Error(test.Format(method, err.Error(), "nil"))
		}
	})
}

func TestSetBucketPublicReadWrite(t *testing.T) {
	bucketNamePrefix := "baidubce-sdk-go-test-for-set-bucket-public-rw-"
	method := "SetBucketPublicReadWrite"

	around(t, method, bucketNamePrefix, func(bucketName string) {
		err := bosClient.SetBucketPublicReadWrite(bucketName, nil)
		if err != nil {
			t.Error(test.Format(method, err.Error(), "nil"))
		}
	})
}

func TestGetBucketAcl(t *testing.T) {
	bucketNamePrefix := "baidubce-sdk-go-test-for-get-bucket-acl-"
	method := "GetBucketAcl"

	around(t, method, bucketNamePrefix, func(bucketName string) {
		_, err := bosClient.GetBucketAcl(bucketName, nil)
		if err != nil {
			t.Error(test.Format(method, err.Error(), "nil"))
		}
	})
}

func around(t *testing.T, method, bucketNamePrefix string, f func(string)) {
	bucketName := bucketNamePrefix + strconv.Itoa(int(time.Now().Unix()))
	err := bosClient.CreateBucket(bucketName, nil)

	if err != nil {
		t.Error(test.Format(method+" at creating bucket", err.Error(), "nil"))
	} else {
		if f != nil {
			f(bucketName)
		}

		err = bosClient.DeleteBucket(bucketName, nil)
		if err != nil {
			t.Error(test.Format(method+" at deleting bucket", err.Error(), "nil"))
		}
	}
}
