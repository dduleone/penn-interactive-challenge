#!/bin/bash
mkdir data
curl https://datasets.imdbws.com/title.basics.tsv.gz | gunzip -c > data/title.basics.tsv