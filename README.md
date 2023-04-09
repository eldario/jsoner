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
