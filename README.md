# object
object is a small library which allows you to create new struct instance or set instance value with default、yaml、json tag
```shell
go get github.com/go-pkg-utils/object
```
For example:
---------------------
default tag:
---------------------
```go
func main() {
    db := object.NewWithDefault[DB]()

    data, _ := json.Marshal(db)

    fmt.Println(string(data))
}

type DB struct {
    Redis struct {
        Host string `default:"127.0.0.1"`
    }
}

// output:  {"Redis":{"Host":"127.0.0.1"}}
```
yaml tag: 
---------------------
db.yaml
```yaml
redis:
    Host: 127.0.0.1
```
```go
func main() {
    v := viper.New()
    v.SetConfigFile("db.yaml")
    v.ReadInConfig()

    db := object.NewWithYaml[DB](v)

    data, _ := json.Marshal(db)

    fmt.Println(string(data))
}

type DB struct {
    Redis struct {
        Host string `yaml:"redis.Host"`
    }
}

// type DB struct {
//     Redis struct {
//         Host string `yaml:"Host"`
//     } `yaml:"redis"`
// }

// output:  {"Redis":{"Host":"127.0.0.1"}}
```

json tag:
------
```go
func main() {
    str := `{
        "redis": {
            "Host": "127.0.0.1"
        }
    }`

    db := object.NewWithJson[DB](str)

    data, _ := json.Marshal(db)

    fmt.Println(string(data))
}

type DB struct {
    Redis struct {
        Host string `json:"redis.Host"`
    }
}
 
// type DB struct {
//     Redis struct {
//         Host string `json:"Host"`
//     } `json:"redis"`
// }

// output:  {"Redis":{"Host":"127.0.0.1"}}
```
