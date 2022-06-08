# configuration-fetch-action

> This project is powered by AWS SDK

A Github Action to fetches configurations from AWS AppConfig files, powered by AWS SDK
## Requirements
- AWS AppConfig Application should exist already
- AWS Appconfig Application Configuration profile should exist already
- AWS AppConfig Environment should exist already

- [AWS configuration](https://github.com/aws-actions/configure-aws-credentials) should be set beforehand. It is useful to have a AWS IAM secret access key pair stored in the Github Secrets with permissions to access the public API.


## Usage

### Inputs

| Input                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `appName`  | AWS AppConfig application name input    |
| `profileName`  | AWS AppConfig profile name input    |
| `env`| AWS AppConfig environment name|
| `region`| AWS Region|



### Example workflow

```yaml
  - name: Configure AWS Dev account credentials
    uses: aws-actions/configure-aws-credentials@v1
    with:
      aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID_DEV }}
      aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY_DEV }}
      aws-region: "eu-west-1"

  - name: Fetch AWS AppConfig content
    uses: carlsberg/configuration-fetch-action@main
    with:
      appName: "beer-application"
      profileName: "beer-delivery-service"
      env: "dev"
      region: "eu-west-1"
```
### Other tips 

Having repo with a previously convened directory structure can be helpfull.
Something along the following lines: 

```
.
└── beer-application
    ├── dev
    │   ├── beer-delivery-service
    │   │   └── config.json
    │   └── beer-order-service
    │       └── config.json
    └── tst
        ├── beer-delivery-service
        │   └── config.json
        └── beer-order-service
            └── config.json
```

