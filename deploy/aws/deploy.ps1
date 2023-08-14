param(
    [Parameter(Mandatory=$true)]
    [string]$domainName
)

$hostedZoneId = (((aws route53 list-hosted-zones |
        ConvertFrom-Json).HostedZones |
        Where-Object -FilterScript {$_.Name -EQ "$domainName."} |
        Select-Object -First 1).Id).Split("/")[2]

$certificateArn = ((aws acm list-certificates --region us-west-2 | ConvertFrom-Json).CertificateSummaryList |
                Where-Object -FilterScript {$_.DomainName -EQ "$domainName"} |
                Select-Object -First 1).CertificateArn

$serviceName = $domainName.Split(".")[0]

aws cloudformation create-stack --stack-name $serviceName --parameters `
    ParameterKey=ServiceName,ParameterValue=$serviceName `
    ParameterKey=CertificateArn,ParameterValue=$certificateArn `
    ParameterKey=HostedZoneId,ParameterValue=$hostedZoneId `
    ParameterKey=DomainName,ParameterValue=$domainName `
    --template-body file://infra.yaml `
    --capabilities CAPABILITY_IAM

