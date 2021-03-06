package cos

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int

	// CI client
	Client *cos.Client

	// Copy source client
	CClient *cos.Client

	// test_object
	TestObject string

	// special_file_name
	SepFileName string
}

var (
	UploadID string
	PartETag string
)

// 开启存储桶日志服务
func (s *CosTestSuite) putBucketLogging() {
  client := s.Client
  //.cssg-snippet-body-start:[put-bucket-logging]
  opt := &cos.BucketPutLoggingOptions{
      LoggingEnabled: &cos.BucketLoggingEnabled{
          TargetBucket: TargetBucket,
      },  
  }   
  resp, err := client.Bucket.PutLogging(context.Background(), opt)
  
  //.cssg-snippet-body-end
}

// 获取存储桶日志服务
func (s *CosTestSuite) getBucketLogging() {
  client := s.Client
  //.cssg-snippet-body-start:[get-bucket-logging]
  v, resp, err := client.Bucket.GetLogging(context.Background())
  
  //.cssg-snippet-body-end
}

//.cssg-methods-pragma


func TestCOSTestSuite(t *testing.T) {
	suite.Run(t, new(CosTestSuite))
}

func (s *CosTestSuite) TestBucketLogging() {
	// 将 examplebucket-1250000000 和 ap-guangzhou 修改为真实的信息
	u, _ := url.Parse("https://examplebucket-1250000000.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_KEY"),
			SecretKey: os.Getenv("COS_SECRET"),
		},
	})
  s.Client = c

	// 开启存储桶日志服务
	s.putBucketLogging()

	// 获取存储桶日志服务
	s.getBucketLogging()

	//.cssg-methods-pragma
}
