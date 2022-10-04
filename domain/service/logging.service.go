package service

// ここの層はinterfaceを提供するのみでOK！(DDDの場合)

type LoggingLogic interface {
	LoggingSetting()
}
