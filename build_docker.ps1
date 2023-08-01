param(
    [Parameter(Mandatory=$true)]
    [string]$region,

    [Parameter(Mandatory=$true)]
    [string]$location,

    [Parameter(Mandatory=$true)]
    [string]$gtag
)

$accountId = (aws sts get-caller-identity | ConvertFrom-Json).Account

docker build --build-arg GTAG=$gtag --build-arg LOCATION=$location -t willitsnow:latest .
docker tag willitsnow:latest "$accountId.dkr.ecr.$region.amazonaws.com/willitsnow_private:latest"
aws ecr get-login-password --region $region | docker login --username AWS --password-stdin "$accountId.dkr.ecr.$region.amazonaws.com"
docker push "$accountId.dkr.ecr.$region.amazonaws.com/willitsnow_private:latest"