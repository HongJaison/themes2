package productlist

var List = map[string]string{
	"productlist": `{{define "productlist"}}
<ul class="products-list product-list-in-box">
    {{range $key, $data := .Data}}
    <li class="item">
        <div class="product-img">
            <img src="{{index $data "img"}}" alt="Product Image">
        </div>
        <div class="product-info">
            <a href="javascript:void(0)" class="product-title">{{index $data "title"}}
                {{if eq (index $data "has_tabel") "true"}}
                    <span class="label label-{{index $data "labeltype"}} pull-right">{{index $data "label"}}</span>
                {{end}}
            </a>
            <span class="product-description">
                {{index $data "description"}}
            </span>
        </div>
    </li>
    {{end}}
</ul>
{{end}}`,
}