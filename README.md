# stori-challenge
## Author: Erik Gomez

### Testing
```
curl --location 'https://yjizyf69i6.execute-api.us-east-1.amazonaws.com/stori-challenge/report' \
--header 'Content-Type: application/json' \
--data-raw '{
  "destination_email_address": "your.email.address@gmail.com"
}'
```

### For further iterations
- Improve documentation, specially explain how AWS were configured.
- Use AWS RDS to store transactions, write a local implementation too.
- Change configuration instantiation since it is being instantiated every time the lambda function is called. Also add a refresh endpoint so there is no need to restart the application to impact the new configuration.
- Improve logging, differentiating between different criticality and add a middleware to handle it.
- Improve error handling, add a middleware.
- Add a local implementation for every repository.
- Use AWS Config and AWS Secrets to store config variables instead of getting from environment variables.
- Improve response message.
- Fix some methods that are not using the context, it will probably be fixed when fixing logs and error handling.
- Improve test coverage
