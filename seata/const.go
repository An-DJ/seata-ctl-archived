package seata

const (
	HTTPProtocol                 = "http://"
	HTTPSProtocol                = "https://"
	APIPrefix                    = "/api/v1"
	LoginURL                     = APIPrefix + "/auth/login"
	AdminURL                     = APIPrefix + "/admin"
	HealthCheckURL               = AdminURL + "/status"
	ConfigurationURL             = AdminURL + "/configuration"
	GetConfigurationURL          = AdminURL + "/configuration/get"
	RegistryConfigurationURL     = ConfigurationURL + "/registry"
	ConfigCenterConfigurationURL = ConfigurationURL + "/config-center"
	ReloadConfigurationURL       = ConfigurationURL + "/reload"
	TryTxnURL                    = AdminURL + "/txn"
	TryBeginTxnURL               = TryTxnURL + "/begin"
	TryCommitTxnURL              = TryTxnURL + "/commit"
	TryRollBackTxnURL            = TryTxnURL + "/rollback"
)

const (
	CodeOK = "200"
)

type ConfigType int

const (
	REGISTRY_CONF ConfigType = iota
	CONFIG_CENTER_CONF
	NORMAL_CONFIG
)

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
