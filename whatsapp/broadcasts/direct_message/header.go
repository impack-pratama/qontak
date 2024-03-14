package direct_message

type HeaderParameterKV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type HeaderParameter struct {
	Format     string              `json:"format"`
	Parameters []HeaderParameterKV `json:"params"`
}
