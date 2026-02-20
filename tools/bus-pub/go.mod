module github.com/uvalib/apts-bus-definitions

go 1.26.0

//require github.com/uvalib/librabus-sdk/uvalibrabus v0.0.0-20250801130056-157231a1fcac

// for local development
replace github.com/uvalib/apts-bus-definitions/uvaaptsbus => ../../uvaaptsbus

require github.com/uvalib/apts-bus-definitions/uvaaptsbus v0.0.0-00010101000000-000000000000

require (
	github.com/aws/aws-sdk-go-v2 v1.41.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.32.9 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.19.9 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.32.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.30.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.35.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.41.6 // indirect
	github.com/aws/smithy-go v1.24.0 // indirect
)
