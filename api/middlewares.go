package api

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/microservices-demo/user/users"
)

// Middleware decorates a service.
type Middleware func(Service) Service

// LoggingMiddleware logs method calls, parameters, results, and elapsed time.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) Login(ctx context.Context, username, password string) (user users.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Login",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Login(ctx, username, password)
}

func (mw loggingMiddleware) Register(ctx context.Context, username, password, email, first, last string) (string, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Register",
			"username", username,
			"email", email,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Register(ctx, username, password, email, first, last)
}

func (mw loggingMiddleware) PostUser(ctx context.Context, user users.User) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "PostUser",
			"username", user.Username,
			"email", user.Email,
			"result", id,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.PostUser(ctx, user)
}

func (mw loggingMiddleware) GetUsers(ctx context.Context, id string) (u []users.User, err error) {
	defer func(begin time.Time) {
		who := id
		if who == "" {
			who = "all"
		}
		mw.logger.Log(
			"method", "GetUsers",
			"id", who,
			"result", len(u),
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetUsers(ctx, id)
}

func (mw loggingMiddleware) PostAddress(ctx context.Context, add users.Address, id string) (string, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "PostAddress",
			"street", add.Street,
			"number", add.Number,
			"user", id,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.PostAddress(ctx, add, id)
}

func (mw loggingMiddleware) GetAddresses(ctx context.Context, id string) (a []users.Address, err error) {
	defer func(begin time.Time) {
		who := id
		if who == "" {
			who = "all"
		}
		mw.logger.Log(
			"method", "GetAddresses",
			"id", who,
			"result", len(a),
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetAddresses(ctx, id)
}

func (mw loggingMiddleware) PostCard(ctx context.Context, card users.Card, id string) (string, error) {
	defer func(begin time.Time) {
		cc := card
		cc.MaskCC()
		mw.logger.Log(
			"method", "PostCard",
			"card", cc.LongNum,
			"user", id,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.PostCard(ctx, card, id)
}

func (mw loggingMiddleware) GetCards(ctx context.Context, id string) (a []users.Card, err error) {
	defer func(begin time.Time) {
		who := id
		if who == "" {
			who = "all"
		}
		mw.logger.Log(
			"method", "GetCards",
			"id", who,
			"result", len(a),
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetCards(ctx, id)
}

func (mw loggingMiddleware) Delete(ctx context.Context, entity, id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Delete",
			"entity", entity,
			"id", id,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Delete(ctx, entity, id)
}

func (mw loggingMiddleware) Health() (health []Health) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Health",
			"result", len(health),
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Health()
}

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(requestCount metrics.Counter, requestLatency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		Service:        s,
	}
}

func (s *instrumentingService) Login(ctx context.Context, username, password string) (users.User, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "login").Add(1)
		s.requestLatency.With("method", "login").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Login(ctx, username, password)
}

func (s *instrumentingService) Register(ctx context.Context, username, password, email, first, last string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "register").Add(1)
		s.requestLatency.With("method", "register").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Register(ctx, username, password, email, first, last)
}

func (s *instrumentingService) PostUser(ctx context.Context, user users.User) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "postUser").Add(1)
		s.requestLatency.With("method", "postUser").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PostUser(ctx, user)
}

func (s *instrumentingService) GetUsers(ctx context.Context, id string) (u []users.User, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "getUsers").Add(1)
		s.requestLatency.With("method", "getUsers").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetUsers(ctx, id)
}

func (s *instrumentingService) PostAddress(ctx context.Context, add users.Address, id string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "postAddress").Add(1)
		s.requestLatency.With("method", "postAddress").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PostAddress(ctx, add, id)
}

func (s *instrumentingService) GetAddresses(ctx context.Context, id string) ([]users.Address, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "getAddresses").Add(1)
		s.requestLatency.With("method", "getAddresses").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetAddresses(ctx, id)
}

func (s *instrumentingService) PostCard(ctx context.Context, card users.Card, id string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "postCard").Add(1)
		s.requestLatency.With("method", "postCard").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PostCard(ctx, card, id)
}

func (s *instrumentingService) GetCards(ctx context.Context, id string) ([]users.Card, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "getCards").Add(1)
		s.requestLatency.With("method", "getCards").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetCards(ctx, id)
}

func (s *instrumentingService) Delete(ctx context.Context, entity, id string) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "delete").Add(1)
		s.requestLatency.With("method", "delete").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Delete(ctx, entity, id)
}

func (s *instrumentingService) Health() []Health {
	defer func(begin time.Time) {
		s.requestCount.With("method", "health").Add(1)
		s.requestLatency.With("method", "health").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Health()
}
