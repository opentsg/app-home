module github.com/opentsg/app-home

go 1.24.2

require github.com/mrmxf/clog v0.7.21

require (
	github.com/go-chi/chi/v5 v5.2.1 // indirect
	github.com/samber/slog-chi v1.15.0 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
)

// internal dev only
// replace github.com/mrmxf/clog => ../../0_mrmxf/clog
