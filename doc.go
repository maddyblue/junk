/*
 * Copyright (c) 2014 Matt Jibson <matt.jibson@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
Party is a tool to automate the use of third party packages.

Many existing Go package managers provide solutions to the problem of relying on
other projects' code. Go handles the use of other packages elegantly, but is
prone to problems when that code has breaking changes or is removed from the
Internet. As a solution, the FAQ suggests copying any third party repositories
into your own project and changing your import paths to match. This is a
recursive process, as you must then perform the same procedure for those copied
packages and so on, until all leaf packages have no external imports.

party is a tool that performs all of this work. It:

1. Rewrites all external imports in both your project and the third party
directory from "other.com/user/package" to
"host.com/user/your_project/_third_party/other.com/user/package".

2. Updates or copies files from "$GOPATH/src/other.com/user/package" to
"$GOPATH/src/host.com/user/your_project/_third_party/other.com/user/package".

3. Returns to step 1 until no more rewrites are performed.

Usage

From your project's root directory, for the first run:

	$ party -c

Subsequent invocations can omit -c.

The flags are:
	-c
		create the third party directory if it does not exist
	-u
		run "go get -d -u <pkg>" on packages imported by party for the
		current package
	-d="_third_party"
		set the third party directory
	-v
		enable verbose output
	-n
		dry run: actions are printed instead of run
	-r
		import third party packages as relative, local imports

App Engine and Relative Imports

When working on an App Engine project, use the -r flag. This will cause all
third party imports to be local packages. This is needed because otherwise App
Engine will run init() routines from your .go files twice if your app is also in
your $GOPATH.

Bugs

There is no guarantee that the source directory exists. For example, a package
may have a OS-specific file (file_windows.go) which imports another package. On
non-Windows, that other package will not have been fetched during "go get" and
will thus not exist. These packages must be fetched using "go get" or some other
method by hand. Or, running party on a Windows machine (in this case) will
correctly copy those files. These cases are reported when running party.

party performs the equivalent of gofmt on any files affected by the import
rewrite, thus the diff may be somewhat larger than just the import lines.

Third party imports that use local (relative) packages will fail, as those
packages are assumed to be part of the standard library. Those packages will not
be copied.

*/
package main
