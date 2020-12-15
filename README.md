TheraPigeon-wide bug reporting microservice. Sends bug reports to the appropriate Slack workspace.

A Docker image is automatically built when code is pushed to the master branch.

### Endpoints

## Ping

Pings the service to see if it's up and running.

Request: 

```
url/ping
```

Response:

```
pong
```

## Soup

Submits a bug report to the SOUP workspace on slack.

Request:

```
url/soup?bug=write your bug submission string here
```