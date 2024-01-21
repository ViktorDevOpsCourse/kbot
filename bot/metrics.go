package bot

import (
	"context"
	"go.opentelemetry.io/otel"
)

func pmetrics(ctx context.Context) {
	meter := otel.GetMeterProvider().Meter("kbot_joke_counter")

	counter, _ := meter.Int64Counter("joke")

	counter.Add(ctx, 1)
}
