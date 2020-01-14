# certs

Creating a package that will handle certificates, &c. in the kohrVid auth
project. For now this just contains a helper package that finds the cert files
but this will probably change.

<!-- vim-markdown-toc GFM -->

* [Generate certificates](#generate-certificates)
* [Test](#test)

<!-- vim-markdown-toc -->

## Generate certificates

To generate the certificates used in this app locally:

1. Install [OpenSSL](https://www.openssl.org/)
2. Generate the server key:

        openssl genrsa -out server.key 2048
        openssl req -new -x509 -sha256 -key server.key \
          -out server.crt -days 3650 \
          -subj "/C=GB/ST=London/L=London/OU=auth/CN=kohrvid.com"

3. Generate certificate signing request:

        openssl req -new -sha256 -key server.key -out server.csr \
          -subj "/C=GB/ST=London/L=London/OU=auth/CN=kohrvid.com"

4. Generate the client certificate:

        openssl x509 -req -sha256 -in server.csr -signkey server.key \
          -out client.crt -days 3650

These certificates need to be generated in the root directory of their respective
services.

## Test

The run the test suite:

    gocov test -count=1 ./helpers | gocov report
