package handlers

import (
	"gofirstapp/internal/client"
	"net/http"
)

type Handlers struct {
	client client.Client
}

type HandlersOptions struct {
	client client.Client
}

type WithOptions = func(o *HandlersOptions) *HandlersOptions

func WithCustomClient(c client.Client) WithOptions {
	return func(o *HandlersOptions) *HandlersOptions {
		o.client = c
		return o
	}
}

func NewHandlers(options ...WithOptions) *Handlers {
	opts := &HandlersOptions{
		client: http.DefaultClient,
	}

	for _, o := range options {
		opts = o(opts)
	}

	return &Handlers{
		client: opts.client,
	}
}