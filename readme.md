[![build-and-publish](https://github.com/David-VTUK/echolife-exporter/actions/workflows/main.yml/badge.svg)](https://github.com/David-VTUK/echolife-exporter/actions/workflows/main.yml)

# Prometheus Exporter for the EchoLife HG612 VDSL Modem

## Requirements

* EchoLife HG612 with unlocked firmware
* The IP address of the modem defined in the environment variable `VDSL_IP`
* Code is packaged as a Docker container automatically via Github CI - https://hub.docker.com/r/virtualthoughts/echolife-exporter