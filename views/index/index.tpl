{{append . "HeadStyles" "/static/css/custom.css"}}
{{append . "HeadScripts" "/static/js/custom.js"}}


<div class="container">
    <div class="row vertical-offset-75">

    {{template "alert.tpl" .}}

        <div class="btn-group btn-group-justified" role="group" aria-label="...">
            <a href="#" class="btn btn-default" role="button">Site A</a>
            <a href="#" class="btn btn-default" role="button">Site A</a>
            <a href="#" class="btn btn-default" role="button">Site A</a>
        </div>
        <div class="btn-group btn-group-justified" role="group" aria-label="...">
            <a href="#" class="btn" role="button">Site C</a>
            <a href="#" class="btn" role="button">Site D</a>
            <a href="#" class="btn" role="button">Site E</a>
            <a href="#" class="btn" role="button">Site F</a>
        </div>


        {{range $blockName, $lister := .thelinks}}
            <div class="panel panel-info">
                <div class="panel-heading">
                    <h3 class="panel-title">{{$blockName}}</h3>
                </div>
                <div class="panel-body">

                    <div class="row clearfix">
                    {{range $k, $site := $lister}}
                        <div class="col-md-3">
                            <a href="{{$site.Url}}" target="_blank">{{$site.Name}}</a>
                        </div>
                    {{end}}
                    </div>
                </div>
            </div>

        {{end}}
    </div>
</div>
