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

func CreateAlert(alert Alert) int {
	len := len(Alerts)
	if alert.Color == "" {
		alert.Color = "info"
	}
	if alert.Cause == "" {
		alert.Cause = " "
	}
	Alerts = append(Alerts, alert)
	return len
}

func DeleteAlert(alert int) {
	Alerts[alert] = Alert{}
}

func UpdateDescription(alert int, title string, description string) {
	Alerts[alert].Title = title
	Alerts[alert].Description = description
}
