## Updating lambda
1. `./build.sh`
2. `aws lambda update-function-code --function-name serverlessgotest --zip-file
   fileb://./deployment.zip --region eu-west-1`

## Creating Lambda
1. `aws lambda create-function --function-name <lambda-function-name> --zip-file fileb://./deployment.zip --runtime go1.x --handler main --role arn:aws:iam::ACCOUNT_ID:role/<new-lamda-fn-role-name>  --region us-west-1`
2. go to lambda services console and add `API Gateway` trigger
3. scroll down to the bottom to configure new trigger. Alternatively go to
   Amazon API Gateway and configure new trigger.
4. hit `Actions` dropdown and `Create Method`
5. add method [GET, POST, PUT, ...] and configure it to trigger new lambda
   function
