# docker-machine-driver-nerdalize

To compile the project : `make`
To use the driver with docker-machine:

`cp dist/docker-machine-driver-nerdalize_<PLATFORM> /usr/local/bin/docker-machine-driver-nerdalize`

Then:
```docker-machine create -d nerdalize \
 --cloudstack-api-url REDACTED \
  --cloudstack-api-key REDACTED \
  --cloudstack-secret-key REDACTED \
  --cloudstack-template "Rancher Nerdalize Edition (1632)" \
  --cloudstack-zone REDACTED \
  --cloudstack-service-offering "Fast Instance" --cloudstack-ssh-user "nerdalize" \
  --cloudstack-expunge --cloudstack-ssh-key-path REDACTED \
  docker-machine-name
```
