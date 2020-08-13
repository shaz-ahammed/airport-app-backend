## Golang application starter
Golang web application starter template using Gin framework (https://github.com/gin-gonic/gin). Aims to provide secure default configuration.

Uses some ideas from:
* https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html
* https://youtu.be/rWBSMsLG8po

### Features
* Middleware to add security headers in response including strict CSP policy. If isTlsEnabled flag is set to true HSTS header will be added as well
```
X-Frame-Options: DENY
X-Content-Type-Options: nosniff
Content-Security-Policy: default-src 'none'; upgrade-insecure-requests;
Referrer-Policy: no-referrer
Strict-Transport-Security: max-age=94608000 ;includeSubDomains; preload
```

* Example middleware to configure cache headers returned in response
```
Expires: 0
Pragma: no-cache
Cache-Control: no-store
```

* Example middleware to handle favicon.ico requests and return HTTP 204 No Content

* Console request logging middleware using rs/zerolog
```
Wed, 17 Jun 2020 20:14:14 +0200 INF Request client-ip=127.0.0.1 content-length=0 http-status=200 latency=0 method=GET request-path=/ping user-agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0"
Wed, 17 Jun 2020 20:14:15 +0200 INF Request client-ip=127.0.0.1 content-length=0 http-status=200 latency=0 method=GET request-path=/ping user-agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0"
Wed, 17 Jun 2020 20:14:17 +0200 WRN Request client-ip=127.0.0.1 content-length=0 http-status=404 latency=0 method=GET request-path=/ user-agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0"
Wed, 17 Jun 2020 20:14:17 +0200 WRN Request client-ip=127.0.0.1 content-length=0 http-status=404 latency=0 method=GET request-path=/ user-agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0"
```
