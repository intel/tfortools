# Contributing to Templates for Tools

Templates for Tools is an open source project licensed under the [Apache v2 License] (https://opensource.org/licenses/Apache-2.0)

## Coding Style

Templates for Tools follows the standard formatting recommendations and language idioms set out
in [Effective Go](https://golang.org/doc/effective_go.html) and in the
[Go Code Review Comments wiki](https://github.com/golang/go/wiki/CodeReviewComments).

## Certificate of Origin

In order to get a clear contribution chain of trust we use the [signed-off-by language] (https://01.org/community/signed-process)
used by the Linux kernel project.

## Patch format

Beside the signed-off-by footer, we expect each patch to comply with the following format:

```
Change summary

More detailed explanation of your changes: Why and how.
Wrap it to 72 characters.
See [here] (http://chris.beams.io/posts/git-commit/)
for some more good advices.

Fixes #NUMBER (or URL to the issue)

Signed-off-by: <contributor@foo.com>
```

For example:

```
Fix poorly named identifiers
  
One identifier, fnname, in func.go was poorly named.  It has been renamed
to fnName.  Another identifier retval was not needed and has been removed
entirely.

Fixes #1
    
Signed-off-by: Mark Ryan <mark.d.ryan@intel.com>
```

## Pull requests

We accept github pull requests.

## Quality Controls

We request you give quality assurance some consideration by:

* Adding go unit tests for changes where it makes sense.
* Enabling [Travis CI](https://travis-ci.org/intel/tfortools) on your github fork of Templates for Tools to get continuous integration feedback on your dev/test branches.

## Issue tracking

If you find a bug please [open a github issue](https://github.com/intel/tfortools/issues/new).
