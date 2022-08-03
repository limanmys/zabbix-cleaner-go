package models

type Audit struct {
	Auditid      int    `json:"auditid"`
	Userid       int    `json:"userid"`
	Clock        int    `json:"clock"`
	Action       int    `json:"action"`
	Resourcetype int    `json:"resourcetype"`
	Note         string `json:"note"`
	IP           string `json:"ip"`
	Resourceid   any    `json:"resourceid"`
	Resourcename string `json:"resourcename"`
}

type Events struct {
	Eventid      int    `json:"eventid"`
	Source       int    `json:"source"`
	Object       int    `json:"object"`
	Objectid     int    `json:"objectid"`
	Clock        int    `json:"clock"`
	Value        int    `json:"value"`
	Acknowledged int    `json:"acknowledged"`
	Ns           int    `json:"ns"`
	Name         string `json:"name"`
	Severity     int    `json:"severity"`
}

type Acknowledges struct {
	Acknowledgeid int    `json:"acknowledgeid"`
	Userid        int    `json:"userid"`
	Eventid       int    `json:"eventid"`
	Clock         int    `json:"clock"`
	Message       string `json:"message"`
	Action        int    `json:"action"`
	OldSeverity   int    `json:"old_severity"`
	NewSeverity   int    `json:"new_severity"`
}

type History struct {
	Itemid int `json:"itemid"`
	Clock  int `json:"clock"`
	Value  any `json:"value"`
	Ns     int `json:"ns"`
}

type Trends struct {
	Itemid   int     `json:"itemid"`
	Clock    int     `json:"clock"`
	Num      int     `json:"num"`
	ValueMin float64 `json:"value_min"`
	ValueAvg float64 `json:"value_avg"`
	ValueMax float64 `json:"value_max"`
}

type TrendsUint struct {
	Itemid   int `json:"itemid"`
	Clock    int `json:"clock"`
	Num      int `json:"num"`
	ValueMin int `json:"value_min"`
	ValueAvg int `json:"value_avg"`
	ValueMax int `json:"value_max"`
}
