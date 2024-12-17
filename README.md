# MAILSERVICE
MAILSERVICE is a > :incoming_envelope: Notification Hub for any stack. 
I decide to use GO, as it is a very lightweight language and the build output image has small footprint and the performance is good
MAILSERVICE consumes mail events from RabbitMQ and send emails over SMTP.
This is still in testing development phase

## Usage

Start worker by running command below

```sh
$ go run ./cmd/mailservice/main.go
```

### Environment variables

| Variable               | Description                          | Required | Default              |
|------------------------|--------------------------------------|----------|----------------------|
| `MAILSERVICE_ENV`       | Environment, reacts on "production"  | *no*     |                      |
| `MAILSERVICE_LOG_LEVEL` | Level of logging                     | *no*     | `debug`              |
| `RABBITMQ_HOST`        | Host of RabbitMQ daemon              | *no*     | `localhost`          |
| `RABBITMQ_PORT`        | Port of RabbitMQ daemon              | *no*     | `5672`               |
| `RABBITMQ_USERNAME`    | RabbitMQ username                    | *no*     | `guest`              |
| `RABBITMQ_PASSWORD`    | RabbitMQ password                    | *no*     | `guest`              |
| `SMTP_PASSWORD`        | Password used for auth to SMTP       | *yes*    |                      |
| `SMTP_PORT`            | Post of SMTP server                  | *no*     | `25`                 |
| `SMTP_HOST`            | Host of SMTP server                  | *no*     | `smtp.sendgrid.net`  |
| `SMTP_USER`            | User used for auth to SMTP           | *no*     | `apikey`             |
| `SENDER_EMAIL`         | Email address of mail sender         | *yes*    |                      |
| `SENDER_NAME `         | Name of mail sender                  | *no*     | `noreply`         |