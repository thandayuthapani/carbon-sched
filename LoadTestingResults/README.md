# Collection of Metrics
This file automates the k6 load test and obtains request latency metrics from influxDB.

## Requirements
    requests
    pandas
    influxdb

## Usage
    Ensure that the docker.io/kkyfury/k6azure:v1 container is running.
    python3 startTest.py -u "<funcurl>:80" -n "func-name"

