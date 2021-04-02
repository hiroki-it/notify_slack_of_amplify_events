module github.com/Hiroki-IT/notify_slack_of_amplify_events

go 1.15

require (
	github.com/aws/aws-lambda-go v1.23.0
	github.com/aws/aws-sdk-go-v2 v1.3.0
	github.com/aws/aws-sdk-go-v2/config v1.1.3
	github.com/aws/aws-sdk-go-v2/service/amplify v1.1.3
	github.com/golang/mock v1.5.0
	github.com/stretchr/testify v1.7.0
)

replace github.com/Hiroki-IT/notify_slack_of_amplify_events/cmd/slack => ./slack
