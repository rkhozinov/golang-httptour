<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{ .Title }}</title>
	</head>
	<body>
		{{range .Content}}
    <div>{{ . }}</div></br>
    {{end}}
	</body>
</html>
