# Projects

Projects is a service which aggregates all the github projects related to micro.

## Dependencies

The projects service currently depends on elasticsearch directly.

```
helm install elasticsearch --set resources.requests.cpu=100m --set resources.requests.memory=256Mi elastic/elasticsearch
```

In future we want to migrate to a "search" service backed by elasticsearch as a shared resource
