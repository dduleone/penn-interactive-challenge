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
/rundocker.sh       Convenience script for running docker
/server.go          Main Go Execution
```

## Notes

### These routes were out of scope, but likely should be implemented

- /
- /movies

### Header Row

The instructions made no mention of the header record in the TSV, so I'm treating it like a normal record.
