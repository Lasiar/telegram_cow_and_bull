<!DOCTYPE html>
<html><head><meta charset="utf-8"><title>Untitled Document.md</title>
</head><body id="preview">
<h1><a id="Telegram_simple_game_0"></a>Telegram simple game</h1>
<h1><a id="Get_started_2"></a>Get started</h1>
<pre><code class="language-sh">git <span class="hljs-built_in">clone</span> https://github.com/Lasiar/telegram_cow_and_bull $(go env|grep GOPATH |awk -F <span class="hljs-string">'='</span> <span class="hljs-string">'{print $2}'</span>|tr <span class="hljs-operator">-d</span> <span class="hljs-string">"\""</span>)/src/telegramGame <span class="hljs-comment">#zsh can substitute variables: GOPATH/src</span>

<span class="hljs-built_in">cd</span>  $(go env|grep GOPATH/src/telegramGame
</code></pre>
<p>Download the tool for managing dependencies:</p>
<pre><code class="language-sh">go get github.com/golang/dep/cmd/dep
</code></pre>
<p>create Vendor</p>
<pre><code class="language-sh">dep ensure -vendor-only
</code></pre>
<p>and build!</p>
<pre><code class="language-sh">go build
</code></pre>
<h3><a id="You_can_also_use_docker_28"></a>You can also use docker:</h3>
<pre><code class="language-sh">docker build -t telegram-game . &amp;&amp; docker run telegram_game:latest  
</code></pre>
<h1><a id="Important_34"></a>Important</h1>
<p>Replace your token in conf.json file</p>
</body></html>
