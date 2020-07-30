package metrics

const (
	nameKey = "name"
	valueKey = "value"
	timestampKey = "@timestamp"
	verticleKey = "Verticle"
)

type Metric interface {
	Values() map[string]interface{}
	AddTag(name string, val interface{})
}

type Gauge struct {
	data		map[string]interface{}

}

func NewGauge(verticle string, name string, value float64, timestamp string) Metric {
	data := make(map[string]interface{})
	data[nameKey] = name
	data[valueKey] = value
	data[timestampKey] = timestamp
	data[verticleKey] = verticle
	return &Gauge{data: data}
}

func (g *Gauge) Values() map[string]interface{} {
	return g.data
}

func (g *Gauge) AddTag(name string, val interface{}) {
	g.data[name] = val
}
