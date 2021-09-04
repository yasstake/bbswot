
from(bucket: "btc")
  |> range(start: 1970-01-01T00:00:00Z)
  |> filter(fn: (r) => r["_measurement"] == "exec")
  |> filter(fn: (r) => r["_field"] == "price" or r["_field"] == "size" or r["_field"]=="side")
  |> drop(columns: ["_start", "_stop", "_measurement"])
  |> pivot(rowKey: ["_time"], columnKey:["_field"], valueColumn: "_value")
  |> count(column: "price")

