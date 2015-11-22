#docker build -t safeie/bayesian-classifier .
docker run --name classifier -p 0.0.0.0:8812:8812 -d safeie/bayesian-classifier