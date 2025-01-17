package consumer

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/go-yaml/yaml"

	"mailservice/internal/config"
	"mailservice/internal/log"
	"mailservice/pkg/amqp"
	"mailservice/pkg/env"
	"mailservice/pkg/eventapi"
)

// EmailConfirmationURI returns unique URL for user to confirm his identity.
func EmailConfirmationURI(token string) string {
	url := env.FetchDefault("CONFIRM_URL", "http://example.com/#{}")
	return strings.Replace(url, "#{}", token, 1)
}

// ResetPasswordURI returns unique URL for user to reset password.
func ResetPasswordURI(token string) string {
	url := env.FetchDefault("RESET_URL", "http://example.com/#{}")
	return strings.Replace(url, "#{}", token, 1)
}

func amqpURI() string {
	host := env.FetchDefault("RABBITMQ_HOST", "localhost")
	port := env.FetchDefault("RABBITMQ_PORT", "5672")
	username := env.FetchDefault("RABBITMQ_USERNAME", "guest")
	password := env.FetchDefault("RABBITMQ_PASSWORD", "guest")

	return fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
}

func requireEnvs() {
	env.Must(env.Fetch("SMTP_PASSWORD"))
	env.Must(env.Fetch("SENDER_EMAIL"))
}

// Run starts the application.
func Run(path, tag string) {
	requireEnvs()
	conf := new(config.Config)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err).
			Msgf("can not read file %s", path)
	}

	if err := yaml.Unmarshal([]byte(content), &conf); err != nil {
		log.Fatal().Err(err).
			Msgf("can not unmarshal configuration %s", path)
	}

	if err := conf.Validate(); err != nil {
		log.Fatal().Err(err).
			Msgf("configuration file %s is not valid", path)
	}

	serveMux := amqp.NewServeMux(amqpURI(), tag, conf.Exchanges, conf.Keychain)

	for id := range conf.Events {
		eventConf := conf.Events[id]
		serveMux.HandleFunc(eventConf.Key, eventConf.Exchange, func(raw eventapi.RawEvent) {
			log.Info().Msgf("processing event %s", eventConf.Key)

			event, err := eventapi.Unmarshal(raw)
			if err != nil {
				log.Error().
					Err(err).
					Fields(raw).
					Msg("can not unmarshal event")
				return
			}

			record, err := event.FixAndValidate(conf.Languages[0].Code)
			if err != nil {
				log.Error().Err(err).Fields(raw).Msg("event is not valid")
				return
			}

			log.Info().Str("uid", record.User.UID).Str("email", record.User.Email).Msgf("event received")

			// Checks, that language is supported.
			if !conf.ContainsLanguage(record.Language) {
				log.Error().
					Str("language", record.Language).
					Msg("language is not supported")
				return
			}

			if strings.TrimSpace(eventConf.Expression) != "" {
				result, err := expr.Eval(eventConf.Expression, raw)
				if err != nil {
					log.Error().Err(err).Msg("expression evaluation failed")
				}

				match, ok := result.(bool)
				if !ok {
					log.Error().Err(err).Msg("expression result is not boolean")
					return
				}

				log := log.Info().
					Str("uid", record.User.UID).
					Str("email", record.User.Email).
					Interface("match", result)

				if !match {
					log.Msgf("skipped")
					return
				}

				log.Msgf("matched")
			}

			switch eventConf.Key {
			case "user.email.confirmation.token":
				raw["EmailConfirmationURI"] = EmailConfirmationURI(record.Token)
			case "user.password.reset.token":
				raw["ResetPasswordURI"] = ResetPasswordURI(record.Token)
			}

			// log.Info().Str("ecb",raw["record"].(string)).Msg("raw")
			tpl := eventConf.Template(record.Language)
			content, err := tpl.Content(raw)
			if err != nil {
				log.Error().Err(err).Msg("template execution failed")
				return
			}

			email := Email{
				FromAddress: env.Must(env.Fetch("SENDER_EMAIL")),
				FromName:    env.FetchDefault("SENDER_NAME", "mailservice"),
				ToAddress:   record.User.Email,
				Subject:     tpl.Subject,
				Reader:      bytes.NewReader(content),
			}

			password := env.Must(env.Fetch("SMTP_PASSWORD"))
			conf := SMTPConf{
				Host:     env.FetchDefault("SMTP_HOST", "smtp.sendgrid.net"),
				Port:     env.FetchDefault("SMTP_PORT", "25"),
				Username: env.FetchDefault("SMTP_USER", "apikey"),
				Password: password,
			}

			if err := NewEmailSender(conf, email).Send(); err != nil {
				log.Error().Err(err).Msg("failed to send email")
				return
			}
		})
	}

	log.Info().Msg("waiting for events")
	if err := serveMux.ListenAndServe(); err != nil {
		log.Panic().Err(err).Msg("connection failed")
	}
}
