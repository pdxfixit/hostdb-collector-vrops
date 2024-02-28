FROM scratch

LABEL maintainer="Ben Sandberg <info@pdxfixit.com>" \
      name="hostdb-collector-vrops" \
      vendor="PDXfixIT, LLC"

COPY hostdb-collector-vrops /usr/bin/
COPY config.yaml /etc/hostdb-collector-vrops/

ENTRYPOINT [ "/usr/bin/hostdb-collector-vrops" ]
