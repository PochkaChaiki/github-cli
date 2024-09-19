# GITHUB-CLI INTRODUCTION
===

This CLI has functionality only to manage issues.
The idea is borrowed from the task of Donovan's "The Go Programming Language"
---
To build this project run following in main directory:
``` go build . ```

Before using github-cli one must add github fine-grained token as env variable:
``` export GITHUB_TOKEN=<YOUR TOKEN> ```

To run github-cli:
```./github-cli <command> -o <repo's owner> -r <repo name> <flags>```

===
For CLI making is used Cobra framework.
