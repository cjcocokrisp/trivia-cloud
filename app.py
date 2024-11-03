#!/usr/bin/env python3
import os

import aws_cdk as cdk
from dotenv import load_dotenv

from trivia_cloud.trivia_cloud_stack import TriviaCloudStack

load_dotenv()

app = cdk.App()
TriviaCloudStack(app, os.getenv("APP_NAME"),)

app.synth()
