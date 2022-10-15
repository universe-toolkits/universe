package universe

type BrokerOptions struct {
}

type Broker struct {
	transporter string
}

func (b *Broker) Call(action string, params interface{}) interface{} {
	return b.transporter
}
