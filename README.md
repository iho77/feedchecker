# feedchecker
Check log streams against Threat Intelligence feeds with rsyslog, Kafka and some go code
Tested on real load with ~32000 IoCs in dns list, 20000 IoCs on URL list and 4000 IoCs in IP list. Under 25 000 EPS load consume around 90% CPU and ~9 Gb RAM on single virtual PC.

Feeds used for test purposes were taken from here https://rstcloud.net/

For usage rsyslog, Kafka and MongoDB should be deployed

Event pipeline: event sources (SIEM) -> rsyslog -> Kafka -> [workers] -> backend -> frontend -> user
Alarm pipeline:    [workers] -> Kafka -> jira connector -> Jira




