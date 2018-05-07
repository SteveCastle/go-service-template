# Shrike Micro-service Framework

## Description

A simple micro-service framework for quickly deploying docker containers with a Go web server for personal projects.

## Build

1.  `docker build -t shrike:latest .`

## Run

1.  `docker run -p 1323:1323 -v {absolute-path-to-static-dir}:/docker/bin/static shrike:latest`
