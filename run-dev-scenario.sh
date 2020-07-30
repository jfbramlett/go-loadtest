#!/bin/sh

docker run -v `pwd`:/etc/config/nwp-load-test nwp/nwp-load-test:latest -elasticUrl http://alias-rsrv:9200 -elasticIndex metrics-2020-07 -scenario /etc/config/nwp-load-test/dev-scenario.json