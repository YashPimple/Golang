## Simple-go-AWS-lambda

This project demonstrates the creation and execution of a simple AWS Lambda function written in Golang.

## Instructions

To use this project, follow the steps below:

1. Create an IAM Role using the AWS CLI command:

```bash
aws iam create-role --role-name lambda-ex --assume-role-policy-document '{"Version": "2012-10-17","Statement":[{"Effect": "Allow", "Principal": {"Service": "lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}'
```

2. Attach a policy to the created IAM Role using the AWS CLI command:

```bash
aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

3. Build the Go code by executing the following command:

```bash
go build main.go
```

4. Create a ZIP file containing the built binary using the command:

```bash
zip main.zip main
```

5. Create a Lambda function using the AWS CLI command:

```bash
aws lambda create-function --function-name Golang --runtime go1.x --role arn:aws:iam::036841205020:role/lambda-ex --handler main --zip-file fileb://main.zip
```

6. Invoke the Lambda function using the AWS CLI command:

```bash 
aws lambda invoke --function-name Golang --cli-binary-format raw-in-base64-out --payload '{"What is your name?": "Yash", "How old are you?": 19}' output.txt
```

Note: Make sure you have the AWS CLI installed and properly configured with the necessary credentials before executing the above commands.

Feel free to modify the payload in the last step according to your requirements.
