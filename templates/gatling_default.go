package templates

const (
	// GatlingBase is the default template for a generated Gatling file
	GatlingBase = `package com.yourcompany.perf.yourservice

// Generated by Replay Zero at {{ now }}

import io.gatling.core.Predef._
import io.gatling.core.structure.ScenarioBuilder
import io.gatling.http.Predef._
import io.gatling.http.protocol.HttpProtocolBuilder

class YourServiceBaseline extends Simulation {

val TP99_RESPONSE_TIME = 1600
val TP90_RESPONSE_TIME = 800
val TP50_RESPONSE_TIME = 600
val AVAILABILITY_RATE = 99

val testDuration: Integer = java.lang.Integer.getInteger("duration", 10)
val rampDuration: Integer = java.lang.Integer.getInteger("ramp", 2)
require(rampDuration > 0, "Ramp duration must be a natural number!")

val steadyStateDuration: Integer = testDuration.intValue() - rampDuration.intValue()
require(steadyStateDuration > 0, "Total test duration must be greater than ramp duration!")

val initialTps: Int = java.lang.Integer.getInteger("initial", 5)
val steadyStateTps: Int = java.lang.Integer.getInteger("steady", 100)
val host: String = System.getProperty("host", "TODO: put your service's base URL here")
val httpProtocol: HttpProtocolBuilder = http.baseURL(host)

{{ range $index, $event := . -}}
val scenario_{{$index}}: ScenarioBuilder = scenario("scenario_{{$index}}")
	.exec(http("http_{{$index}}"))
	.{{lower $event.HTTPMethod}}("{{$event.Endpoint}}")
	{{- if gt (len $event.ReqHeaders) 0}}
	.headers(
		{{- range $h_index, $header := $event.ReqHeaders}}
		{{- if $h_index -}},{{end}}
		"{{$header.Name}}" = "{{$header.Value}}"
		{{- end}}
	)
	{{- end}}
	{{- if $event.ReqBody}}
	.body(StringBody(
	"""
	{{ $event.ReqBody }}
	"""
	)){{- end }}

{{ end -}}
{{- println -}}
setUp(
	{{- $toSkip := decr (len .)}}
	{{- range $index, $event := . }}
	scenario_{{$index}}.inject(rampUsersPerSec(initialTps) to steadyStateTps during (rampDuration minutes),
		constantUsersPerSec(steadyStateTps) during (steadyStateDuration minutes)){{if ne $index $toSkip }},{{end}}
	{{- end }}
)
	.protocols(httpProtocol)
	.assertions(
		{{- range $index, $event := . -}}
			{{- range $k, $v := mkMap "1:99" "2:90" "3:50"}}
		details("http_{{$index}}").responseTime.percentile{{$k}}.lessThan(TP{{$v}}_RESPONSE_TIME),
			{{- end}}	
		{{- end}}
		global.failedRequests.percent.lessThan(100 - AVAILABILITY_RATE)
	)
}
`
)
