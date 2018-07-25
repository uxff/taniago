<header id="topbar" class="navbar navbar-default navbar-fixed-top bs-docs-nav" role="banner">
  <div class="container">
    <div class="navbar-header">
      <button class="navbar-toggle collapsed" type="button" data-toggle="collapse" data-target=".bs-navbar-collapse">
        <span class="sr-only">导航</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a style="font-size: 14px;" class="navbar-brand" rel="home" href="/" >
        <strong>Channel</strong>
      </a>
    </div>

    <nav class="navbar-collapse bs-navbar-collapse collapse" role="navigation" style="height: 1px;">
      <ul itemscope="itemscope" itemtype="http://www.schema.org/SiteNavigationElement" class="nav navbar-nav">
        <li itemprop="name"><a itemprop="url" href='{{urlfor "UsersController.Index"}}'>
          <span class="glyphicon glyphicon-facetime-video">A</span>
        </a></li>
        <li itemprop="name"><a itemprop="url" href="">
          <span class="glyphicon glyphicon-heart">B</span>
        </a></li>
        <li itemprop="name">
            <a itemprop="url" href="javascript:;" class="dropdown-toggle" data-hover="dropdown">
          <span class="glyphicon glyphicon-credit-card">C</span> <b class="caret"></b>
        </a>
            <ul class="dropdown-menu">
                <li role="presentation" class="dropdown-header">Dropdown header</li>
                <li><a href="#">What a action</a></li>
                <li><a href="#">Something else here</a></li>
                <li role="presentation" class="divider"></li>
                <li role="presentation" class="dropdown-header">Dropdown header</li>
                <li><a href="#">Separated link</a></li>
                <li><a href="#">One more separated link</a></li>
            </ul>
        </li>
      </ul>

      <ul itemscope="itemscope" itemtype="http://www.schema.org/SiteNavigationElement" class="nav navbar-nav navbar-right">
        <li class="dropdown">
          <a href="#" role="button" class="dropdown-toggle" data-hover="dropdown">
            <span class='glyphicon glyphicon-info-sign'></span> Account <b class="caret"></b>
          </a>
          <ul itemprop="name" class="dropdown-menu">
            {{if .IsLogin}}
                <li itemprop="name" ><a itemprop="url" href='{{urlfor "UsersController.Logout"}}'>
                  <span class='glyphicon glyphicon-log-out'></span> 退出
                </a></li>
            {{else}}
                <li itemprop="name" ><a itemprop="url" href='{{urlfor "UsersController.Login"}}'>
                  <span class='glyphicon glyphicon-globe'></span> 登录
                </a></li>
            {{end}}
          </ul>
        </li>
      </ul>
    </nav>
  </div>

</header>
