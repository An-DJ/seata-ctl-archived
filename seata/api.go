package seata

const (
	HTTPProtocol   = "http://"
	HTTPSProtocol  = "https://"
	APIPrefix      = "/api/v1"
	LoginURL       = APIPrefix + "/auth/login"
	AdminURL       = APIPrefix + "/admin"
	HealthCheckURL = AdminURL + "/status"
	RegistryURL = AdminURL + "/registry"
	ConfigCenterURL        = AdminURL + "/config-center"
	ConfigurationURL       = AdminURL + "/configuration"
	ReloadConfigurationURL = AdminURL + "/reload"
	TryTxnURL              = AdminURL + "/txn"
)

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
