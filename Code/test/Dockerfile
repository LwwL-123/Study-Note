FROM arm64v8/ubuntu
COPY ./full_cpu  /full_cpu
RUN chmod +x /full_cpu
ENTRYPOINT ["/full_cpu", "-cpunum"]
CMD ["1"]