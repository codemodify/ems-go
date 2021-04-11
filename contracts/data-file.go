package contracts

type ESSObject struct {
	Creation_ms int64   `json:"creation_ms,omitempty"`
	ID          int64   `json:"id,omitempty"`
	SOC         float64 `json:"soc,omitempty"`
	E_kwh       float64 `json:"e_kwh,omitempty"`
	P_kw        float64 `json:"p_kw,omitempty"`
	P_ch_kw     int     `json:"p_ch_kw,omitempty"`
	P_disch_kw  int     `json:"p_disch_kw,omitempty"`
}

type MeterObject struct {
	Creation_ms int64   `json:"creation_ms,omitempty"`
	ID          int64   `json:"id,omitempty"`
	P_kw        float64 `json:"p_kw,omitempty"`
	Q_kva       int64   `json:"q_kva,omitempty"`
}

type TargetObject struct {
	Creation_ms int64   `json:"creation_ms,omitempty"`
	ID          int64   `json:"id,omitempty"`
	P_kw        float64 `json:"p_kw,omitempty"`
}

type DataFile struct {
	ESS    []ESSObject    `json:"ess,omitempty"`
	Meter  []MeterObject  `json:"meter,omitempty"`
	Target []TargetObject `json:"target,omitempty"`
}
