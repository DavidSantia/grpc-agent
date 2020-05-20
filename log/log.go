package log

import (
	"context"
	"github.com/newrelic/go-agent/v3/integrations/logcontext/nrlogrusplugin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

func InitLog() {
	logger := logrus.New()
	logger.SetFormatter(nrlogrusplugin.ContextFormatter{})
	logger.SetLevel(logrus.TraceLevel)
}

func DoLog(txn *newrelic.Transaction, level int32, message, host string) {
	logger := logrus.New()
	logger.SetFormatter(nrlogrusplugin.ContextFormatter{})

	ctx := newrelic.NewContext(context.Background(), txn)
	logger.WithContext(ctx).WithField("client", host).Log(logrus.Level(level), message)
}
