package rabbitmq

import (
	"fmt"

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
		fmsg := fmt.Sprintf(msg.Message, msg.Args...)

		r.channel.Publish("", r.queue.Name, false, false, amqp.Publishing{
			AppId:       r.cfg.AppID,
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("[%s]: %s", msg.Level, fmsg)),
		})
	}
}

func (r *RabbitMQLogger) Print(message string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: message,
	}
}

func (r *RabbitMQLogger) Println(message string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: message + "\n",
	}
}

func (r *RabbitMQLogger) Debug(message string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Debug,
		Message: message,
	}
}

func (r *RabbitMQLogger) Info(message string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: message,
	}
}

func (r *RabbitMQLogger) Warn(message string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Warn,
		Message: message,
	}
}

func (r *RabbitMQLogger) Error(message string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Error,
		Message: message,
	}
}

func (r *RabbitMQLogger) Fatal(message string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Fatal,
		Message: message,
	}
}

func (r *RabbitMQLogger) Printf(message string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: message,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Debugf(message string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Debug,
		Message: message,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Infof(message string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: message,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Warnf(message string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Warn,
		Message: message,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Errorf(message string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Error,
		Message: message,
		Args:    args,
	}
}

func (r *RabbitMQLogger) Fatalf(message string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Fatal,
		Message: message,
		Args:    args,
	}
}
