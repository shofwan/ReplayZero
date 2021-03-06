package main

const (
	testKarateExpected = `# Generated by Replay Zero at 18 Feb 20 12:22 PST
Feature:

  Background:
	* url 'http://localhost:8080'

	Scenario: test scenario c1487b92-01a0-4b08-b66d-52c597e88e67
		Given path '/test/api'
		And header User-Agent = 'curl/7.54.0'
		And header Accept = '*/*'
		And header Content-Length = '22'
		And request
		"""
		this is a test payload
		"""

		When method POST
		Then status 200
		And match header X-Real-Server == 'test.server'
		And match header Content-Length == '22'
		And match header Date == 'Date: Tue, 18 Feb 2020 20:42:12 GMT'
		And match response ==
		"""
		Test payload back atcha
		"""

	Scenario: test scenario c1487b92-01a0-4b08-b66d-52c597e88e67
		Given path '/test/api'
		And header User-Agent = 'curl/7.54.0'
		And header Accept = '*/*'
		And header Content-Length = '22'
		And request
		"""
		this is a test payload
		"""

		When method POST
		Then status 200
		And match header X-Real-Server == 'test.server'
		And match header Content-Length == '22'
		And match header Date == 'Date: Tue, 18 Feb 2020 20:42:12 GMT'
		And match response ==
		"""
		Test payload back atcha
		"""


`

	testGatlingExpected = `package com.yourcompany.perf.yourservice

// Generated by Replay Zero at 18 Feb 20 12:22 PST

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

val scenario_0: ScenarioBuilder = scenario("scenario_0")
	.exec(http("http_0"))
	.post("/test/api")
	.headers(
		"User-Agent" = "curl/7.54.0",
		"Accept" = "*/*",
		"Content-Length" = "22"
	)
	.body(StringBody(
	"""
	this is a test payload
	"""
	))

val scenario_1: ScenarioBuilder = scenario("scenario_1")
	.exec(http("http_1"))
	.post("/test/api")
	.headers(
		"User-Agent" = "curl/7.54.0",
		"Accept" = "*/*",
		"Content-Length" = "22"
	)
	.body(StringBody(
	"""
	this is a test payload
	"""
	))


setUp(
	scenario_0.inject(rampUsersPerSec(initialTps) to steadyStateTps during (rampDuration minutes),
		constantUsersPerSec(steadyStateTps) during (steadyStateDuration minutes)),
	scenario_1.inject(rampUsersPerSec(initialTps) to steadyStateTps during (rampDuration minutes),
		constantUsersPerSec(steadyStateTps) during (steadyStateDuration minutes))
)
	.protocols(httpProtocol)
	.assertions(
		details("http_0").responseTime.percentile1.lessThan(TP99_RESPONSE_TIME),
		details("http_0").responseTime.percentile2.lessThan(TP90_RESPONSE_TIME),
		details("http_0").responseTime.percentile3.lessThan(TP50_RESPONSE_TIME),
		details("http_1").responseTime.percentile1.lessThan(TP99_RESPONSE_TIME),
		details("http_1").responseTime.percentile2.lessThan(TP90_RESPONSE_TIME),
		details("http_1").responseTime.percentile3.lessThan(TP50_RESPONSE_TIME),
		global.failedRequests.percent.lessThan(100 - AVAILABILITY_RATE)
	)
}
`
)
