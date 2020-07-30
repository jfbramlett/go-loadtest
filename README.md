# Simple LoadTest Framework
This project holds a simple framework for performing load testing. It is build using Go which offers a powerful environment
for executing concurrent actions via GoRoutines. 

The code for the framework is in the pkg folder.

To use the framework you can import the code in pkg or you can modify the code in cmd/loadtest/main.go to reflect
you current test needs.

The framework is essentially given a "test" which is simply an interface with a *Run(context context.Context) error* which
which contains the code for what actions you want to take. You then given this code to load test runner which executes the
given testing using the defined run profile. A "run profile" defines the concurrency (simulating users), the frequency of tests
 (tests can be run at a steady state, at random intervals, or a combination of the two with some tests run at a steady rate
 with others using a random profile), and a specified ramp up strategy which outlines how we should reach our max concurrency
 (a Noop ramp strategy starts the full load right away while a smooth ramp strategy gradually increases load over a period of time).
 
 After the test completes you can view metrics of the run which includes successes, errors, avg execute time, etc...
 
A predefined "test" is provided that can execute against our HTTP endpoints. This test, NWPRunner, is provided a 
scenario file which outlines the endpoints to execute. You can then run this with the following command line
arguments: 

- -scenario <file> -> the name of the file holding our test config
- -concurrency <int - defaults to 25> -> represents the number of concurrent requests (i.e. users)
- -length <int - defaults to 300> -> the length of the test in seconds
- -interval <int - defaults to 2> -> the number of seconds between requests
- -ramp <int - defaults to 0> -> the ramp strategy - 0=noop, smooth=1, random=2 (defines rollout of users hitting the system, noop will start all concurrent users together, smooth will ramp users in 10 intervals, random will gradual increases 10% of users at random intervals over a ramp period)
- -profile <int - defaults to 1> -> the run profile - 1=static, 2=random, 3=partial random (this drives the actual interval used between tests, static will always wait interval, random will wait between 0-interval, partial does a half static and the rest random)
- -minTime <int - defaults to 1000> -> used by reporting to determine our min bound for execution time
- -maxTime <int - defaults to 1500> -> used by reporting to determine our max bound for execution time   
 
The above relies on a scenario file to define what our execution looks like. A sample is defined in dev-scenario.json.
In this file you define your auth details and then the endpoints to execute. Lastly you define a runtype which
should be set to "all" or "random" depending on whether you want to execute all the endpoints as part of a single test
or execute one of them at random. Using "random" gives you better metrics as only a single call is being made
in a test where "all" reports at the test level not at the individual http call level so you have to do some 
math to get info at the call level.

You can also configure whether metrics are published to Prometheus and/or ElasticSearch via cli arguments.
- -publishToPrometheus <true/false> -> flag indicating if we should publish metrics to Prometheus
- -prometheusUrl <url - defaults to http://localhost:9091> -> the url we want to push metrics to
- -publishToElastic <true/false> -> flag indicating if we should publish metrics to Elastic
- -elasticUrl <url - defaults to http://localhost:9200> -> the elastic url we are publishing to
- -elasticIndex <string - defaults to metrics> -> the elastic index to publish to


To run this you can run locally if you have Go 1.11 or greater installed, however the easiest way to run is
via docker. First build your local image with:

```shell script
./build-docker.sh
```

That will generate your image, you can then run the image with:

```shell script
docker run -v `pwd`:/etc/config nwp/nwp-load-test:latest -scenario /etc/config/dev-scenario.json
```

The above runs the latest image mapping the current working directory as /etc/config and then provides
the command line argument of *-scenario /etc/config/dev-scenario.json* to the program.

The scenario file is treated as a Go template when read in. Replacements are done against key-values from os.Environ()
allowing you to external certain sensitive information.

If using the elastic publisher the following metrics are published:
- perf_total_requests the total number of requests made
- perf_concurrent_requests the number of simulated users
- perf_duration the duration of the tests
- perf_total_time the cumulative time spent in the test if run concurrently
- perf_passed the number of passed tests
- perf_errors the number of failed tests
- perf_avg_time the avg request time (in ms)
- perf_max_request_time the longest request time (in ms)
- perf_min_request_time the fatest request time (in ms)
- perf_requests_below_threshold the # of requests faster than our lower bound threshold
- perf_requests_below_threshold_pct the % of requests faster than our lower bound threshold
- perf_requests_between_threshold # of requests between our lower and upper thresholds
- perf_requests_between_threshold_pct # of requests between our lower and upper thresholds
- perf_requests_above_threshold # of requests above upper thresholds
- perf_requests_above_threshold_percent # of requests above our upper thresholds


 Ideally a "real" framework like Gatling or JMeter would be used but sometimes you just need something simple you can
 code in hence this framework.