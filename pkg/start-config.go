package pkg

const (
	//Enviroment Paths
	CONFIGENVPATH string = "config.env"

	//Out directory
	OUTDIRKEY          string = "OUT_DIR"
	OUTDIRDEFAULTVALUE string = "~/"
)

// In App Enviroment Values
var (
	OUTDIR = GetEnvValue(CONFIGENVPATH, OUTDIRKEY)
)

// Default Enviroment Values
var DEFAULTVALUES = []string{OUTDIRKEY + "=" + OUTDIRDEFAULTVALUE}
