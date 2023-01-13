docker-compose up -d influxdb grafana
echo "Grafana dashboard at http://localhost:3000/d/k6/k6-load-testing-results"
docker-compose run --rm k6 run /scripts/funcinvoker.js
