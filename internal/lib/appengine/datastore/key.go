package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/110y/gae-go-grpc/internal/lib/appengine/internal"
)

func NameKey(ctx context.Context, kind, name string, parent *datastore.Key) *datastore.Key {
	return &datastore.Key{
		Kind:      kind,
		Name:      name,
		Parent:    parent,
		Namespace: internal.GetAppEngineNamespace(ctx),
	}
}
