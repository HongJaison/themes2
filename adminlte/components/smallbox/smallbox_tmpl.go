package smallbox

var List = map[string]string{
	"smallbox": `{{define "smallbox"}}
    <div class="small-box bg-{{.Color}}">
        <div class="inner">
            <h3>{{langHtml .Value}}</h3>
            <p>{{langHtml .Title}}</p>
        </div>
        <div class="icon">
            <i class="fa {{.Icon}}"></i>
        </div>
        <a href="{{.Url}}" class="small-box-footer">
            {{lang "more"}}
            <i class="fa fa-arrow-circle-right"></i>
        </a>
    </div>
{{end}}`,
}