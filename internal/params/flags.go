package params

// Flags
const (
	AgentFlag                = "agent"
	AgentFlagUsage           = "Scan origin name"
	DefaultAgent             = "ASTCLI"
	DebugFlag                = "debug"
	DebugUsage               = "Debug mode with detailed logs"
	RetryFlag                = "retry"
	RetryDefault             = 3
	RetryUsage               = "Retry requests to AST on connection failure"
	RetryDelayFlag           = "retry-delay"
	RetryDelayDefault        = 3
	RetryDelayUsage          = "Time between retries in seconds, use with --" + RetryFlag
	SourcesFlag              = "file-source"
	SourcesFlagSh            = "s"
	TenantFlag               = "tenant"
	TenantFlagUsage          = "Checkmarx tenant"
	AsyncFlag                = "async"
	WaitDelayFlag            = "wait-delay"
	ScanTimeoutFlag          = "scan-timeout"
	SourceDirFilterFlag      = "file-filter"
	SourceDirFilterFlagSh    = "f"
	IncludeFilterFlag        = "file-include"
	IncludeFilterFlagSh      = "i"
	ProjectIDFlag            = "project-id"
	BranchFlag               = "branch"
	BranchFlagSh             = "b"
	ScanIDFlag               = "scan-id"
	BranchFlagUsage          = "Branch to scan"
	MainBranchFlag           = "branch"
	ScaResolverFlag          = "sca-resolver"
	ScaResolverParamsFlag    = "sca-resolver-params"
	AccessKeyIDFlag          = "client-id"
	AccessKeySecretFlag      = "client-secret"
	AccessKeyIDFlagUsage     = "The OAuth2 client ID"
	AccessKeySecretFlagUsage = "The OAuth2 client secret"
	InsecureFlag             = "insecure"
	InsecureFlagUsage        = "Ignore TLS certificate validations"
	ScanInfoFormatFlag       = "scan-info-format"
	FormatFlag               = "format"
	FormatFlagUsageFormat    = "Format for the output. One of %s"
	FilterFlag               = "filter"
	BaseURIFlag              = "base-uri"
	ProxyFlag                = "proxy"
	ProxyFlagUsage           = "Proxy server to send communication through"
	ProxyTypeFlag            = "proxy-auth-type"
	ProxyTypeFlagUsage       = "Proxy authentication type, (basic or ntlm)"
	TimeoutFlag              = "timeout"
	TimeoutFlagUsage         = "Timeout for network activity, (default 5 seconds)"
	NtlmProxyDomainFlag      = "proxy-ntlm-domain"
	NtlmProxyDomainFlagUsage = "Window domain when using NTLM proxy"
	BaseURIFlagUsage         = "The base system URI"
	BaseAuthURIFlag          = "base-auth-uri"
	BaseAuthURIFlagUsage     = "The base system IAM URI"
	AstAPIKeyFlag            = "apikey"
	AstAPIKeyUsage           = "The API Key to login to AST"
	ClientRolesFlag          = "roles"
	ClientRolesSh            = "r"
	ClientDescriptionFlag    = "description"
	ClientDescriptionSh      = "d"
	UsernameFlag             = "username"
	UsernameSh               = "u"
	PasswordFlag             = "password"
	PasswordSh               = "p"
	ProfileFlag              = "profile"
	ProfileFlagUsage         = "The default configuration profile"
	Help                     = "help"
	TargetFlag               = "output-name"
	TargetPathFlag           = "output-path"
	TargetFormatFlag         = "report-format"
	ProjectName              = "project-name"
	ScanTypes                = "scan-types"
	ScanTypeFlag             = "scan-type"
	TagList                  = "tags"
	GroupList                = "groups"
	ProjectGroupList         = "project-groups"
	ProjectTagList           = "project-tags"
	IncrementalSast          = "sast-incremental"
	PresetName               = "sast-preset-name"
	Threshold                = "threshold"
	KeyValuePairSize         = 2
	WaitDelayDefault         = 5
	SimilarityIDFlag         = "similarity-id"
	SeverityFlag             = "severity"
	StateFlag                = "state"
	CommentFlag              = "comment"
	SCMTokenFlag             = "token"
	SCMTokenUsage            = "GitHub OAuth token"
	GitHubURLFlag            = "url"
	GitHubURLFlagUsage       = "API base URL"
)

// Parameter values
const (
	IDQueryParam           = "id"
	IDsQueryParam          = "ids"
	IDRegexQueryParam      = "id-regex"
	LimitQueryParam        = "limit"
	OffsetQueryParam       = "offset"
	ScanIDQueryParam       = "scan-id"
	ScanIDsQueryParam      = "scan-ids"
	TagsKeyQueryParam      = "tags-keys"
	TagsValueQueryParam    = "tags-values"
	StatusesQueryParam     = "statuses"
	StatusQueryParam       = "status"
	BranchNameQueryParam   = "branch-name"
	ProjectIDQueryParam    = "project-id"
	FromDateQueryParam     = "from-date"
	ToDateQueryParam       = "to-date"
	SeverityQueryParam     = "severity"
	StateQueryParam        = "state"
	GroupQueryParam        = "group"
	QueryQueryParam        = "query"
	NodeIDsQueryParam      = "node-ids"
	IncludeNodesQueryParam = "include-nodes"
	SortQueryParam         = "sort"
	Profile                = "default"
	BaseURI                = ""
	BaseIAMURI             = ""
	Tenant                 = ""
	Branch                 = ""
)

// Results
const (
	SastType = "sast"
	KicsType = "kics"
	ScaType  = "sca"
)

// ScaAgent AST Role
const ScaAgent = "SCA_AGENT"

var (
	Version = "dev"
)
