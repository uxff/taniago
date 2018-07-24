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
            <a href="javascript:;" class="btn" role="button">{{$blockName}}</a>
            <div class="btn-group btn-group-justified" role="group" aria-label="...">
                {{range $k, $site := $lister}}
                <a href="{{$site.Url}}" class="btn btn-default" role="button">{{$site.Name}}</a>
                {{end}}
            </div>

        {{end}}
    </div>
</div>
