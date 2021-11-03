#brew install podman
#podman machine init
#podman machine start
podman build -t pushnotification:latest .
#podman pull busybox:latest
podman container stop pushnotification
podman container rm pushnotification
podman pod rm push-notification
#podman run --network=host imageID dnf -y install java
podman pod create --name push-notification 
#to mock the bckend servers create test pods using busybox or equivelant to execute
#integration tests
#Pass additional arguments to the app if required
podman run \
-d --restart=always --pod=push-notification \
--name=pushnotification pushnotification:latest -p 8080:8080
#podman generate kube push-notification > local-deployment.yaml
#podman play kube local-deployment.yaml