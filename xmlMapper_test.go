package eurate

import (
	"testing"
)

func TestXmlMapper(t *testing.T) {
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
 		<Cube time="2016-12-09">
            <Cube currency="USD" rate="1.0559"/>
            <Cube currency="JPY" rate="121.48"/>
            <Cube currency="BGN" rate="1.9558"/>
            <Cube currency="CZK" rate="27.027"/>
            <Cube currency="DKK" rate="7.4369"/>
            <Cube currency="GBP" rate="0.83935"/>
            <Cube currency="HUF" rate="314.45"/>
            <Cube currency="PLN" rate="4.4454"/>
            <Cube currency="RON" rate="4.4994"/>
            <Cube currency="SEK" rate="9.691"/>
            <Cube currency="CHF" rate="1.0756"/>
            <Cube currency="NOK" rate="8.9805"/>
            <Cube currency="HRK" rate="7.5327"/>
            <Cube currency="RUB" rate="66.6988"/>
            <Cube currency="TRY" rate="3.6707"/>
            <Cube currency="AUD" rate="1.4141"/>
            <Cube currency="BRL" rate="3.5815"/>
            <Cube currency="CAD" rate="1.3914"/>
            <Cube currency="CNY" rate="7.2931"/>
            <Cube currency="HKD" rate="8.1934"/>
            <Cube currency="IDR" rate="14072.51"/>
            <Cube currency="ILS" rate="4.04"/>
            <Cube currency="INR" rate="71.341"/>
            <Cube currency="KRW" rate="1237"/>
            <Cube currency="MXN" rate="21.5825"/>
            <Cube currency="MYR" rate="4.6708"/>
            <Cube currency="NZD" rate="1.4745"/>
            <Cube currency="PHP" rate="52.615"/>
            <Cube currency="SGD" rate="1.5079"/>
            <Cube currency="THB" rate="37.664"/>
            <Cube currency="ZAR" rate="14.545"/>
        </Cube>
	</Cube>
</gesmes:Envelope>`


	m := XmlMapper(xmlSting)

	var examples = []struct {
		date string
		currency string
		expected float64
	}{
		{"2016-12-12", "USD", 1.0596},
		{"2016-12-12", "NZD", 1.4793},
		{"2016-12-12", "ZAR", 14.5441},
		//
		{"2016-12-09", "USD", 1.0559},
		{"2016-12-09", "NZD", 1.4745},
		{"2016-12-09", "ZAR", 14.545},
	}

	for _, example := range examples {
		actual := m[example.date][example.currency]
		if actual != example.expected {
			t.Errorf("XmlMapper(%s, %s) == %f, expected %f",  example.date, example.currency, actual, example.expected)
		}
	}

}
