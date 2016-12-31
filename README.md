# VDN - valid domain name

[![GoDoc](https://godoc.org/github.com/simon-xia/vdn?status.png)](https://godoc.org/github.com/simon-xia/vdn)

VDN(valid domain name) is used to check if s string is a valid domain name

# Criteria
A complete domain name is valid if it meets the following criteria:

- A complete domain name must have one or more subdomain names and one top-level domain name.
- A complete domain name must use dots (.) to separate domain names.
- Domain names must use only alphanumeric characters and dashes (-).
- Domain and subdomain names must not begin or end with dashes (-).
- Domain names must mot have more than 63 characters.
- The top-level domain name must be one of the predefined top-level domain names, like (com), (org), or (ca)


# TODO

- [ ] unicode check

# Ref

[rfc1035](https://tools.ietf.org/html/rfc1035)

[rfc3696](https://tools.ietf.org/html/rfc3696)

[top-level domain db](https://www.iana.org/domains/root/db)