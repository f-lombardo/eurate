package eurate

import (
	"testing"
)

func TestEurateXmlMapper(t *testing.T) {
	var xmlSting =
		`<?xml version="1.0" encoding="UTF-8"?>
<gesmes:Envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
	<gesmes:subject>Reference rates</gesmes:subject>
	<gesmes:Sender>
		<gesmes:name>European Central Bank</gesmes:name>
	</gesmes:Sender>
	<Cube>
		<Cube time='2016-12-12'>
			<Cube currency='USD' rate='1.0596'/>
			<Cube currency='JPY' rate='122.69'/>
			<Cube currency='BGN' rate='1.9558'/>
			<Cube currency='CZK' rate='27.027'/>
			<Cube currency='DKK' rate='7.4371'/>
			<Cube currency='GBP' rate='0.83900'/>
			<Cube currency='HUF' rate='314.02'/>
			<Cube currency='PLN' rate='4.4530'/>
			<Cube currency='RON' rate='4.5073'/>
			<Cube currency='SEK' rate='9.7445'/>
			<Cube currency='CHF' rate='1.0772'/>
			<Cube currency='NOK' rate='8.9498'/>
			<Cube currency='HRK' rate='7.5335'/>
			<Cube currency='RUB' rate='64.8661'/>
			<Cube currency='TRY' rate='3.7269'/>
			<Cube currency='AUD' rate='1.4174'/>
			<Cube currency='BRL' rate='3.5953'/>
			<Cube currency='CAD' rate='1.3918'/>
			<Cube currency='CNY' rate='7.3225'/>
			<Cube currency='HKD' rate='8.2211'/>
			<Cube currency='IDR' rate='14121.82'/>
			<Cube currency='ILS' rate='4.0497'/>
			<Cube currency='INR' rate='71.4485'/>
			<Cube currency='KRW' rate='1237.73'/>
			<Cube currency='MXN' rate='21.4472'/>
			<Cube currency='MYR' rate='4.6750'/>
			<Cube currency='NZD' rate='1.4793'/>
			<Cube currency='PHP' rate='52.720'/>
			<Cube currency='SGD' rate='1.5145'/>
			<Cube currency='THB' rate='37.764'/>
			<Cube currency='ZAR' rate='14.5441'/>
		</Cube>
	</Cube>
</gesmes:Envelope>`


	m := eurateXmlMapper(xmlSting)

	input := "NZD"
	expected := 1.4793
	actual := m[input]
	if actual != expected {
		t.Errorf("eurateXmlMapper(%s) == %f, expected %f", input, actual, expected)
	}

}
