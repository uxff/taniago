{{append . "HeadStyles" "/static/css/custom.css"}}
{{append . "HeadScripts" "/static/js/custom.js"}}


<div class="container">
    <div class="row vertical-offset-75">

    {{template "alert.tpl" .}}

        <div class="btn-group btn-group-justified" role="group" aria-label="...">
            <a href="#" class="btn btn-default" role="button">SiteA</a>
            <a href="#" class="btn btn-default" role="button">SiteA</a>
            <a href="#" class="btn btn-default" role="button">SiteA</a>
        </div>
        <div class="btn-group btn-group-justified" role="group" aria-label="...">
            <a href="#" class="btn" role="button">SiteA</a>
            <a href="#" class="btn" role="button">SiteA</a>
            <a href="#" class="btn" role="button">SiteA</a>
            <a href="#" class="btn" role="button">SiteA</a>
        </div>
    </div>
</div>
