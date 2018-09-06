{{$Base := .Base | ToLower}}
{{$Name := .DataSource.Name | ToLower}}

package models

{{if eq $Base $Name}}
func {{$Base}}DataSource() (string, string) {
	ds := datasource{
		name:   "{{.DataSource.Name}}",
		host:   "{{.Host}}",
		port:   {{.Port}},
		user:   "{{.User}}",
		pwd:    "{{.Pwd}}",
		driver: "{{.Driver}}",
	}
	return ds.driver, ds.ConnString()
}
{{end}}

func {{$Base}}DbName() string {
	return "{{.Base}}"
}
