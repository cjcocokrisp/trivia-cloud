#!/usr/bin/env python3
import os

import aws_cdk as cdk
from dotenv import load_dotenv

from trivia_cloud.trivia_cloud_stack import TriviaCloudStack

load_dotenv()

app = cdk.App()
TriviaCloudStack(app, os.getenv('APP_NAME'), 
                 env=cdk.Environment(account=os.getenv('AWS_ACCOUNT_ID'), region=os.getenv('AWS_DEFAULT_REGION')),
)

app.synth()
