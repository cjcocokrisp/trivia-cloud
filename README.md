<div align="center">
    <img src="src/webapp/src/logo.svg" alt="Logo" width="50%" height="50%"/>
    <h3>A Real Time Trivia Game Built Using AWS</h3>
</div>

## About

Trivia Cloud is a real time trivia game built using AWS. It offers real time gameplay similar to something like Kahoot. It used the [Open Trivia Database](https://opentdb.com/) for questions to allow for different questions based on a variety of categories. Game flow is simple, players will join a lobby and then answer questions and then after answering they will be told if they are right or wrong. Unlike other games like this, the scores of the others will not be revealed until the end of the game to leave you in the dark about how the people you are playing against are doing. Built for the Fall 2024 COMP.4600 Selected Topics: Cloud Computing final project.

## Technology Used

Trivia Cloud uses a variety of different technologies. As mentioned it uses AWS for its main services. Below is a list of technologies and AWS services that this project is using. When creating or joining a game, Websockets are used to connect players to each other to communicate game data amongest themselves and with the server.

### Main Technologies
- React.JS (Frontend Website)
- Go (Backend Functionality)

### AWS Services
- Lambda (Websocket functions and methods)
- API Gateway (Websocket API route management)
- DynamoDB (Hold data about current connections and games)
- S3 (Hosts compiled React site)
- CDK (Used for deployment and management of AWS resources)

### Other Technologies
- GitHub Actions (Used for automated deployments)
- Bash Scripting (Used for generating a config file based on the output of the CDK deploy, building the React site, and uploading it to S3)

## Installation

To run Trivia Cloud locally and on your own AWS account follow these steps

1. Make a `.env` file and copy the below contents into the file and replace it with your information.
2. Run `cdk deploy --outputs-file ./result.json` to create AWS resources (make sure you have the outputs-file flag or it will not work, the bash script relies on the json file).
3. Run `./react-s3-deploy.sh` (if you get permission errors do `chmod +x ./react-s3-deploy.sh`) to upload the built React app to the S3 bucket.

Template .env file for project:
```
APP_NAME=<WHATEVER YOU WANT TO NAME THE APP>
AWS_ACCESS_ID=<YOUR AWS ACCESS ID>
AWS_SECRET_ID=<YOUR AWS SECRET ID>
AWS_ACCOUNT_ID=<YOUR AWS ACCOUNT ID>
AWS_DEFAULT_REGION=<THE AWS REGION YOU WANT IT TO BE DEPLOYED TO>
```

## Acknowledgements

- Christopher Coco - Developer 
- Adam Corkhum - Developer
- Layann Shaban - Developer
- Vina Dang - Developer
- Professor Johannes Weis - Professor for course
- [Open Trivia Database](https://opentdb.com/) - For providing an API to use for questions
- [AWS Samples](https://github.com/aws-samples) - For the various samples provided
- [Namecheap Logo Maker](https://www.namecheap.com/logo-maker/) - Resource used for generating the logo