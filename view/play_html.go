package view

const play_html = `<html>
<head>
<title>{{.Title}}</title>
<style>
	body           {padding: 0px; margin: 0px; text-align: center; font-family: Arial;}
	h3             {padding: 1px; margin: 2px;}
	#jsRequiredMsg {padding: 3px; margin: 3px; font-weight: bold; background: #faa;}
	#controls      {padding: 2px;}
	#controls *    {margin-left: 3px; margin-right: 3px;}
	#view          {position: relative; padding: 1px;}
	#img           {background: #fff; border: 1px solid black;}
	#errMsg        {visibility: hidden; position: absolute; top: 10px; right: 0px; width: 100%; color: #ff3030; font-weight: bold;}
	#footer        {margin-top: 5px; font-size: 90%; font-style: italic;}
</style>
</head>

<body>

<noscript>
	<div id="jsRequiredMsg">
		Your web browser must have JavaScript enabled in order for this page to display correctly!
	</div>
</noscript>

<h3>{{.Title}}</h3>

<div id="controls">
	
	<button id="newGame" onclick="newGame()">New Game</button>
	
	<a href="/help" target="_blank">Help</a>
	
	<a href="https://github.com/skiptomyliu/gomoku" target="_blank" title="Visit Home Page">github</a>
</div>

<div id="view">
	<img id="img" width="{{.Width}}" height="{{.Height}}"
		onload="errMsg.style.visibility = 'hidden'; imgLoaded = true;"
		onerror="errMsg.style.visibility = 'visible'; setTimeout('imgLoaded = true;', 1000);"
		onmousedown="imgClicked(event)"/>
	<div id="errMsg">Connection Error or Application Closed!</div>
</div>

<div id="footer">
	<a href="https://github.com/skiptomyliu/gomoku">https://github.com/skiptomyliu/gomoku</a>
</div>

<script>
	var runId = {{.RunId}};
	var playing = false, imgLoaded = true;
	
	// HTML elements:
	var img            = document.getElementById("img"),
		errMsg         = document.getElementById("errMsg"),
		pauseResumeBtn = document.getElementById("pauseResume");
	
	// Disable image dragging and right-click context menu:
	img.oncontextmenu = img.ondragstart = function() { return false; }
	
	pauseResume();
	refresh();
	
	// Kick-off:
	console.log("Kick off...");
	setInterval(checkRunId, 10000);
	
	function pauseResume() {
		playing = !playing;
		imgLoaded = true;
		// pauseResumeBtn.innerText = playing ? "Freeze" : "Resume";
	}


	function refresh() {
		if (playing && imgLoaded) {
			imgLoaded = false;
			img.src = "/img?quality=75&t=" + new Date().getTime();
			setTimeout(refresh, 1000);
		}
		else
			setTimeout(refresh, 1000);
	}
	
	
	function imgClicked(e) {
		if (!playing)
			return;
		// Relative mouse coordinates inside image:
		var x, y;
		if (document.all) { // For IE, this is enough (exact):
			x = e.offsetX;
			y = e.offsetY;
		} else {            // For other browsers:
	    	x = e.clientX;
	    	y = e.clientY;
	    	for (var el = img; el; el = el.offsetParent) {
	    		x -= el.offsetLeft - el.scrollLeft + el.clientLeft + 7;
        		y -= el.offsetTop - el.scrollTop + el.clientTop + 7;
	    	}
    	}
    	
		var r = new XMLHttpRequest();
		r.open("GET", "/clicked?x=" + x + "&y=" + y + "&b=" + e.button + "&t=" + new Date().getTime(), true);
		r.send(null);
	}
	
	function checkRunId() {
		console.log("checking run id...");
		if (!playing)
			return;
		var r = new XMLHttpRequest();
		r.open("GET", "/runid?t=" + new Date().getTime(), true);
		r.onreadystatechange = function() {
			if (r.readyState == 4 && r.status == 200 && runId != r.responseText)
				window.location.reload(); // App was restarted, reload page
		};
		r.send(null);
	}
	
	function newGame() {
		var r = new XMLHttpRequest();
		r.open("GET", "/new?t=" + new Date().getTime(), true);
		r.onreadystatechange = function() {
			if (r.readyState == 4 && r.status == 200 && runId != r.responseText)
				// New game was started, reload page
				if (!playing)
					pauseResume();
		};
		r.send(null);
	}
</script>

</body>
</html>
`