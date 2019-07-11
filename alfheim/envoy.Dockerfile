FROM envoyproxy/envoy:latest
COPY ./alfheim/envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml