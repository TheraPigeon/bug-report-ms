TheraPigeon-wide bug reporting microservice. Sends bug reports to the appropriate Slack workspace.

A Docker image is automatically built when code is pushed to the master branch.

### URL

```
https://therapigeonbugreporterms-env.eba-cpajmpju.us-east-1.elasticbeanstalk.com/
```

### Endpoints

## Ping

Pings the service to see if it's up and running.

Request: **GET method**

```
url/ping
```

Response:

```
pong
```

## Soup

Submits the user's bug report to the SOUP workspace on slack.

Request: **POST method**

```
url/soup/?bug=write your bug submission string here
```