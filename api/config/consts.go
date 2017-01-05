package config

// Key for CheckBundleConfig options
type Key string

// Constants per type as defined in
// https://login.circonus.com/resources/api/calls/check_bundle
const (
	DefaultCIDRegex  = "[0-9]+"
	DefaultUUIDRegex = "[[:xdigit:]]{8}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{12}"
	// Endpoint prefix & cid regex
	AccountPrefix              = "/account"
	AccountCIDRegex            = "^" + AccountPrefix + "/([0-9]+|current)$"
	AcknowledgementPrefix      = "/acknowledgement"
	AcknowledgementCIDRegex    = "^" + AcknowledgementPrefix + "/" + DefaultCIDRegex + "$"
	AlertPrefix                = "/alert"
	AlertCIDRegex              = "^" + AlertPrefix + "/" + DefaultCIDRegex + "$"
	AnnotationPrefix           = "/annotation"
	AnnotationCIDRegex         = "^" + AnnotationPrefix + "/" + DefaultCIDRegex + "$"
	BrokerPrefix               = "/broker"
	BrokerCIDRegex             = "^" + BrokerPrefix + "/" + DefaultCIDRegex + "$"
	CheckBundleMetricsPrefix   = "/check_bundle_metrics"
	CheckBundleMetricsCIDRegex = "^" + CheckBundleMetricsPrefix + "/" + DefaultCIDRegex + "$"
	CheckBundlePrefix          = "/check_bundle"
	CheckBundleCIDRegex        = "^" + CheckBundlePrefix + "/" + DefaultCIDRegex + "$"
	CheckPrefix                = "/check"
	CheckCIDRegex              = "^" + CheckPrefix + "/" + DefaultCIDRegex + "$"
	ContactGroupPrefix         = "/contact_group"
	ContactGroupCIDRegex       = "^" + ContactGroupPrefix + "/" + DefaultCIDRegex + "$"
	DashboardPrefix            = "/dashboard"
	DashboardCIDRegex          = "^" + DashboardPrefix + "/" + DefaultCIDRegex + "$"
	GraphPrefix                = "/graph"
	GraphCIDRegex              = "^" + GraphPrefix + "/" + DefaultUUIDRegex + "$"
	MaintenancePrefix          = "/maintenance"
	MaintenanceCIDRegex        = "^" + MaintenancePrefix + "/" + DefaultCIDRegex + "$"
	MetricClusterPrefix        = "/metric_cluster"
	MetricClusterCIDRegex      = "^" + MetricClusterPrefix + "/" + DefaultCIDRegex + "$"
	MetricPrefix               = "/metric"
	MetricCIDRegex             = "^" + MetricPrefix + "/[0-9]+_[^[:space:]]+$"
	OutlierReportPrefix        = "/outlier_report"
	OutlierReportCIDRegex      = "^" + OutlierReportPrefix + "/" + DefaultCIDRegex + "$"
	ProvisionBrokerPrefix      = "/provision_broker"
	ProvisionBrokerCIDRegex    = "^" + ProvisionBrokerPrefix + "/[a-z0-9]+-[a-z0-9]+$"
	RuleSetGroupPrefix         = "/rule_set_group"
	RulesetGroupCIDRegex       = "^" + RuleSetGroupPrefix + "/" + DefaultCIDRegex + "$"
	RuleSetPrefix              = "/rule_set"
	RuleSetCIDRegex            = "^" + RuleSetPrefix + "/[0-9]+_.+$"
	UserPrefix                 = "/user"
	UserCIDRegex               = "^" + UserPrefix + "/([0-9]+|current)$"
	WorksheetPrefix            = "/worksheet"
	WorksheetCIDRegex          = "^" + WorksheetPrefix + "/" + DefaultUUIDRegex + "$"

	NumSeverityLevels = 5

	//
	// default settings for api.NewCheckBundle()
	//
	DefaultCheckBundleMetricLimit = -1 // unlimited
	DefaultCheckBundleStatus      = "active"
	DefaultCheckBundlePeriod      = 60
	DefaultCheckBundleTimeout     = 10

	//
	// common (apply to more than one check type)
	//
	AsyncMetrics = Key("async_metrics")

	//
	// httptrap
	//
	SecretKey = Key("secret")

	//
	// "http"
	//
	AuthMethod   = Key("auth_method")
	AuthPassword = Key("auth_password")
	AuthUser     = Key("auth_user")
	Body         = Key("body")
	CAChain      = Key("ca_chain")
	CertFile     = Key("certificate_file")
	Ciphers      = Key("ciphers")
	Code         = Key("code")
	Extract      = Key("extract")
	// HeaderPrefix is special because the actual key is dynamic and matches:
	// `header_(\S+)`
	HeaderPrefix = Key("header_")
	HTTPVersion  = Key("http_version")
	KeyFile      = Key("key_file")
	Method       = Key("method")
	Payload      = Key("payload")
	ReadLimit    = Key("read_limit")
	Redirects    = Key("redirects")
	URL          = Key("url")

	//
	// reserved - config option(s) can't actually be set - here for r/o access
	//
	ReverseSecretKey = Key("reverse:secret_key")
	SubmissionURL    = Key("submission_url")
)
