package db

import "testing"

func TestTraverseDb(t *testing.T) {
	query5dAll := `from(bucket: "btc") |> range(start: -5d) |> limit(n: 10)`

	TraverseDb(query5dAll)

	queryBoard5dAll := `from(bucket: "btc") |> range(start: -5d) |> filter(fn:(r)=>r["_mesurement"] == "board") |> last)`
	TraverseDb(queryBoard5dAll)

}

func TestTraverseDb2(t *testing.T) {
	query := `
from(bucket: "btc")
  |> range(start: -5d)
  |> filter(fn: (r) => r["_measurement"] == "exec")
  |> filter(fn: (r) => r["_field"] == "price" or r["_field"] == "size" or r["_field"] == "side")
  |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")

`
	TraverseDb(query)
}

func TestTraverseDb3(t *testing.T) {
	query := `
from(bucket: "btc")
  |> range(start: -5d)
  |> filter(fn: (r) => r["_measurement"] == "board")
  |> filter(fn: (r) => r["_field"] == "price" or r["_field"] == "size" or r["_field"] == "side")
  |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")

`
	TraverseDb(query)
}
