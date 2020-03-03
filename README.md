# nats-listener

> A NATS client that can be configured to listen to subjects and output incoming messages.

## Environment Variables

| Variable | Description | Default |
|:---:|:---|:---:|
| `NATS_SERVER` |The NATS address of the server you're connecting to. | `nats://127.0.0.1:4222` |
| `NATS_SUBJECT` | The subject you'd like to listen in on. | `>` |
