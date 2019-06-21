#!/bin/bash

./kill.sh
./dockerBuildPush.sh
# docker pull patientplatypus/secretsquirrel_nest:latest
# docker run -it --rm --name patientplatypus/secretsquirrel_nest:latest
docker run -it -p 8000:8000 --name secretsquirrel_nest patientplatypus/secretsquirrel_nest:latest