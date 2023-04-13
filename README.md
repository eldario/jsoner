![](https://habrastorage.org/webt/pt/xw/aw/ptxwawfevwtpf3z_-j2ficijci8.png)
## Description

Simple data converter from `csv` format to `json`.

The idea is a template with data taken from the header of the `csv` file.

## Examples

I have a `sample.csv` file like this

```csv
id,name,contacts_email,phone,contacts_address
1,Ivan,ivan@gmail.com,+4464564641,"5 avenue"
2,Irina,b.irina@gmail.com,+889789464,"21 avenue"
```

I want to get a `output.json` as a result and that each row be in the format, like `sample.json`:

```json
{
  "IDE": "$id",
  "Person": {
    "Name": "$name",
    "Contacts": {
      "EMAIL": "$contacts_email",
      "Phone": "$phone",
      "address": "$contacts_address"
    }
  }
}
```

### Usage

Command & result into file `output.json`:
```sh
$ go run main.go -path=sample.csv -sample=sample.json -output=output.json
Result in output.json  
```

Result in `output.json`
```shell
$ cat output.json
[
 {
  "IDE": "1",
  "Person": {
   "Contacts": {
    "EMAIL": "ivan@gmail.com",
    "Phone": "+4464564641",
    "address": "5 avenue"
   },
   "Name": "Ivan"
  }
 },
 {
  "IDE": "2",
  "Person": {
   "Contacts": {
    "EMAIL": "b.irina@gmail.com",
    "Phone": "+889789464",
    "address": "21 avenue"
   },
   "Name": "Irina"
  }
 }
]                                    
```

For a more compact look, you can use the flag `--compact`:

```shell
$ go run main.go -path=sample.csv -sample=sample.json -output=output.json --compact
```

```shell
$  cat output.json
[{"IDE":"1","Person":{"Contacts":{"EMAIL":"ivan@gmail.com","Phone":"+4464564641","address":"5 avenue"},"Name":"Ivan"}},{"IDE":"2","Person":{"Contacts":{"EMAIL":"b.irina@gmail.com","Phone":"+889789464","address":"21 avenue"},"Name":"Irina"}}]  
```

To output to `stdout` use the `--show` flag:

```shell
$ go run main.go -path=sample.csv -sample=sample.json --compact --show
[{"IDE":"1","Person":{"Contacts":{"EMAIL":"ivan@gmail.com","Phone":"+4464564641","address":"5 avenue"},"Name":"Ivan"}},{"IDE":"2","Person":{"Contacts":{"EMAIL":"b.irina@gmail.com","Phone":"+889789464","address":"21 avenue"},"Name":"Irina"}}]

```

### Available formatters


| Formatter                       | Description                                                                                                        |
|---------------------------------|--------------------------------------------------------------------------------------------------------------------|
| `$field_name`                   | `field_name` - field name from your csv                                                                            |
| `$splitOrNull(field_name,0,4)`  | Splits value from column `field_name` and gets a part from `0` to `4` position. If value is empty - returns `null` |
| `$stringOrEmpty(field_name)`    | Gets current value from `field_name`. If value is empty - returns empty string "", not `null`                      |
| `$int(field_name)`              | Make a given value from `field_name` a type of `Integer`                                                           |
| `$date(field_name,"d/m/y","2")` | Create a `DateTime` value from given value in `field_name` with format using `d m y h i s`. The last parameter is responsible for the deviation from the date. Can be either positive or negative  |