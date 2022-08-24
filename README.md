# aws demo1 service
## Dockerfile 
software: python+apache2

code from: git clone https://github.com/yz124/demo1.git

test url: /, /cgi-bin/test.py
## buildspec.yml
CodeBuild docker build, push to ECR

environment variable: AWS_ACCOUNT_ID, AWS_DEFAULT_REGION, IMAGE_REPO_NAME
## index.html, test.py
code files
## my-project
cdk project
