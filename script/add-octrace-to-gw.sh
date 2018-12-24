#!/bin/bash

sed -i 's|"golang.org/x/net/context"|&\
	"go.opencensus.io/exporter/stackdriver/propagation"\
	"go.opencensus.io/trace"|g' ./app/api/proto/api.pb.gw.go

sed -i 's/defer cancel()/&\
		httpFormat := \&propagation.HTTPFormat{}\
		sc, ok := httpFormat.SpanContextFromRequest(req)\
		if ok {\
			sctx, span := trace.StartSpanWithRemoteParent(ctx, "grpc-gateway", sc)\
			defer span.End()\
			ctx = sctx\
		}/g' ./app/api/proto/api.pb.gw.go
