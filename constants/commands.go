package constants

//  Constants for Actions
const (
	Read           = "read"
	Create         = "create"
	Update         = "update"
	Rollback       = "rollback"
	Edit           = "edit"
	Delete         = "delete"
	Describe       = "describe"
	Clear          = "clear"
	List           = "list"
	ChangePassword = "change-password"
	Search         = "search"
	BustCache      = "bustcache"
	AddMember      = "add-members"
	DeleteMember   = "delete-members"
	Restore        = "restore"
)

// Nouns
const (
	NounSecret       = "secret"
	NounPermission   = "permission"
	NounPolicy       = "policy"
	NounPolicies     = "policies"
	NounAuth         = "auth"
	NounToken        = "token"
	NounUser         = "user"
	NounWhoAmI       = "whoami"
	EvaluateFlag     = "eval"
	Init             = "init"
	NounClient       = "client"
	NounAwsProfile   = "awsprofile"
	NounCliConfig    = "cli-config"
	NounConfig       = "config"
	NounRole         = "role"
	NounUsage        = "usage"
	NounGroup        = "group"
	NounAuthProvider = "auth-provider"
	NounLogs         = "logs"
	NounAudit        = "audit"
	NounPrincipal    = "principal"
	NounPki          = "pki"
)

// Cli-Config only
const (
	Editor = "editor"
)

// Flags
const (
	Key               = "key"
	Value             = "value"
	Profile           = "profile"
	Path              = "path"
	ID                = "id"
	Data              = "data"
	Username          = "auth.username"
	Tenant            = "tenant"
	Password          = "auth.password"
	SecurePassword    = "auth.securePassword"
	CurrentPassword   = "currentPassword"
	NewPassword       = "newPassword"
	AuthProvider      = "auth.provider"
	Encoding          = "encoding"
	Beautify          = "beautify"
	Plain             = "plain"
	Filter            = "filter"
	Verbose           = "verbose"
	Config            = "config"
	Dev               = "dev"
	AuthType          = "auth.type"
	AwsProfile        = "auth.awsprofile"
	GcpProject        = "auth.gcp.project"
	GcpToken          = "auth.gcp.token"
	GcpServiceAccount = "auth.gcp.service"
	GcpAuthType       = "auth.gcp.type"
	AuthClientSecret  = "auth.client.secret"
	AuthClientID      = "auth.client.id"
	AzureAuthClientID = "AZURE_CLIENT_ID"
	Query             = "query"
	SearchLinks       = "search.links"
	Limit             = "limit"
	Cursor            = "cursor"
	RefreshToken      = "refreshtoken"
	Output            = "out"
	Overwrite         = "overwrite"
	ClientID          = "client.id"
	ClientSecret      = "client.secret"
	Version           = "version"
	StartDate         = "startdate"
	EndDate           = "enddate"
	Force             = "force"
)

// Data Flags
const (
	// shared
	DataExternalID  = "external.id"
	DataDescription = "desc"
	DataAttributes  = "attributes"
	DataProvider    = "provider"

	// user
	DataUsername       = "username"
	DataSecurePassword = "securePassword"
	DataPassword       = "password"

	// permission
	DataSubject   = "subjects"
	DataEffect    = "effect"
	DataAction    = "actions"
	DataCondition = "conditions"
	DataCidr      = "cidr"
	DataResource  = "resources"

	// role
	DataName = "name"

	// auth provider
	DataType      = "type"
	DataTenantID  = "azure.tenant.id"
	DataAccountID = "aws.account.id"

	//group
	DataGroupName = "group.name"
)

// Security
const (
	StoreType = "store.type"
	Type      = "type"
	Store     = "store"
)

// Hidden Flags
const (
	StorePath = "store.path"
)

// Encodings
const (
	Yaml      = "yaml"
	Json      = "json"
	YamlShort = "yml"
)

func GetShortFlag(flag string) string {
	switch flag {
	case Tenant:
		return "t"
	case Config:
		return "c"
	default:
		return ""
	}
}