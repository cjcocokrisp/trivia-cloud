from aws_cdk import (
    # Duration,
    Stack,
    aws_dynamodb as dynamodb,
    RemovalPolicy
)
from constructs import Construct

class TriviaCloudStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        table = dynamodb.Table(self, "Data", 
                               partition_key=dynamodb.Attribute(name="gameID", type=dynamodb.AttributeType.NUMBER),
                               removal_policy=RemovalPolicy.DESTROY)
