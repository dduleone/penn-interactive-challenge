#!/bin/bash
docker build -t penn-interactive-challenge .
docker run -it --rm -p 8080:80 --name penn-interactive-challenge penn-interactive-challenge