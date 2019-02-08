# golang-iwlist

Golang scanner and parser for wireless networks

## Usage

```
package main

import (
	"fmt"
	"github.com/MonkeyBuisness/golang-iwlist"
)

func main() {
	cells, err := wlist.Scan("wls1")
	if err != nil {
		panic(err)
	}
	
	fmt.Println(cells)
}

```

## Output

```
[{"cell_number":"01","mac":"BA:69:F4:70:7D:1D","essid":"Guest","mode":"Master","frequency":2.432,"frequency_units":"GHz","channel":5,"encryption_key":false,"encryption":"off","signal_quality":70,"signal_total":70,"signal_level":-36}]
```

## Util

```
$ sudo wlist -i wls1 
```

## Tests

```
$ go test
```