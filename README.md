# go-ping

A simple Go server/microservice example for [Docker's Go Language Guide](https://docs.docker.com/language/golang/).

This app has helm charts that can be used to deploy into K8 or Openshift environments. The image is configured to be pulled from a public repository

Notable features:

* Uses helm charts that can be used for deployment
* Includes a [multi-stage `Dockerfile`](https://github.com/olliefr/docker-gs-ping/blob/main/Dockerfile.multistage), which actually is a good example of how to build Go binaries _for production releases_.

## License

[Apache-2.0 License](LICENSE)
