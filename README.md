# https-echo

## Echo the requested url and redirect to HTTPS

[![Docker Image](https://img.shields.io/badge/docker-vibrato%2Fhttps--echo-blue.svg)](https://hub.docker.com/r/vibrato/https-echo/)
[![Go Report](https://goreportcard.com/badge/github.com/vibrato/https-echo)](https://goreportcard.com/report/github.com/vibrato/https-echo)

This is a tiny Go Web server that listens to HTTP requests and redirects them to HTTPS. Intended to run beside an HTTPS protected website or be used in conjunction with a load-balancer to keep HTTPS Everywhere. It compiles to a less than 10MB and can be used in a "FROM scratch" container.

To run, simply download and execute.

    ./https-echo

## Contributing

Bug reports and pull requests are welcome on GitHub at [https://github.com/vibrato/https-echo](https://github.com/vibrato/https-echo). This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The Software is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
