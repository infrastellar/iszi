package base

type CommandConfig struct {
	LogsDir    string     `json:"logs"`
	Repository string     `json:"repository"`
	Programs   []*Program `json:"programs,omitempty"`
}

type Program struct {
	Name        string `json:"name"`
	RootAccount Root   `json"root_account,omitempty"`
	Path        string `json:"path,omitempty"`
}

type Environment struct {
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`
}

type Space struct {
	Name      string        `json:"name"`
	Config    BackendConfig `json:"config"`
	Operator  string        `json:"operator,omitempty"`
	Region    string        `json:"region,omitempty"`
	AccountId string        `json:"account_id,omitempty"`
}

type BackendConfig struct {
	Bucket        string `json:"bucket"`
	DynamodbTable string `json:"dynamodb_table"`
	Encrypt       bool   `json:"encrypt"`
	KmsKeyID      string `json:"kms_key_id"`
	Region        string `json:"region"`
	Key           string `json:"key,omitempty"`
	RoleArn       string `json:"role_arn,omitempty"`
}

type Region struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id"`
	Stature   string `json:"stature"`
}

type Root struct {
	Region    string `json:"region"`
	AccountId string `json:"account_id"`
}

type Mission struct {
	Name    string `json:"name"`
	Vars    string `json:"vars,omitempty"`
	Path    string `json:"path,omityempty"`
	Version string `json:"version,omitempty"`
}

type BaseStation struct {
	Name, Region, AccountId, RootAccountId, RootRegion, Operator, OperatorArn string
}

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
