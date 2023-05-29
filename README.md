# rssfinder

![](https://img.shields.io/github/go-mod/go-version/numb95/rssfinder?style=for-the-badge)
![](https://img.shields.io/github/license/numb95/rssfinder?style=for-the-badge)
![](https://img.shields.io/github/actions/workflow/status/numb95/rssfinder/go.yml?style=for-the-badge)

Install the rssfinder using the following command:

```bash 
go get github.com/numb95/rssfinder
```
Import it in the project:
```go
import "github.com/numb95/rssfinder"
```

## Index

- [func FindRSSFeeds(url string) ([]string, error)](<#func-findrssfeeds>)


## func FindRSSFeeds

```go
func FindRSSFeeds(url string) ([]string, error)
```

FindRSSFeeds finds the RSS feeds in a webpage.

## Example Code

```go
package main

import (
	"fmt"
	"github.com/numb95/rssfinder"
)

func main() {
	url := "https://goodarzi.net"
	rss, err := rssfinder.FindRSSFeeds(url)
	if err != nil {
		panic(err)
	}
	for _, feed := range rss {
		fmt.Println(feed)
	}
}
```
## License

This project is licensed under the [MIT license](https://opensource.org/license/mit/)
