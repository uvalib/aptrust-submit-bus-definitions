module github.com/uvalib/apts-bus-pub

go 1.26.0

//require github.com/uvalib/aptrust-submit-bus-definitions/uvaaptsbus v0.0.0-20250801130056-157231a1fcac

// for local development
replace github.com/uvalib/aptrust-submit-bus-definitions/uvaaptsbus => ../../uvaaptsbus

require github.com/uvalib/aptrust-submit-bus-definitions/uvaaptsbus v0.0.0-20260303144802-713a35019895

require (
	github.com/aws/aws-sdk-go-v2 v1.41.3 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.32.11 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.19.11 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.32.21 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.0.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.30.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.35.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.41.8 // indirect
	github.com/aws/smithy-go v1.24.2 // indirect
)
