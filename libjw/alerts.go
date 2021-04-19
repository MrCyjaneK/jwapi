package libjw

type Alert struct {
	Title       string
	Description string
	Color       string
	Cause       string
	Callbacks   []AlertCallback
}
type AlertCallback struct {
	Title    string
	Endpoint string
}

var Alerts []Alert
