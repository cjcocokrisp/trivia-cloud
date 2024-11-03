from aws_cdk import (
    # Duration,
    Stack,
    RemovalPolicy,
    CfnOutput,
    aws_dynamodb as dynamodb,
    aws_s3 as s3,
    aws_s3_deployment as s3_deployment
)
from constructs import Construct

class TriviaCloudStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        table = dynamodb.Table(self, 'Data', 
                               partition_key=dynamodb.Attribute(name='gameID', type=dynamodb.AttributeType.NUMBER),
                               removal_policy=RemovalPolicy.DESTROY)

        website_bucket = s3.Bucket(self, 'websiteBucket', website_index_document='index.html', 
                                   block_public_access=s3.BlockPublicAccess(
                                       block_public_policy=False, block_public_acls=False
                                    ), public_read_access=True, removal_policy=RemovalPolicy.DESTROY, auto_delete_objects=True)
        
        s3_deployment.BucketDeployment(self, 'websiteDeploy', sources=[s3_deployment.Source.asset('./src/frontend/build/')], 
                                       destination_bucket=website_bucket)
        
        CfnOutput(self, 'BucketExport', value=website_bucket.bucket_website_url, export_name='websiteBucketName')