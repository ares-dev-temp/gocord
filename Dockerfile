FROM debian:stable-slim

COPY GoCord /bin/GoCord

ENV PORT=8991

CMD ["/bin/GoCord"]
