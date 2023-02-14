## get cpu and memory average
```
POST /metricbeat-*/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "@timestamp": {
              "gte": "now-5m",
              "lt": "now"
            }
          }
        },
        {
          "term": {
            "host.hostname": $HOSTNAME
          }
        }
      ]
    }
  },
  "aggs": {
    "cpu_usage_avg": {
      "avg": {
        "field": "host.cpu.usage"
      }
    },
    "memory_usage_avg": {
      "avg": {
        "field": "system.memory.used.pct"
      }
    }
  }
}
```
