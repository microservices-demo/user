package db

import (
	"context"
	"fmt"
	"os"

	"github.com/microservices-demo/user/users"

	otext "github.com/opentracing/opentracing-go/ext"
	stdopentracing "github.com/opentracing/opentracing-go"
)

type TracingMiddleware struct {
	next   Database
}

// Middleware decorates a database.
type Middleware func(Database) TracingMiddleware

// DbTracingMiddleware traces database calls.
func DbTracingMiddleware() Middleware {
	return func(next Database) TracingMiddleware {
		return TracingMiddleware{
			next:   next,
		}
	}
}

func (mw TracingMiddleware) Init() error {
	return mw.next.Init()
}

func (mw TracingMiddleware) CreateAddress(ctx context.Context, a *users.Address, userid string) error {
	span := startSpan(ctx, "create address on db")
	err := mw.next.CreateAddress(a, userid)
	finishSpan(span, len(fmt.Sprintf("%#v", a)))
	return err
}

func (mw TracingMiddleware) CreateCard(ctx context.Context, c *users.Card, userid string) error {
	span := startSpan(ctx, "create card on db")
	err := mw.next.CreateCard(c, userid)
	finishSpan(span, len(fmt.Sprintf("%#v", c)))
	return err
}

func (mw TracingMiddleware) CreateUser(ctx context.Context, u *users.User) error {
	span := startSpan(ctx, "create user on db")
	err := mw.next.CreateUser(u)
	finishSpan(span, len(fmt.Sprintf("%#v", u)))
	return err
}

func (mw TracingMiddleware) Delete(ctx context.Context, entity, id string) error {
	span := startSpan(ctx, "delete from db")
	err := mw.next.Delete(entity, id)
	finishSpan(span, len(fmt.Sprintf("%#v", entity)))
	return err
}

func (mw TracingMiddleware) GetAddress(ctx context.Context, n string) (users.Address, error) {
	span := startSpan(ctx, "address from db")
	a, err := mw.next.GetAddress(n)
	finishSpan(span, len(fmt.Sprintf("%#v", a)))
	return a, err
}

func (mw TracingMiddleware) GetAddresses(ctx context.Context) ([]users.Address, error) {
	span := startSpan(ctx, "addresses from db")
	a, err := mw.next.GetAddresses()
	finishSpan(span, len(fmt.Sprintf("%#v", a)))
	return a, err
}

func (mw TracingMiddleware) GetCard(ctx context.Context, n string) (users.Card, error) {
	span := startSpan(ctx, "card from db")
	c, err := mw.next.GetCard(n)
	finishSpan(span, len(fmt.Sprintf("%#v", c)))
	return c, err
}

func (mw TracingMiddleware) GetCards(ctx context.Context) ([]users.Card, error) {
	span := startSpan(ctx, "cards from db")
	c, err := mw.next.GetCards()
	finishSpan(span, len(fmt.Sprintf("%#v", c)))
	return c, err
}


func (mw TracingMiddleware) GetUserByName(ctx context.Context, n string) (users.User, error) {
	span := startSpan(ctx, "user from db")
	u, err := mw.next.GetUserByName(n)
	finishSpan(span, len(fmt.Sprintf("%#v", u)))
	return u, err
}

func (mw TracingMiddleware) GetUser(ctx context.Context, n string) (users.User, error) {
	span := startSpan(ctx, "user from db")
	u, err := mw.next.GetUser(n)
	finishSpan(span, len(fmt.Sprintf("%#v", u)))
	return u, err
}

func (mw TracingMiddleware) GetUsers(ctx context.Context) ([]users.User, error) {
	span := startSpan(ctx, "users from db")
	us, err := mw.next.GetUsers()
	finishSpan(span, len(fmt.Sprintf("%#v", us)))
	return us, err
}

func (mw TracingMiddleware) GetUserAddresses(ctx context.Context, u *users.User) error {
	span := startSpan(ctx, "user addresses from db")
	err := mw.next.GetUserAddresses(u)
	finishSpan(span, len(fmt.Sprintf("%#v", u.Addresses)))
	return err
}

func (mw TracingMiddleware) GetUserCards(ctx context.Context, u *users.User) error {
	span := startSpan(ctx, "user cards from db")
	err := mw.next.GetUserCards(u)
	finishSpan(span, len(fmt.Sprintf("%#v", u.Cards)))
	return err
}

func (mw TracingMiddleware) Ping() error {
	return mw.next.Ping()
}

func startSpan(ctx context.Context, n string) stdopentracing.Span {
	var span stdopentracing.Span
	span, ctx = stdopentracing.StartSpanFromContext(ctx, n)
	otext.SpanKindRPCClient.Set(span)
	span.SetTag("db.type", os.Getenv("USER_DATABASE"))
	span.SetTag("peer.address", os.Getenv("MONGO_HOST"))
	return span
}

func finishSpan(span stdopentracing.Span, size int) {
	span.SetTag("db.query.result.size", size)
	span.Finish()
}
