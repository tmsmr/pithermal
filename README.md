# pithermal
*`cat /sys/class/thermal/thermal_zone0/temp` as Go Module.*


## Usage

```go
import "github.com/tmsmr/pithermal"

...

temp, err := pithermal.GetCpuTemp()
if err != nil {
    panic(err)
}
```

---
*Not worth a Module at the moment, but there may be more features in the future...*
