import os
from aws_cdk import (
    Duration,
    Stack,
    RemovalPolicy,
    CfnOutput,
    aws_dynamodb as dynamodb,
    aws_s3 as s3,
    aws_s3_deployment as s3_deployment,
    aws_apigatewayv2 as apigateway,
    aws_lambda as _lambda,
    aws_lambda_go_alpha as go_lambda,
    aws_apigatewayv2_integrations as integrations
)
from constructs import Construct

class TriviaCloudStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # Constructs for websocket connects + backend
        data_table = dynamodb.Table(
            self, 
            f'DataTable{construct_id}', 
            partition_key=dynamodb.Attribute(name='gameId', type=dynamodb.AttributeType.STRING),
            removal_policy=RemovalPolicy.DESTROY
        )

        player_table = dynamodb.Table(
            self, 
            f'PlayerTable{construct_id}', 
            partition_key=dynamodb.Attribute(name='connectionId', type=dynamodb.AttributeType.STRING),
            removal_policy=RemovalPolicy.DESTROY
        )

        websocket_api = apigateway.WebSocketApi(
            self, 
            f'WebsocketAPI{construct_id}'
        )

        connect_lambda = go_lambda.GoFunction(self, 
            f'ConnectID{construct_id}', 
            function_name=f'{construct_id}Connect',
            environment= {
                'DATA_TABLE': data_table.table_name,
                'PLAYER_TABLE': player_table.table_name,
            },
            entry='src/api/connect', 
            timeout=Duration.seconds(300)
        )
        
        disconnect_lambda = go_lambda.GoFunction(
            self, 
            f'DisconnectID{construct_id}', 
            function_name=f'{construct_id}Disconnect',
            environment= {
                'DATA_TABLE': data_table.table_name,
                'PLAYER_TABLE': player_table.table_name,
            },
            entry='src/api/disconnect', 
            timeout=Duration.seconds(300)
        )

        data_table.grant_read_write_data(connect_lambda)
        data_table.grant_read_write_data(disconnect_lambda)
        player_table.grant_read_write_data(connect_lambda)
        player_table.grant_read_write_data(disconnect_lambda)

        websocket_api.add_route('$connect', 
            integration=integrations.WebSocketLambdaIntegration(f'ConnectIntegration{construct_id}', connect_lambda)
        )
        websocket_api.add_route('$disconnect', 
            integration=integrations.WebSocketLambdaIntegration(f'DisconnectIntegration{construct_id}', disconnect_lambda)
        )

        # Constructs for serverless react app
        website_bucket = s3.Bucket(
            self, 'WebsiteBucket', 
            website_index_document='index.html', 
            block_public_access=s3.BlockPublicAccess(block_public_policy=False, block_public_acls=False), 
            public_read_access=True, 
            removal_policy=RemovalPolicy.DESTROY, 
            auto_delete_objects=True
        )

        CfnOutput(self, 'WebsiteBucketName', value=website_bucket.bucket_name, export_name='WebsiteBucketName')
        CfnOutput(self, 'WebsiteBucketURL', value=website_bucket.bucket_website_url, export_name='WebsiteBucketURL')
        CfnOutput(self, 'WebsocketApiEndpoint', value=websocket_api.api_endpoint, export_name='WebsocketApiEndpoint')