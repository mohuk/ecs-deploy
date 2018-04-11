# ECS Deploy
> Binary to update and deploy services on AWS ECS

## Install
```bash
# install with go get and make sure you have $GOPATH/bin in your path
λ go get github.com/mohuk/ecs-deploy/cmd/ecs-deploy

# make with source
λ git clone https://github.com/mohuk/ecs-deploy
λ cd ecs-deploy
λ make build

# create cross platform binaries
λ git clone https://github.com/mohuk/ecs-deploy
λ cd ecs-deploy
λ make cross
```
or you can download the pre-compiled version of binaries from [here](https://github.com/mohuk/ecs-deploy/releases/tag/0.1.0).

## Usage

Global Help:
```bash
λ ecs-deploy --help

Usage of ./ecs-deploy_darwin_386:
  -c string
    	ECS Cluster name
  -i string
    	Image name
  -r string
    	Region (default "us-east-1")
  -s string
    	ECS Service to deploy
```

Deploy:
```bash
λ ecs-deploy -c ecs-cluster -s ecs-service -i ecs-image
```
