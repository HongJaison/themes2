# GoAdmin Official Themes

- [adminlte](https://github.com/HongJaison/themes2/tree/master/adminlte)
- [sword](https://github.com/HongJaison/themes2/tree/master/sword)

## How to use

- Import the theme
- Set in the global configuration of GoAdmin

```go

package main

import (
	...
	_ "github.com/HongJaison/themes2/adminlte"
	...
)

func main()  {
	
	...
	
	cfg := config.Config{
    		...
    		
    		Theme: "adminlte",
    		
    		...
    	}
	
	...
 
}

```

## How to modify and make it work

Use the Makefile under each theme directory.