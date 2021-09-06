

from(bucket: "btc")
  |> range(start: -5d)
  |> filter(fn: (r) => r["_measurement"] == "oi")
  |> drop(columns: ["size", "_start", "_stop"])


