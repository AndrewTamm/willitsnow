# Will It Snow?

This is a single page webapp which displays the weather forecast for the area and indicates if it will snow in the next week.
It is a single Docker container which has a static React frontend served by a golang backend, which also serves as a cache
and proxy to a Lambda function which fetches the weather information.

## How to use locally

1. Clone the repository.
2. Install and configure the [AWS Cli](https://github.com/aws/aws-cli) based on your OS.
3. The backend expects to find a Lambda function named `FetchWeatherLambda`. The cloudformation templates will create a function for this purpose.
4. Change directory to `/frontend`
5. Install the dependencies using `npm install`.
6. Run the app using `npm start`.
7. Open your browser and navigate to `http://localhost:3000`.

## Building the container

The `build_docker.ps1` script automates the process of building a Docker container image, tagging it appropriately for Amazon Elastic Container Registry (ECR), logging in to the ECR, and pushing the built image to ECR. This allows for a seamless deployment and integration of Docker containers with Amazon Web Services (AWS).

### Prerequisites

1. [Docker](https://www.docker.com/) installed and running on your machine.
2. [AWS Command Line Interface (CLI)](https://aws.amazon.com/cli/) installed and configured.
3. AWS permissions to:
   - Fetch the account ID using the STS service.
   - Log in to ECR.
   - Push images to ECR.
4. The Dockerfile for building the container image should be present in the directory from which the script is run.

### Usage

To use the script, navigate to the directory containing the Dockerfile and execute the script in a PowerShell environment providing the necessary parameters:

```powershell
./build_docker.ps1 -region yourAWSregion -location yourLocation -gtag yourGTag
```

Replace `yourAWSregion` with the AWS region for your ECR, `yourLocation` with the desired location parameter for the Docker build, and `yourGTag` with the desired GTag parameter for the Docker build.

### Parameters

- `region` (Mandatory): The AWS region where your ECR resides.
- `location` (Mandatory): A build argument representing the location used within the Docker build process.
- `gtag` (Mandatory): A build argument representing the GTag used within the Docker build process.

### Output

Upon successful execution, the script will:

1. Build a Docker image named `willitsnow:latest` using the provided `location` and `gtag` parameters.
2. Tag the image appropriately for pushing to AWS ECR.
3. Authenticate with the ECR.
4. Push the Docker image to your ECR repository named `willitsnow_private`.

### Important Notes

1. The script retrieves your AWS account ID using the STS service and uses this ID to correctly tag and push your Docker image.
2. Ensure your Docker environment is running and you have appropriate permissions to build and push images.
3. The destination ECR repository `willitsnow_private` should already exist in the specified region. If it doesn't, you'll need to create it before running the script.

### Disclaimer

This script interacts with AWS services and may incur charges, especially when pushing images to ECR. Always ensure you understand the actions being performed and review the script for any unintended side effects.

## How to deploy

This PowerShell script is designed to automate the process of:

1. Retrieving the hosted zone ID from Amazon Route53.
2. Retrieving the ARN (Amazon Resource Name) of a certificate from Amazon Certificate Manager (ACM) that matches a provided domain name.
3. Creating a CloudFormation stack using the retrieved hosted zone ID, certificate ARN, and other parameters based on the provided domain name.

### Prerequisites

1. [AWS Command Line Interface (CLI)](https://aws.amazon.com/cli/) installed and configured.
2. PowerShell environment (Windows PowerShell or PowerShell Core).
3. AWS permissions to:
    - List Route53 hosted zones.
    - List ACM certificates.
    - Create CloudFormation stacks.
    - IAM capabilities for the CloudFormation stack.
4. The `infra.yaml` CloudFormation template should exist in the same directory as the script.

### Usage

To use the script, execute it in a PowerShell environment and provide the domain name as a parameter.

```powershell
cd deploy/aws
./deploy.ps1 -domainName yourdomain.com
```

### Parameters

- `domainName` (Mandatory): The domain name for which you want to retrieve the hosted zone ID and certificate ARN. This will also be used to determine the service name by extracting the subdomain.

### Output

Upon successful execution, the script will create a CloudFormation stack using the provided domain, hosted zone ID from Route53, and certificate ARN from ACM.

### Important Notes

1. This script assumes that the domain's hosted zone is already created in Route53.
2. The certificate for the domain should exist in the ACM (Amazon Certificate Manager) under the region `us-west-2`.
3. The infra.yaml template file contains the following parameters:
   - ServiceName: A tag to place on all resources created by this formation.
   - CertificateArn: The ARN of the ACM certificate.
   - HostedZoneId: The ID of the HostedZone for the DNS domain.
   - DomainName: The DNS domain name.
   - AlternativeHostedZoneId: The ID of the HostedZone for the DNS domain. By default, this value is empty.
   - AlternativeDomainName: An alternative DNS domain name. By default, this value is empty.
   - ContainerPort: The port number the application inside the Docker container is binding to. The default is 8080.
   - HostPort: The published port exposed on the Docker host. The default is 8080.

   Ensure that you provide appropriate values for these parameters when deploying the CloudFormation stack using the script, and be aware of the default values.
4. Ensure that your AWS CLI is correctly configured with the necessary permissions to execute all commands in the script.
5. Always validate the `infra.yaml` template before deploying it with the script to ensure there are no errors.

### Disclaimer

This script interacts with your AWS resources and may incur charges. Always ensure you understand the actions being performed, and review the script and CloudFormation template for any unintended side effects.

---

## Contributing

If you want to contribute to this project, please follow these steps:

1. Fork this repository.
2. Create a new branch.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Special Thanks

Heavily inspired  by the fine work at [Will It Snow In PDX?](https://www.willitsnowinpdx.com/) who gave me the idea and made
me realize that I needed the same information for SEA.
