package pkg

const (
	CONFIGENVNAME string = ".youtoolconfig"

	OUTDIRKEY string = "OUT_DIR"
)

var (
	HOMEDIR       = getHomeDir()
	CONFIGENVPATH = HOMEDIR + "/" + CONFIGENVNAME

	OUTDIRDEFAULTVALUE = HOMEDIR
	OUTDIR             = GetEnvValue(CONFIGENVPATH, OUTDIRKEY)
)

// Default Enviroment Values
var DEFAULTVALUES = []string{OUTDIRKEY + "=" + OUTDIRDEFAULTVALUE}
