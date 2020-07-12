# Elasticsearch

Elasticsearch is being used as non-structured database which manages the employee's information and salary.

## Environment Variable

|**ENVIRONMENT VARIABLE**|**DEFAULT VALUE**|**DESCRIPTION**|
|------------------------|-----------------|---------------|
| ELASTIC_PASSWORD | elastic | Password for elasticsearch database server |

## Quickstart

```yaml
---
cluster.name: "ot-microservices"
network.host: 0.0.0.0

discovery.type: single-node

xpack.license.self_generated.type: trial
xpack.security.enabled: true
xpack.monitoring.collection.enabled: true
```

```shell
make build-image
```
