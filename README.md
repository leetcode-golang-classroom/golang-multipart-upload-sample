# golang-multipart-upload-sample

This repository is for demo how to implement multipart file upload on golang

## what is multipart upload?

  A way to send large files by breaking them into smaller parts and encoding them into a single HTTP request.

## Multipart upload vs regular upload

* Encoding: Multipart MIME format vs plain binary/text
* File size: Unlimited (chunks) vs request size limits.
* Error handling: Retry chunks vs re-upload everything.
* Flexibility: Supports extra metadata easily