package dbprovider

import (
	"os"

	"github.com/maxzurawski/utilities/stringutils"
)

func EnvironmentPreparations() {
	_ = os.Setenv("SERVICE_NAME", "sensortypes")
	_ = os.Setenv("HTTP_PORT", "8101")
	_ = os.Setenv("EUREKA_SERVICE", "http://xdevicesdev.home:8761")
	userHomeDir := stringutils.UserHomeDir()
	_ = os.Setenv("DB_PATH", userHomeDir+"/.databases/xdevices/test/sensortypes.db")
	_ = os.Setenv("CONNECT_TO_RABBIT", "FALSE")
	InitDbManager()
}
