package longwar

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", Home)
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

const homeTemplateHTML = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
	<script type="text/javascript" src="../js/thirdparty/crafty.js"></script>
	<script type="text/javascript" src="../js/game.js"></script>
	<title>LongWar</title>
	<style>
	body, html { margin:0; padding: 0; overflow:hidden }
	</style>
</head>
<body>
</body>
</html>
`

var homeTemplate = template.Must(template.New("home").Parse(homeTemplateHTML))
