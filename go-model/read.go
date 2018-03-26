package main

import (
  "fmt"
  "os"

  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/aws/endpoints"
  "github.com/aws/aws-sdk-go-v2/aws/external"
  "github.com/aws/aws-sdk-go-v2/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
  fmt.Fprintf(os.Stderr, msg+"\n", args...)
  os.Exit(1)
}

// Lists all objects in a bucket using pagination
//
// Usage:
// listObjects <bucket>
func main() {
  if len(os.Args) < 2 {
    exitErrorf("you must specify a bucket")
  }

  defaultResolver := endpoints.NewDefaultResolver()
  s3CustResolverFn := func(service, region string) (aws.Endpoint, error) {
    if service == "s3" {
      return aws.Endpoint{
        URL:           "localhost:9000",
        SigningRegion: "custom-signing-region",
      }, nil
    }

    return defaultResolver.ResolveEndpoint(service, region)
  }

  cfg, err := external.LoadDefaultAWSConfig()
  if err != nil {
    panic("failed to load config, " + err.Error())
  }
  cfg.Region = "us-east-1"
  cfg.EndpointResolver = aws.EndpointResolverFunc(s3CustResolverFn)

  // Create the S3 service client with the shared config. This will
  // automatically use the S3 custom endpoint configured in the custom
  // endpoint resolver wrapping the default endpoint resolver.
  svc := s3.New(cfg)

  req := svc.ListObjectsRequest(&s3.ListObjectsInput{Bucket: &os.Args[1]})
  p := req.Paginate()
  for p.Next() {
    page := p.CurrentPage()
    for _, obj := range page.Contents {
      fmt.Println("Object: ", *obj.Key)
    }
  }

  if err := p.Err(); err != nil {
    exitErrorf("failed to list objects, %v", err)
  }
}
