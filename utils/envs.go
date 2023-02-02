package utils

import (
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	logger = ctrl.Log.WithName("utils")
)

func PaymentBaseURL() string {
	return getEnvVar("PAYMENT_BASE_URL")
}

func getEnvVar(name string) string {
	v := os.Getenv(name)
	if v == "" {
		logger.Error(nil, "Environment variable not set", "name", name)
	}
	return v
}
