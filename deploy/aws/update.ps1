param(
    [Parameter(Mandatory=$true)]
    [string]$domainName,

    [string]$alternativeDomainName
)

$hostedZoneId = (((aws route53 list-hosted-zones |
        ConvertFrom-Json).HostedZones |
        Where-Object -FilterScript {$_.Name -EQ "$domainName."} |
        Select-Object -First 1).Id).Split("/")[2]

$certificateArn = ((aws acm list-certificates --region us-west-2 | ConvertFrom-Json).CertificateSummaryList |
        Where-Object -FilterScript {$_.DomainName -EQ "$domainName"} |
        Select-Object -First 1).CertificateArn

$serviceName = $domainName.Split(".")[0]

if ($alternativeDomainName) {
    $alternativeHostedZoneId = (((aws route53 list-hosted-zones |
            ConvertFrom-Json).HostedZones |
            Where-Object -FilterScript {$_.Name -EQ "$alternativeDomainName."} |
            Select-Object -First 1).Id).Split("/")[2]
}

aws cloudformation create-change-set --change-set-name SampleChangeSet --stack-name $serviceName --parameters `
    ParameterKey=ServiceName,ParameterValue=$serviceName `
    ParameterKey=CertificateArn,ParameterValue=$certificateArn `
    ParameterKey=HostedZoneId,ParameterValue=$hostedZoneId `
    ParameterKey=DomainName,ParameterValue=$domainName `
    ParameterKey=AlternativeDomainName,ParameterValue=$alternativeDomainName `
    ParameterKey=AlternativeHostedZoneId,ParameterValue=$alternativeHostedZoneId `
    --template-body file://infra.yaml `
    --capabilities CAPABILITY_IAM

