package ak

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/getsentry/sentry-go"
	httptransport "github.com/go-openapi/runtime/client"
	log "github.com/sirupsen/logrus"
	"goauthentik.io/api"
	"goauthentik.io/internal/constants"
)

func doGlobalSetup(outpost api.Outpost, globalConfig api.Config) {
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyMsg:  "event",
			log.FieldKeyTime: "timestamp",
		},
	})
	switch outpost.Config[ConfigLogLevel].(string) {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
	log.WithField("logger", "authentik.outpost").WithField("hash", constants.BUILD()).WithField("version", constants.VERSION).Info("Starting authentik outpost")

	if globalConfig.ErrorReporting.Enabled {
		sentryEnv := fmt.Sprintf("%s-outpost-%s", globalConfig.ErrorReporting.Environment, outpost.Type)
		dsn := "https://a579bb09306d4f8b8d8847c052d3a1d3@sentry.beryju.org/8"
		log.WithField("env", sentryEnv).Debug("Error reporting enabled")
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              dsn,
			Environment:      sentryEnv,
			TracesSampleRate: float64(globalConfig.ErrorReporting.TracesSampleRate),
		})
		if err != nil {
			log.WithField("env", sentryEnv).WithError(err).Warning("Failed to initialise sentry")
		}
	}
}

// GetTLSTransport Get a TLS transport instance, that skips verification if configured via environment variables.
func GetTLSTransport() http.RoundTripper {
	value, set := os.LookupEnv("AUTHENTIK_INSECURE")
	if !set {
		value = "false"
	}
	tlsTransport, err := httptransport.TLSTransport(httptransport.TLSClientOptions{
		InsecureSkipVerify: strings.ToLower(value) == "true",
	})
	if err != nil {
		panic(err)
	}
	return tlsTransport
}
