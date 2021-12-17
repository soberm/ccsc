package ccsc

type Config struct {
	ClientContractAddress        string
	EthereumHost                 string
	PrivateKey                   string
	CallDurationMeasurementsFile string
	CostsMeasurementsFile        string
	CallerConfig                 CallerConfig
}

type CallerConfig struct {
	SourceHost                 string
	TargetHost                 string
	PrivateKey                 string
	ProxyContractAddress       string
	ServerContractAddress      string
	SrcRelayContractAddress    string
	TargetRelayContractAddress string
	CallerMeasurementFile      string
}
