# trackem - a hack project
Playing with new tech deserves a playground application, preferably one with actual real world value. In this particular case, I wanted to experiment with [up](https://github.com/apex/up), [dynamodb](https://aws.amazon.com/dynamodb), and others. Rather than build another [shrty](https://github.com/bkono/go-shrty) in go, I decided to tackle a basic tracking pixel service.

## Features
With a goal of simplicity, the core features of this tool are:

- Client side link generation, with all metadata encoded in the provided URL
- Up-ready handler that records visits to the link in dynamodb
- More to come, see Next Steps

## Usage
Coming soon. Flag help docs are accurate for now.

### Next steps:

- [x] Add server cmd
- [x] Listen for url, serve static transparent gif
- [x] Decode url query param in dynamodb
- [ ] Move to one command, one main.go with `server` and `create` as args
- [ ] Deploy on Up
- [ ] Add dashboard, SES and/or SNS alerts
- [ ] Consider additional AWS tech to dabble in

