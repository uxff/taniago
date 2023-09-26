{{append . "HeadStyles" "/static/css/custom.css"}}
{{append . "HeadScripts" "/static/js/custom.js"}}


<div class="container">
        <div class="row">
            <div class="panel panel-default">
                <div class="panel-heading">
                    <div class="text-muted " style="display:inline">您的位置：</div>
                    <a href="/picset">图集首页</a>/{{if ne .fullParentName "."}}<a href="{{.parentLink}}">{{.fullParentName}}</a>/{{end}}{{if ne .curDirName "."}}{{.curDirName}}{{end}}
                </div>
            </div>
        </div>
    <div class="row">
        {{range $k, $aname := .thedirnames}}
        <div class="col-sm-12 col-md-12" style="padding:20px;">
            <div class="row text-center" style="padding: 0px;vertical-align: middle;text-align: center;/*display: table-cell;vertical-align: bottom;*/">
                <a href="{{$aname.Url}}">
                    <img src="{{$aname.Thumb}}" alt="{{$aname.Name}}" style="width:auto; height:auto;max-height: 100%; max-width: 100%;display: inline-block;background-size:contain;">
                </a>
            </div>
            <p style="text-align:center;">{{if $aname.IsDir}}<span class='glyphicon glyphicon-folder-open'></span>{{end}}&nbsp;&nbsp;{{$aname.Name}}</p>
        </div>
        {{end}}
    </div>
    {{template "layouts/paginator.html" .}}

</div>


{{/*<script src="/static/js/bootstrap.min.js"></script>*/}}
{{/*<script type="text/javascript" src="/static/js/twitter-bootstrap-hover-dropdown.min.js"></script>*/}}
{{/*<script type="text/javascript" src="/static/js/bootstrap-admin-theme-change-size.js"></script>*/}}
