# party

A tool for Go to automate the use of third party package versions in the [Google/Camlistore way](http://golang.org/doc/faq#get_version).

Many existing Go package managers provide solutions to the problem of relying on other projects' code. Go handles the use of other packages elegantly, but is prone to problems when that code has breaking changes or is removed from the Internet. As a solution, the FAQ suggests copying any third party repositories into your own project and changing your import paths to match. This is a recursive process, as you must then perform the same procedure for those copied packages and so on, until all leaf packages have no external imports.

party is a tool that performs all of this work. It

1. Rewrites all external imports in both your project and the `third_party` directory from `other.com/user/package` to `host.com/user/your_project/third_party/other.com/user/package`.
1. Updates or copies files from `$GOPATH/src/other.com/user/package` to `$GOPATH/src/host.com/user/your_project/third_party/other.com/user/package`.
1. Returns to step 1 until no more rewrites are performed.

# usage

1. Install with `go get github.com/mjibson/party`.
1. From your project's root directory, run `party -c`.

The `-c` flag will create the `third_party` directory if it does not exist. On further uses, invoking `party` with no arguments is sufficient. This is for protection so that `party` is not invoked at, say, `$GOPATH/src`, which would perform path rewriting and file updating for many files. When run without `-c`, `party` will fail if `third_party` does not exist.

A `-v` flag is available for verbosity. A `-n` flag is available to perform a dry run, in which no actions are taken.

### app engine / relative imports

When working on an App Engine project, use the `-r` flag. This will cause all `third_party` imports to be local packages. This is needed because otherwise App Engine will run `init()` routines from your `.go` files twice.

# updating

To update a third_party package to a newer version, fetch the most recent version of it with `go get -u other.com/user/package`, and run `party` again in your project's directory.

# goven

goven is a precursor to party, and was useful enough to get a mention in the FAQ. However, goven has suffered from neglect and so has some bugs, lacks features, and is not cross-platform. party strives to be an improvement to goven in all ways.

# bugs

There is no guarantee that the source directory exists. For example, a package may have a OS-specific file (`file_windows.go`) which imports another package. On non-Windows, that other package will not have been fetched during `go get` and will thus not exist. These packages must be fetched using `go get` or some other method by hand. Or, running `party` on a Windows machine (in this case) will correctly copy those files. These cases are reported when running `party`.

`party` performs the equivalent of `gofmt` on any files affected by the import rewrite, thus the diff may be somewhat larger than just the import lines.

Third party imports that use local (relative) packages will fail, as those packages are assumed to be part of the standard library. Those packages will not be copied.
