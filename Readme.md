# Mirror

An echo server you could run locally to view the request data. You could add filter on methods, url (body, regex yet to do) and see it. 

`-port - port to run application`
`-methods GET,PUT,DELETE,POST`  customize what methods to log, will be adding filter based on url, body, headers.

This is similar to hookbin, https://hookb.in/b9ONNWPLMoFNeLz1NoLL where we can inspect the requests.
like https://devdinu.free.beeceptor.com where we can proxy with rules. We can't use the above in prod systems as we've privacy concern. This lets us do it locally.


This lib is still in development.

## Running
`go build && ./mirror`

runs the application in default port 8080, could run `./send_requests` to send sample requests, which logs the request information
(body, url, method, headers)


## Proxy

`config/flags.go` has []Proxy url with the backend. This will be read from config file later.
```
  {MatchingUrl: ".*/static/.*", Backend: "http://localhost:8888"},
  {MatchingUrl: ".*/log/.*", Backend: "http://localhost:8081"}
```
so if the requests matches static it forwards to `localhost:8888`

1. run python static server `python -m SimpleHTTPServer  8888`
2. run mirror application in port 8081 `mirror -port 8081`
3. run actual proxy server `mirror` or ensure its running already

`curl http://localhost:8080/static/some` should give the simplehttpserver static files response, similarly for url `http://localhost:8080/log/some` will be redirected to echo service running in `8081`


## Use

This service could be extended to 
* Partitioning/Sharding
* redirecting requests matching url (eg: read servives, and write service separately)
* Redirect requests to temporary cluster when failure in existing backends

[Screencast Video](https://youtu.be/H_Sk4xxKJkg)
