# linetocmd
parse string command line to golang cmd
# example

```go
const command = `ping "127.0.0.1" -c 1`
cmd, err := Parse(command)
if err != nil {
	t.Fatalf(err.Error())
}
data, err := cmd.Output()
if err != nil {
	t.Fatalf(err.Error())
}
t.Logf("%s", data)
```
