package tracinghook

import (
	"github.com/rs/zerolog"
)

type TracingHook struct{}

func (h TracingHook) Run(event *zerolog.Event, level zerolog.Level, msg string) {
	ctx := event.GetCtx()

	if traceId := ctx.Value("traceId"); traceId != nil {
		event.Str("traceId", traceId.(string))
	}
}
