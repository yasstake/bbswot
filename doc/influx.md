# Influxdb 格納スキーマ（Bybit)

## データ構造

###　約定

* measurement
  - exec
* key
  - (none)
* field
  - side   ["Buy"|"Sell"]
  - price  (float64)
  - size   (float64)


### 板情報
* measurement
  - board
* key
  - side   ["Buy"|"Sell"]
  - price  (価格を１０倍にして小数点をカットし文字列化)
* field 
  - price  (float64)
  - size   (float64)

### オープンインタレスト

* measurement 
  - oi
* key
  - (none)
* field
  - size   (int64)

### ファンディングレート

* measurement
   - funding
* key
  - (none)
* field 
  - rate   (float64)



```javascript
ohlc = (tables=<-) =>
     tables     
     |> reduce(
      identity: {
        total: 0.0,
        high: 0.0,
        low: 0.0,
        close: 0.0
      },
      fn: (r, accumulator) => ({
        total: accumulator.total + r._value,
        high: if accumulator.total + r._value > accumulator.high then accumulator.total + r._value else accumulator.high,
        low: if accumulator.total + r._value < accumulator.low then accumulator.total + r._value else accumulator.low,
        close: accumulator.close + r._value,
      })
    )
```