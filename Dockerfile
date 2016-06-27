FROM scratch
MAINTAINER Kelsey Hightower <kelsey.hightower@gmail.com>
ADD app-healthz /app-healthz
ENTRYPOINT ["/app-healthz"]
