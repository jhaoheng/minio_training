package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go-v2/aws/endpoints"
  "github.com/aws/aws-sdk-go-v2/aws/external"
)

func main() {
  // Using the SDK's default configuration, loading additional config
  // and credentials values from the environment variables, shared
  // credentials, and shared configuration files
  cfg, err := external.LoadDefaultAWSConfig()
  if err != nil {
    panic("unable to load SDK config, " + err.Error())
  }

  // Set the AWS Region that the service clients should use
  cfg.Region = endpoints.UsWest2RegionID
  fmt.Println(cfg)

}
