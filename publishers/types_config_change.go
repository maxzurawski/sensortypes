package publishers

import (
	"github.com/labstack/gommon/log"
	"github.com/xdevices/sensortypes/config"
	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/publishing"
)

type typesPublisher struct {
	*publishing.Publisher
}

func (p *typesPublisher) PublishDeleteChange(previous, current interface{}) {
	p.publishRaw("delete", previous, current)
}

func (p *typesPublisher) PublishSaveChange(previous, current interface{}) {
	p.publishRaw("save", previous, current)
}

func (p *typesPublisher) PublishUpdateChange(previous, current interface{}) {
	p.publishRaw("update", previous, current)
}

func (p *typesPublisher) publishRaw(routingKeySuffix string, previous, current interface{}) {
	if !config.Config().ConnectToRabbit() {
		log.Info("connection to rabbit disabled")
		return
	}
	p.PublishConfigurationChanged(crosscutting.RoutingKeySensorTypes.String()+"."+routingKeySuffix,
		config.Config().ServiceName(),
		previous,
		current,
	)
}