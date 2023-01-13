Installation:
1. docker-compose up -d influxdb grafana
2. docker-compose run --rm k6 
3. k6 run /scripts/funcinvoker.js


Information:
In funcinvoker.js, currently data of a function #47658c9cff05caaffd85722b6f81dd163f680f03d4026474a098f966e2f528b3 is being read from a csv file and requests are being sent as per the invocation count pattern every minute. This test is currently set for 15 min.