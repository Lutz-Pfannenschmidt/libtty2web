# libtty2web
It's a cli wrapper for [tty2web](https://github.com/kost/tty2web) and all functionality is provided by the tty2web binary.

It's simply a command builder and runner for [tty2web](https://github.com/kost/tty2web).

## Installation
```bash
go get github.com/Lutz-Pfannenschmidt/libtty2web
```

## Usage

Every tty2web option is available as a function in the libtty2web package using the `With` prefix. (e.g. `libtty2web.WithPermitWrite()`)

to see all features and options of tty2web, please visit the [tty2web](https://github.com/kost/tty2web) repository, run `tty2web --help` or use the `GetHelpMsg()` function from go.

```go
package main

import "github.com/Lutz-Pfannenschmidt/libtty2web"

func main() {
	test := libtty2web.NewTty2Web("bash")
	test.SetBinary("tty2web")
	test.AddOptions(
		libtty2web.WithPermitWrite(),
		libtty2web.WithOnce(),
	)
	err := test.Run()
	if err != nil {
		panic(err)
	}
	test.Wait()
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)