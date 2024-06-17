package redis

import (
	"context"

	"github.com/jamesstocktonj1/chlogger"
	"github.com/redis/go-redis/v9"
)

var _ chlogger.Logger = &RedisLogger{}

// RedisLogger is a logger implementation that logs messages to Redis
type RedisLogger struct {
	cfg RedisConfig
	msg chan chlogger.Message
	db  *redis.Client
	ctx context.Context
}

// RedisConfig is the configuration for a RedisLogger
type RedisConfig struct {
	Host     string
	Username string
	Password string
	Topic    string
}

// NewLogger creates a new RedisLogger
func NewLogger(cfg RedisConfig) *RedisLogger {
	return &RedisLogger{
		cfg: cfg,
		msg: make(chan chlogger.Message),
		ctx: context.Background(),
	}
}

func (r *RedisLogger) Init() error {
	r.db = redis.NewClient(&redis.Options{
		Addr:     r.cfg.Host,
		Username: r.cfg.Username,
		Password: r.cfg.Password,
	})

	err := r.db.Ping(r.ctx).Err()
	if err != nil {
		return err
	}

	go r.daemon()

	return nil
}

func (r *RedisLogger) Close() error {
	close(r.msg)
	r.db.Close()
	return nil
}

func (r *RedisLogger) daemon() {
	for msg := range r.msg {
		r.db.RPush(r.ctx, r.cfg.Topic, msg.String()).Err()
	}
}

func (r *RedisLogger) Print(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg,
	}
}

func (r *RedisLogger) Println(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg + "\n",
	}
}

func (r *RedisLogger) Debug(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Debug,
		Message: msg,
	}
}

func (r *RedisLogger) Info(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg,
	}
}

func (r *RedisLogger) Warn(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Warn,
		Message: msg,
	}
}

func (r *RedisLogger) Error(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Error,
		Message: msg,
	}
}

func (r *RedisLogger) Fatal(msg string) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Fatal,
		Message: msg,
	}
}

func (r *RedisLogger) Printf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg,
		Args:    args,
	}
}

func (r *RedisLogger) Debugf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Debug,
		Message: msg,
		Args:    args,
	}
}

func (r *RedisLogger) Infof(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Info,
		Message: msg,
		Args:    args,
	}
}

func (r *RedisLogger) Warnf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Warn,
		Message: msg,
		Args:    args,
	}
}

func (r *RedisLogger) Errorf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Error,
		Message: msg,
		Args:    args,
	}
}

func (r *RedisLogger) Fatalf(msg string, args ...interface{}) {
	r.msg <- chlogger.Message{
		Level:   chlogger.Fatal,
		Message: msg,
		Args:    args,
	}
}
