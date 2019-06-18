ARG VERSION
FROM hashicorp/terraform:$VERSION
LABEL maintainer="Dotnet Mentor <info@dotnetmentor.se>"
RUN apk add --update git bash openssh
WORKDIR /work/
ENTRYPOINT ["/bin/bash"]
