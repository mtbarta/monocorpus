package service

import (
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// LoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
// func LoggingMiddleware(logger log.Logger) Middleware {
// 	return func(next Service) Service {
// 		return loggingMiddleware{logger, next}
// 	}
// }

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

// func (mw loggingMiddleware) Login(ctx context.Context, user string, password string) (token string, err error) {
// 	defer func() {
// 		mw.logger.Log("method", "Login", "user", user, "err", err)
// 	}()
// 	return mw.next.Login(ctx, user, password)
// }

// func (mw loggingMiddleware) CreateUser(ctx context.Context, first string, last string, email string, password string) (v string, err error) {
// 	defer func() {
// 		mw.logger.Log("method", "CreateUser", "Email", email, "err", err)
// 	}()
// 	return mw.next.CreateUser(ctx, first, last, email, password)
// }

// InstrumentingMiddleware returns a service middleware that instruments
// the number of integers summed and characters concatenated over the lifetime of
// the service.
// func InstrumentingMiddleware(ints, chars metrics.Counter) Middleware {
// 	return func(next Service) Service {
// 		return instrumentingMiddleware{
// 			ints:  ints,
// 			chars: chars,
// 			next:  next,
// 		}
// 	}
// }

// type instrumentingMiddleware struct {
// 	getUser    metrics.Counter
// 	createUser metrics.Counter
// 	next       Service
// }

// func (mw instrumentingMiddleware) GetUser(ctx context.Context, a string) (int, error) {
// 	v, err := mw.next.GetUser(ctx, a)
// 	mw.getUser.Add(float64(1.0))
// 	return v, err
// }

// func (mw instrumentingMiddleware) CreateUser(ctx context.Context, a string) (string, error) {
// 	v, err := mw.next.CreateUser(ctx, a)
// 	mw.createUser.Add(float64(len(1.0)))
// 	return v, err
// }
