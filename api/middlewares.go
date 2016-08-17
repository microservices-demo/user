package api

import (
	"time"

	"github.com/go-kit/kit/log"
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

func (mw loggingMiddleware) Login(username, password string) (user users.User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Login",
			"username", username,
			"result", user.UserID,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Login(username, password)
}

func (mw loggingMiddleware) Register(username, password, email string) (status bool) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Register",
			"username", username,
			"email", email,
			"result", status,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Register(username, password, email)
}

func (mw loggingMiddleware) PostUser(user users.User) (status bool) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "PostUser",
			"username", user.Username,
			"email", user.Email,
			"result", status,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.PostUser(user)
}

func (mw loggingMiddleware) GetUsers(id string) (u []users.User, err error) {
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
	return mw.next.GetUsers(id)
}

func (mw loggingMiddleware) PostAddress(add users.Address, id string) (status bool) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "PostAddress",
			"street", add.Street,
			"number", add.Number,
			"user", id,
			"result", status,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.PostAddress(add, id)
}

func (mw loggingMiddleware) GetAddresses(id string) (a []users.Address, err error) {
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
	return mw.next.GetAddresses(id)
}

func (mw loggingMiddleware) PostCard(card users.Card, id string) (status bool) {
	defer func(begin time.Time) {
		cc := card
		cc.MaskCC()
		mw.logger.Log(
			"method", "PostCard",
			"card", cc.LongNum,
			"user", id,
			"result", status,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.PostCard(card, id)
}

func (mw loggingMiddleware) GetCards(id string) (a []users.Card, err error) {
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
	return mw.next.GetCards(id)
}
