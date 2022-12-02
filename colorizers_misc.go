package glog

func Password(password string) string {
	return Wrap(password, LoggerConfig.ColorPassword)
}

func Error(err error) string {
	return Wrap(err.Error(), LoggerConfig.ColorError)
}

func Reason(reason string) string {
	return Wrap(reason, LoggerConfig.ColorReason)
}

func File(file string) string {
	return Wrap(file, LoggerConfig.ColorFile)
}
