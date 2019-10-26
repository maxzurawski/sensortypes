package publishers

import (
	"github.com/labstack/gommon/log"
	"github.com/maxzurawski/sensortypes/config"
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/publishing"
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
	p.Reset()
	p.PublishConfigurationChanged(crosscutting.RoutingKeySensorTypes.String()+"."+routingKeySuffix,
		config.Config().ServiceName(),
		previous,
		current,
	)
}
