# configuration-fetch-action

> This project is powered by AWS SDK

A Github Action that fetches configurations from AWS AppConfig files, powered by AWS SDK. It creates a file and sets the absolute path as a github action output for the file location.

## Requirements

- AWS AppConfig Application should exist already
- AWS Appconfig Application Configuration profile should exist already
- AWS AppConfig Environment should exist already
- [AWS configuration](https://github.com/aws-actions/configure-aws-credentials) should be set beforehand. It is useful to have a AWS IAM secret access key pair stored in the Github Secrets with permissions to access the public API.

## Usage

### Inputs

| Input                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `app-name`  | AWS AppConfig application name input    |
| `profile-name`  | AWS AppConfig profile name input    |
| `env`| AWS AppConfig environment name|
| `region`| AWS Region|

### Outputs

| Output                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `absolute-path`  | The location of the file with a copy of the requested configuration     |

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
      app-name: "beer-application"
      profile-name: "beer-delivery-service"
      env: "dev"
      region: "eu-west-1"
```
### Other tips 

We recommend repositories to have the following directory structure: 
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

