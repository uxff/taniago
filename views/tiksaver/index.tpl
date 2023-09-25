{{append . "HeadStyles" "/static/css/custom.css"}}
{{append . "HeadScripts" "/static/js/custom.js"}}


<div class="container">
    <div class="row">
        <div class="panel panel-default">
            <div class="panel-heading">
                <div class="text-muted " style="display:inline">您的位置：</div>
                <a href="/">首页</a>/Tiktok Saver
            </div>
        </div>
    </div>
    <div class="row">
        <div class="col-md-8 col-md-offset-2">
    		<div class="panel panel-default">
			  	<div class="panel-heading text-center">
			    	<h3 class="panel-title"><strong>
                        Download Tiktok
                    </strong></h3>
			 	</div> 

			  	<div class="panel-body">
                    <form accept-charset="utf-8" role="form" class="form-horizontal" method="POST" action="/tiksaver/download">
                        <div class="form-group">
                            <label class="col-sm-3 control-label">Tiktok Link</label>
                            <div class="col-sm-8">
                                <input class="form-control" placeholder="https://vt.tiktok.com/ZSNeVCqCK" name="link" value="" type="text" required />
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="col-sm-3 control-label">Mirror</label>
                            <div class="col-sm-8">
                                <input class="form-control" placeholder="Do Mirror after downloaded" name="desc" value="" type="text" />
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="" class="col-sm-3 control-label">Resize</label>
                            <div class="col-sm-8">
                                <input class="form-control" name="quotaSpace" value="0" type="number" />
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="" class="col-sm-3 control-label">Adding Text to Speech</label>
                            <div class="col-sm-8">
                                <input class="form-control" placeholder="0.01" name="price" value="0" type="number"  />
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="" class="col-sm-3 control-label">Extract Speech to Text and Show</label>
                            <div class="col-sm-8">
                                <input class="form-control" placeholder="0.01" name="primeCost" value="0" type="number" />
                            </div>
                        </div>
                                
                        <div class="form-group text-center">
                            <div class="col-sm-5"></div>
                            <div class="col-sm-2">
                                <input class="btn btn-success btn-block" type="submit" value="提交">
                            </div>
                            <div class="col-sm-5"></div>
                            </div>
                        </form>
                </div>
			</div>
		</div>
    </div>
    <div class="row">
        <p>{{if ne .link ""}}Downloading {{.link}}{{end}}</p>
        <p>{{if ne .outputfile ""}}{{.outputfile}} is downloaded.{{end}}</p>
        <!-- <p>{{if ne .errmsg ""}}{{.err}}{{end}}</p> -->
    </div>

</div>


{{/*<script src="/static/js/bootstrap.min.js"></script>*/}}
{{/*<script type="text/javascript" src="/static/js/twitter-bootstrap-hover-dropdown.min.js"></script>*/}}
{{/*<script type="text/javascript" src="/static/js/bootstrap-admin-theme-change-size.js"></script>*/}}
