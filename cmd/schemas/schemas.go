package schemas

type Devices struct {
	Eddi  []Eddi  `json:"eddi,omitempty"`
	Zappi []any   `json:"zappi,omitempty"`
	Harvi []Harvi `json:"harvi,omitempty"`
	Libbi []any   `json:"libbi,omitempty"`
	Asn   string  `json:"asn,omitempty"`
	Fwv   string  `json:"fwv,omitempty"`
	Vhub  int     `json:"vhub,omitempty"`
}
type Eddi struct {
	DeviceClass             string  `json:"deviceClass"`
	Sno                     int     `json:"sno"`
	Dat                     string  `json:"dat"`
	Tim                     string  `json:"tim"`
	Ectp1                   int     `json:"ectp1"`
	Ectp2                   int     `json:"ectp2"`
	Ectp3                   int     `json:"ectp3"`
	Ectt1                   string  `json:"ectt1"`
	Ectt2                   string  `json:"ectt2"`
	Ectt3                   string  `json:"ectt3"`
	Bsm                     int     `json:"bsm"`
	Bst                     int     `json:"bst"`
	Dst                     int     `json:"dst"`
	Div                     int     `json:"div"`
	Frq                     float64 `json:"frq"`
	Grd                     int     `json:"grd"`
	Pha                     int     `json:"pha"`
	Pri                     int     `json:"pri"`
	Sta                     int     `json:"sta"`
	Tz                      int     `json:"tz"`
	Vol                     int     `json:"vol"`
	Hpri                    int     `json:"hpri"`
	Hno                     int     `json:"hno"`
	Ht1                     string  `json:"ht1"`
	Ht2                     string  `json:"ht2"`
	R1A                     int     `json:"r1a"`
	R2A                     int     `json:"r2a"`
	Rbc                     int     `json:"rbc"`
	Tp1                     int     `json:"tp1"`
	Tp2                     int     `json:"tp2"`
	BatteryDischargeEnabled bool    `json:"batteryDischargeEnabled"`
	G100LockoutState        string  `json:"g100LockoutState"`
	Cmt                     int     `json:"cmt"`
	Fwv                     string  `json:"fwv"`
	NewAppAvailable         bool    `json:"newAppAvailable"`
	NewBootloaderAvailable  bool    `json:"newBootloaderAvailable"`
}
type Harvi struct {
	DeviceClass string `json:"deviceClass"`
	Sno         int    `json:"sno"`
	Dat         string `json:"dat"`
	Tim         string `json:"tim"`
	Ectp1       int    `json:"ectp1"`
	Ectp2       int    `json:"ectp2"`
	Ectp3       int    `json:"ectp3"`
	Ectt1       string `json:"ectt1"`
	Ectt2       string `json:"ectt2"`
	Ectt3       string `json:"ectt3"`
	Ect1P       int    `json:"ect1p"`
	Ect2P       int    `json:"ect2p"`
	Ect3P       int    `json:"ect3p"`
	Fwv         string `json:"fwv"`
}
