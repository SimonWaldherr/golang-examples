# Non-Std-lib

## About

This directory has Golang code examples that contain at least one import that is not from the standard library.

## Importing from outside Std-lib

### Downloading a package

Golang comes with a builtin tool for downloading and installing packages called `get`.

#### Common Usage: `go get [packages]`
- `go get github.com/spf13/cobra/cobra`
- `go get golang.org/x/net/html`

`go get`also allows you to pass flags such as `-u` for updating a package that is already installed on your computer.

- `go get -u github.com/spf13/cobra/cobra`
- `go get -u golang.org/x/net/html`

For more information about to use `go get` and the other flags that may be used, see the [Golang Documentation](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies).

#### Remote import paths

Import path names are of the form `vcs/path` (vcs => Version Control System).

The supported version control systems are:

	Bazaar      .bzr
	Fossil      .fossil
	Git         .git
	Mercurial   .hg
	Subversion  .svn
	
When using `go get` to retrieve a package such as `github.com/spf13/cobra/cobra`, go get will make a http(s) request for that package. 

Given that different vcs have different implementations, go get must know which vcs it is dealing with before it can successfully download the package.

1. If the path name has vcs extension, such as `github.com/spf13/cobra/cobra.git`, then go get knows to download the package via git.

2. If the path name does not have a vcs extension, such as `github.com/spf13/cobra/cobra`, then the html page must have a meta of the following form that indicates which vcs was used.
	
    `<meta name="go-import" content="import-prefix vcs repo-root">`

    The meta tag for `github.com/spf13/cobra/cobra` is `<meta name="go-import" content="github.com/spf13/cobra git https://github.com/spf13/cobra.git">`

For more information on remote import paths, see the [Golang Documentation](https://golang.org/cmd/go/#hdr-Import_path_syntax)
