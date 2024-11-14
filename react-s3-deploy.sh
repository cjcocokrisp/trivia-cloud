#!/bin/bash

file=./src/webapp/src/config.js

# get ouputs from cdk output json
bucket=`jq -r '.TriviaCloud.WebsiteBucketName' ./result.json`
apiendpoint=`jq '.TriviaCloud.WebsocketApiEndpoint' ./result.json`

# create config js file
rm -rf $file
touch $file

# output cdk outputs to the config file
echo "const config = {" >> $file
echo "    API_ENDPOINT: $apiendpoint," >> $file
echo "};" >> $file
echo "export default config;" >> $file

# compile react app
cd src/webapp
npm run build

# deploy to s3 bucket
aws s3 sync build/ s3://$bucket/
