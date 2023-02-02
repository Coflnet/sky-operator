package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	ctrl "sigs.k8s.io/controller-runtime"
  k8smetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	api "github.com/Coflnet/sky-operator/target/dir"
	"github.com/Coflnet/sky-operator/utils"
)

var (
	// PreApi Gauge
	PreAPIGauge = promauto.NewGauge(prometheus.GaugeOpts{Name: "sky_controller_pre_api_useres", Help: "Pre API Users"})

	logger = ctrl.Log.WithName("metrics")
)

func Start() {
	initialize()
	go func() {
		for {
			loop()
			time.Sleep(30 * time.Second)
		}
	}()
}

func initialize() {
  k8smetrics.Registry.MustRegister(PreAPIGauge)
}

func loop() {
	go updatePreAPIGauge()
}

func updatePreAPIGauge() {

	client, err := api.NewClient(utils.PaymentBaseURL())
	if err != nil {
		logger.Error(err, "Could not create payment client")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := client.ProductsServiceServiceSlugCountGet(ctx, api.ProductsServiceServiceSlugCountGetParams{ServiceSlug: "pre_api"})

	if err != nil {
		logger.Error(err, "Error while getting pre api count")
		return
	}

	PreAPIGauge.Set(float64(count))
}
