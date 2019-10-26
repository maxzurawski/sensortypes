package publishers

import (
	"github.com/xdevices/sensortypes/config"
	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/publishing"
)

var configChangedPublisher *publishing.Publisher
var typesPublisherInstance *typesPublisher

func Init() {
	if configChangedPublisher == nil && config.Config().ConnectToRabbit() {
		configChangedPublisher = config.Config().InitPublisher()
		// NOTE: once declared - even if we disconnect, exchange will stay there in rabbitmq
		configChangedPublisher.DeclareTopicExchange(crosscutting.TopicConfigurationChanged.String())
	}
}

func TypesConfigChangePublisher() *typesPublisher {
	if typesPublisherInstance == nil {
		typesPublisherInstance = &typesPublisher{
			configChangedPublisher,
		}
	}
	return typesPublisherInstance
}
