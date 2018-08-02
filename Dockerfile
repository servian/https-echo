FROM scratch
LABEL maintainer="Tristan Morgan <tristan@vibrato.com.au>"
LABEL Description="HTTPS_ECHO, echo url and redirect http to https"
EXPOSE 80
WORKDIR /
COPY https-echo /
ENTRYPOINT ["/https-echo"]
