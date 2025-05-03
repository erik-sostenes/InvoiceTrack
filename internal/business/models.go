package business

import "encoding/xml"

type (
	Map map[string]map[string]float64

	DailyStatement struct {
		FUECD         string       `xml:"FUECD,attr"`
		XMLName       xml.Name     `xml:"estadodecuenta"`
		AccountKey    string       `xml:"clv_subcuenta,attr"`
		OperationDate string       `xml:"fecha_oper,attr"`
		Settlements   []Settlement `xml:"liquidaciones>liquidacion"`
	}

	Settlement struct {
		NumLiq   int       `xml:"num_liq,attr"`
		Invoices []Invoice `xml:"facturas>factura"`
	}

	Invoice struct {
		Type     string    `xml:"tipo,attr"`
		Concepts []Concept `xml:"conceptos>concepto"`
	}

	Concept struct {
		TotalAmount float64 `xml:"monto_total"`
	}
)
