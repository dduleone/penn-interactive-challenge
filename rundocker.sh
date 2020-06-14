#!/bin/bash
docker build -t penn-interactive-challenge .
docker run -it --rm -p 80:80 --name penn-interactive-challenge penn-interactive-challenge