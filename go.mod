module github.com/golangee/forms-example

go 1.14

replace (
	github.com/golangee/forms => ../forms
	github.com/golangee/i18n => ../i18n
)

require github.com/golangee/forms v0.0.0

require (
	github.com/golangee/i18n v0.0.0
	github.com/lpar/gzipped/v2 v2.0.0-rc1
	github.com/worldiety/i18n v0.0.0-20200217125020-b17c54770fb1
)
