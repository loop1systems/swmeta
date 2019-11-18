# SolarWinds Meta

This application will allow you to answer the question "can I do X with Y
component using the API?"

## Example

Querying for an entity without the verbs:

```console
$ swmeta -h localhost -u admin -p password -s Orion.Nodes --verbs
+--------------------------------+-----------+---------+-----------+-----------+-----------+
|              NAME              | CANCREATE | CANREAD | CANUPDATE | CANDELETE | CANINVOKE |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesThresholds          | false     | true    | false     | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesWebCommunityStrings | false     | true    | false     | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.Nodes                    | true      | true    | true      | true      | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesCustomProperties    | false     | true    | true      | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesStats               | false     | true    | false     | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodeSettings             | true      | true    | true      | true      | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesForecastCapacity    | false     | true    | false     | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
```

Querying for an entity with the verbs:

```console
$ swmeta -h localhost -u admin -p password -s Orion.Nodes --verbs
+--------------------------------+-----------+---------+-----------+-----------+-----------+
|              NAME              | CANCREATE | CANREAD | CANUPDATE | CANDELETE | CANINVOKE |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesThresholds          | false     | true    | false     | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesWebCommunityStrings | false     | true    | false     | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.Nodes                    | true      | true    | true      | true      | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesCustomProperties    | false     | true    | true      | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesStats               | false     | true    | false     | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodeSettings             | true      | true    | true      | true      | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
| Orion.NodesForecastCapacity    | false     | true    | false     | false     | true      |
+--------------------------------+-----------+---------+-----------+-----------+-----------+
+-----------------------------+-----------------------------------------+
|         ENTITYNAME          |               METHODNAME                |
+-----------------------------+-----------------------------------------+
| Orion.Nodes                 | Unmanage                                |
+-----------------------------+-----------------------------------------+
| Orion.Nodes                 | Remanage                                |
+-----------------------------+-----------------------------------------+
| Orion.Nodes                 | PollNow                                 |
+-----------------------------+-----------------------------------------+
| Orion.Nodes                 | GetCountOfElementsPerEngineForLicensing |
+-----------------------------+-----------------------------------------+
| Orion.Nodes                 | ScheduleListResources                   |
+-----------------------------+-----------------------------------------+
| Orion.Nodes                 | GetScheduledListResourcesStatus         |
+-----------------------------+-----------------------------------------+
| Orion.Nodes                 | ImportListResourcesResult               |
+-----------------------------+-----------------------------------------+
| Orion.NodesCustomProperties | CreateCustomProperty                    |
+-----------------------------+-----------------------------------------+
| Orion.NodesCustomProperties | CreateCustomPropertyWithValues          |
+-----------------------------+-----------------------------------------+
| Orion.NodesCustomProperties | ModifyCustomProperty                    |
+-----------------------------+-----------------------------------------+
| Orion.NodesCustomProperties | DeleteCustomProperty                    |
+-----------------------------+-----------------------------------------+
| Orion.NodesCustomProperties | ValidateCustomProperty                  |
+-----------------------------+-----------------------------------------+
```

## TODO

- Allow authentication from environment variables
