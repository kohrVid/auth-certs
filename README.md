# certs

Creating a package that will handle certificates, &c. in the kohrVid auth
project. For now this just contains a helper package that finds the cert files
but this will probably change.

<!-- vim-markdown-toc GFM -->

* [Test](#test)

<!-- vim-markdown-toc -->

## Test

The run the test suite:

    gocov test -count=1 ./helpers | gocov report
