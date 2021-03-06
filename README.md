# penn-interactive-challenge

Go &amp; Terraform Coding Challenge

## In This Repository

```plaintext
/api/               Folder containing API Library
    /api.go         API Handler Functions
    /imdb.go        IMDB API Logic
/terraform/         Folder containing Terraform
    /main.tf        Main Terraform Script
    /userdata.sh    EC2 Userdata For Installing and Running Docker
    /vars.tf        Terraform Variable Configuration
/Dockerfile         Dockerfile for running locally
/README.md          This README
/getdata.sh         Script to fetch data
/rundocker.sh       Convenience script for running docker
/server.go          Main Go Execution
/testdeployment.sh  Script used to generate execution time analysis
```

## How to Use

First, fetch the data:

```bash
./getdata.sh
```

To run as Docker container:

```bash
./rundocker.sh
```

To run as Go script:

```bash
go run server.go
```

To run Terraform:

```bash
cd terraform
terraform apply
```

## Notes

### These routes were out of scope, but likely should be implemented

- /
- /movies

### Header Row

The instructions made no mention of the header record in the TSV, so I'm treating it like a normal record.

### Speed

In order to provide better performance, I would recommend caching results, since they shouldn't change from one query to another. Also, I'd like to experiment with the various methods for buffering and sending responses, to look for performance enhancements.
