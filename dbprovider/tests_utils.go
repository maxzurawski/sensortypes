package dbprovider

import "os"

func environmentPreparations() {
	_ = os.Setenv("SERVICE_NAME", "sensortypes")
	_ = os.Setenv("HTTP_PORT", "8101")
	_ = os.Setenv("EUREKA_SERVICE", "http://xdevicesdev.home:8761")
	_ = os.Setenv("DB_PATH", "/Users/l0cke/.databases/xdevices/test/sensortypes.db")
	InitDbManager()
}
