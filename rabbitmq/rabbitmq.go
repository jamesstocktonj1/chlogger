package rabbitmq

import (
	"github.com/jamesstocktonj1/chlogger"
	amqp "github.com/rabbitmq/amqp091-go"
)

var _ chlogger.Logger = &RabbitMQLogger{}

// RabbitMQLogger is a logger implementation that logs messages to RabbitMQ
type RabbitMQLogger struct {
	cfg     RabbitMQConfig
	msg     chan chlogger.Message
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

// RabbitMQConfig is the configuration for a RabbitMQLogger
type RabbitMQConfig struct {
	Host  string
	Topic string
	AppID string
}

// NewLogger creates a new RabbitMQLogger
func NewLogger(cfg RabbitMQConfig) *RabbitMQLogger {
	return &RabbitMQLogger{
		cfg: cfg,
		msg: make(chan chlogger.Message),
	}
}

func (r *RabbitMQLogger) Init() error {
	var err error

	r.conn, err = amqp.Dial(r.cfg.Host)
	if err != nil {
		return err
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		return err
	}

	r.queue, err = r.channel.QueueDeclare(r.cfg.Topic, false, false, false, false, nil)
	if err != nil {
		return err
	}

	go r.daemon()

	return nil
}

func (r *RabbitMQLogger) Close() error {
	close(r.msg)
	r.channel.Close()
	r.conn.Close()
	return nil
}

func (r *RabbitMQLogger) daemon() {
	for msg := range r.msg {
		r.channel.Publish("", r.queue.Name, false, false, amqp.Publishing{
			AppId:       r.cfg.AppID,
			ContentType: "text/plain",
			Body:        []byte(msg.String()),
		})
	}
}

func (r *RabbitMQLogger) Print(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg,
	}
}

func (r *RabbitMQLogger) Println(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg + "\n",
	}
}

func (r *RabbitMQLogger) Debug(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Debug,
		Message: msg,
	}
}

func (r *RabbitMQLogger) Info(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg,
	}
}

func (r *RabbitMQLogger) Warn(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Warn,
		Message: msg,
	}
}

func (r *RabbitMQLogger) Error(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Error,
		Message: msg,
	}
}

func (r *RabbitMQLogger) Fatal(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Fatal,
		Message: msg,
	}
}

func (r *RabbitMQLogger) Printf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Debugf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Debug,
		Message: msg,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Infof(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Warnf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Warn,
		Message: msg,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Errorf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Error,
		Message: msg,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Fatalf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Fatal,
		Message: msg,
		Args:    args,
	}
}
