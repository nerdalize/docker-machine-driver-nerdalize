# docker-machine-driver-nerdalize

To compile the project : `make`

To use the driver with docker-machine:

```bash
$ cp dist/docker-machine-driver-nerdalize_<PLATFORM> /usr/local/bin/docker-machine-driver-nerdalize
```

Then:
``` bash
$ docker-machine create -d nerdalize \
 --nerdalize-api-url REDACTED \
  --nerdalize-api-key REDACTED \
  --nerdalize-secret-key REDACTED \
  --nerdalize-template "Rancher Nerdalize Edition (1632)" \
  --nerdalize-zone REDACTED \
  --nerdalize-service-offering "Fast Instance" --nerdalize-ssh-user "nerdalize" \
  --nerdalize-expunge  \
  docker-machine-name
```
