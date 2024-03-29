package s3crypto_test

import (
	"reflect"
	"testing"

	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/s3"
	"github.com/unicloud-uos/uos-sdk-go/service/s3/s3crypto"
)

func TestHeaderV2SaveStrategy(t *testing.T) {
	cases := []struct {
		env      s3crypto.Envelope
		expected map[string]*string
	}{
		{
			s3crypto.Envelope{
				CipherKey:             "Foo",
				IV:                    "Bar",
				MatDesc:               "{}",
				WrapAlg:               s3crypto.KMSWrap,
				CEKAlg:                s3crypto.AESGCMNoPadding,
				TagLen:                "128",
				UnencryptedMD5:        "hello",
				UnencryptedContentLen: "0",
			},
			map[string]*string{
				"X-Amz-Key-V2":                     aws.String("Foo"),
				"X-Amz-Iv":                         aws.String("Bar"),
				"X-Amz-Matdesc":                    aws.String("{}"),
				"X-Amz-Wrap-Alg":                   aws.String(s3crypto.KMSWrap),
				"X-Amz-Cek-Alg":                    aws.String(s3crypto.AESGCMNoPadding),
				"X-Amz-Tag-Len":                    aws.String("128"),
				"X-Amz-Unencrypted-Content-Md5":    aws.String("hello"),
				"X-Amz-Unencrypted-Content-Length": aws.String("0"),
			},
		},
		{
			s3crypto.Envelope{
				CipherKey:             "Foo",
				IV:                    "Bar",
				MatDesc:               "{}",
				WrapAlg:               s3crypto.KMSWrap,
				CEKAlg:                s3crypto.AESGCMNoPadding,
				UnencryptedMD5:        "hello",
				UnencryptedContentLen: "0",
			},
			map[string]*string{
				"X-Amz-Key-V2":                     aws.String("Foo"),
				"X-Amz-Iv":                         aws.String("Bar"),
				"X-Amz-Matdesc":                    aws.String("{}"),
				"X-Amz-Wrap-Alg":                   aws.String(s3crypto.KMSWrap),
				"X-Amz-Cek-Alg":                    aws.String(s3crypto.AESGCMNoPadding),
				"X-Amz-Unencrypted-Content-Md5":    aws.String("hello"),
				"X-Amz-Unencrypted-Content-Length": aws.String("0"),
			},
		},
	}

	for _, c := range cases {
		params := &s3.PutObjectInput{}
		req := &request.Request{
			Params: params,
		}
		strat := s3crypto.HeaderV2SaveStrategy{}
		err := strat.Save(c.env, req)
		if err != nil {
			t.Errorf("expected no error, but received %v", err)
		}

		if !reflect.DeepEqual(c.expected, params.Metadata) {
			t.Errorf("expected %v, but received %v", c.expected, params.Metadata)
		}
	}
}
