# Nginx-JSON-Log-Analyzer

![Go Version](https://img.shields.io/github/go-mod/go-version/fantasticmao/nginx-json-log-analyzer)
[![Go Report Card](https://goreportcard.com/badge/github.com/fantasticmao/nginx-json-log-analyzer)](https://goreportcard.com/report/github.com/fantasticmao/nginx-json-log-analyzer)
[![License](https://img.shields.io/github/license/fantasticmao/nginx-json-log-analyzer)](https://github.com/fantasticmao/nginx-json-log-analyzer/blob/main/LICENSE)

## Nginx Configuration

```text
log_format json_log escape=json '{"time_iso8601":"$time_iso8601",'
                                '"remote_addr":"$remote_addr",'
                                '"request_time":$request_time,'
                                '"request":"$request",'
                                '"status":$status,'
                                '"body_bytes_sent":$body_bytes_sent,'
                                '"http_user_agent":"$http_user_agent"}';
access_log /path/to/access.log json_log
```

Related document: http://nginx.org/en/docs/http/ngx_http_log_module.html

## Supported Statistical Indicators

| Supported | Analyze Type | Indicators                                                                   | Required Fields or Libraries                                                                        |
| --------- | ------------ | ---------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| ✅        | 0            | PV and UV                                                                    | $remote_addr                                                                                        |
| ✅        | 1            | Most visited IPs                                                             | $remote_addr                                                                                        |
| ✅        | 2            | Most visited URIs                                                            | $request                                                                                            |
| ✅        | 3            | Most visited User-Agents                                                     | $http_user_agent                                                                                    |
| TODO      | 4            | Most visited User-Countries                                                  | $remote_addr, [MaxMind GeoIP2 Country Database](https://www.maxmind.com/en/geoip2-country-database) |
| TODO      | 5            | Most visited User-Cities                                                     | $remote_addr, [MaxMind GeoIP2 City Database](https://www.maxmind.com/en/geoip2-city)                |
| ✅        | 6            | Top mean response-time URIs                                                  | $request, $request_time                                                                             |
| TODO      | 7            | Top percentile response-time URIs, e.g. p1(min), p50(median), p95, p100(max) | $request, $request_time                                                                             |
