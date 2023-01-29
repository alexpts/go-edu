package middleware

import pm "github.com/alexpts/edu-go/internal/middleware/panic"

// aliases / re-export

type PanicMiddleware = pm.MiddlewarePanic

var (
	ProvideMiddlewarePanic = pm.ProvideMiddlewarePanic
)
