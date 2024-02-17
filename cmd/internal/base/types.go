package base

// CommandConfig represents the config for the command itself
type CommandConfig struct {
	LogsDir    string     `json:"logs"`
	Repository string     `json:"repository"`
	Programs   []*Program `json:"programs,omitempty"`
}

// Program represents a program object and data we need to interact with it
type Program struct {
	Name        string `json:"name"`
	RootAccount Root   `json"root_account,omitempty"`
	Path        string `json:"path,omitempty"`
}

// Environment is a logical collection of resources, and is rather sparse,
// using "spaces" and "missions" to construct its features.
type Environment struct {
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`
}

// Space represents the remote state configuration for a particular set of
// resources. There is a one to one relationship between environment and space.
type Space struct {
	Name      string        `json:"name"`
	Config    BackendConfig `json:"config"`
	Operator  string        `json:"operator,omitempty"`
	Region    string        `json:"region,omitempty"`
	AccountId string        `json:"account_id,omitempty"`
}

// BackendConfig represents the terraform backend configuration that we render
// out in the config.s3.tfbackend file.
type BackendConfig struct {
	Bucket        string `json:"bucket"`
	DynamodbTable string `json:"dynamodb_table"`
	Encrypt       bool   `json:"encrypt"`
	KmsKeyID      string `json:"kms_key_id"`
	Region        string `json:"region"`
	Key           string `json:"key,omitempty"`
	RoleArn       string `json:"role_arn,omitempty"`
}

// Region represents the information regarding the provider region we need to
// attach missions/prcodures too.
type Region struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id"`
	Stature   string `json:"stature"`
}

// Root represents the root account at the provider, this is where we store our
// spaces and other management data. No infrastructure resources are intended
// to be here.
type Root struct {
	Region    string `json:"region"`
	AccountId string `json:"account_id"`
}

// Mission represents the staged terraform modules that are applied to
// environment/regions and their state stored in a space.
type Mission struct {
	Name    string `json:"name"`
	Vars    string `json:"vars,omitempty"`
	Path    string `json:"path,omityempty"`
	Version string `json:"version,omitempty"`
}

// BaseStation represents the terraform module that we render in a new space.
type BaseStation struct {
	Name, Region, AccountId, RootAccountId, RootRegion, Operator, OperatorArn string
}

// BaseStationOutput represents the outputs of a space/base-station terraform
// module, these map to the terraform output command.
type BaseStationOutput struct {
	Config struct {
		Sensitive bool `json:"sensitive"`
		Type      []any
		Value     struct {
			Bucket        string `json:"bucket"`
			DynamodbTable string `json:"dynamodb_table"`
			Encrypt       bool   `json:"encrypt"`
			KmsKeyID      string `json:"kms_key_id"`
			Region        string `json:"region"`
			RoleArn       string `json:"role_arn"`
		} `json:"value"`
	} `json:"config"`
}
