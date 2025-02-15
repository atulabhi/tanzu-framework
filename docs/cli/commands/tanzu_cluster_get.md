## tanzu cluster get

Get details from a cluster

```
tanzu cluster get CLUSTER_NAME [flags]
```

### Options

```
  -h, --help                         help for get
  -n, --namespace string             The namespace from which to get workload clusters. If not provided clusters from all namespaces will be returned
      --show-all-conditions string   List of comma separated kind or kind/name for which we should show all the object's conditions (all to show conditions for all the objects)
      --show-details                 Show details of MachineInfrastructure and BootstrapConfig when ready condition is true or it has the Status, Severity and Reason of the machine's object
      --show-group-members           Expand machine groups whose ready condition has the same Status, Severity and Reason
```

### Options inherited from parent commands

```
      --log-file string   Log file path
  -v, --verbose int32     Number for the log level verbosity(0-9)
```

### SEE ALSO

* [tanzu cluster](tanzu_cluster.md)	 - Kubernetes cluster operations

###### Auto generated by spf13/cobra on 14-Sep-2022
