#!/bin/bash

curl http://${1}:8080/health -s > /dev/null
curl http://${1}:8080/health -s > /dev/null
curl http://${1}:8080/health -s > /dev/null
curl http://${1}:8080/health --create-dirs -o responses/health.json
curl http://${1}:8080/movies/tconst/tt0000492 -s > /dev/null
curl http://${1}:8080/movies/tconst/tt0000492 -s > /dev/null
curl http://${1}:8080/movies/tconst/tt0000492 -s > /dev/null
curl http://${1}:8080/movies/tconst/tt0000492 --create-dirs -o responses/movies_tconst_tt0000492.json
curl http://${1}:8080/movies/startYear/2000 -s > /dev/null
curl http://${1}:8080/movies/startYear/2000 -s > /dev/null
curl http://${1}:8080/movies/startYear/2000 -s > /dev/null
curl http://${1}:8080/movies/startYear/2000 --create-dirs -o responses/movies_startYear_2000.json
curl http://${1}:8080/movies/genre/drama -s > /dev/null
curl http://${1}:8080/movies/genre/drama -s > /dev/null
curl http://${1}:8080/movies/genre/drama -s > /dev/null
curl http://${1}:8080/movies/genre/drama --create-dirs -o responses/movies_genre_draa.json
