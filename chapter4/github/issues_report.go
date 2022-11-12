package github

import "time"

const TemplText = `{{.TotalCount}} тем:
{{range .Items}}--------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | DaysAgo}}
{{end}}`

const TemplHTML = `
<hl>{{.TotalCount}} тем</hl>
<table>
<tr style=’text-align: left’>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td> 
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
