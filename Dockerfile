FROM alpine
COPY ./k8s-api-test /root/
RUN chmod +x /root/k8s-api-test
CMD ["/root/k8s-api-test"]


