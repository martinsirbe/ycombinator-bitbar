# ycombinator-bitbar
![example](https://i.imgur.com/9LjbmH9.png)  
A basic script for fetching top stories from [**Hacker News API**][1] and outputting them in [**BitBar**][2].  

setup
-----
```go get -u github.com/martinsirbe/ycombinator-bitbar```  
```go install github.com/martinsirbe/ycombinator-bitbar```  
Move the `hacker-news.1m.sh` script to your BitBar plugins directory.  
Update `GOPATH` in the `hacker-news.1m.sh`.

dependencies
------------
* [**Hacker News API**][1]
* [**BitBar**][2]
* [**gorun**][3]

## License
See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

[1]: https://github.com/HackerNews/API
[2]: https://getbitbar.com/
[3]: https://github.com/erning/gorun
