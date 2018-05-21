<script type="text/javascript" src="https://code.jquery.com/jquery-2.0.3.min.js"></script>
<script type="text/javascript" src="/iot/sw-register.min.js"></script>
<script type="text/javascript" src="/ga.js?pageTitle={{.data.InstanceName}}"></script>
<script type="text/javascript" src="lc/lumavate-components.js"></script>
<script language="javascript">
	var deviceInfo = {
			options: [],
			header: [navigator.platform, navigator.userAgent, navigator.appVersion, navigator.vendor, window.opera],
			dataos: [
					{ name: 'Windows Phone', value: 'Windows Phone', version: 'OS' },
					{ name: 'Windows', value: 'Win', version: 'NT' },
					{ name: 'iPhone', value: 'iPhone', version: 'OS' },
					{ name: 'iPad', value: 'iPad', version: 'OS' },
					{ name: 'Kindle', value: 'Silk', version: 'Silk' },
					{ name: 'Android', value: 'Android', version: 'Android' },
					{ name: 'PlayBook', value: 'PlayBook', version: 'OS' },
					{ name: 'BlackBerry', value: 'BlackBerry', version: '/' },
					{ name: 'Macintosh', value: 'Mac', version: 'OS X' },
					{ name: 'Linux', value: 'Linux', version: 'rv' },
					{ name: 'Palm', value: 'Palm', version: 'PalmOS' }
			],
			databrowser: [
					{ name: 'Chrome', value: 'Chrome', version: 'Chrome' },
					{ name: 'Firefox', value: 'Firefox', version: 'Firefox' },
					{ name: 'Safari', value: 'Safari', version: 'Version' },
					{ name: 'Internet Explorer', value: 'MSIE', version: 'MSIE' },
					{ name: 'Opera', value: 'Opera', version: 'Opera' },
					{ name: 'BlackBerry', value: 'CLDC', version: 'CLDC' },
					{ name: 'Mozilla', value: 'Mozilla', version: 'Mozilla' }
			],
			init: function () {
					var agent = this.header.join(' '),
							os = this.matchItem(agent, this.dataos),
							browser = this.matchItem(agent, this.databrowser);

					return { os: os, browser: browser };
			},
			matchItem: function (string, data) {
					var i = 0,
							j = 0,
							html = '',
							regex,
							regexv,
							match,
							matches,
							version;

					for (i = 0; i < data.length; i += 1) {
							regex = new RegExp(data[i].value, 'i');
							match = regex.test(string);
							if (match) {
									regexv = new RegExp(data[i].version + '[- /:;]([\\d._]+)', 'i');
									matches = string.match(regexv);
									version = '';
									if (matches) { if (matches[1]) { matches = matches[1]; } }
									if (matches) {
											matches = matches.split(/[._]+/);
											for (j = 0; j < matches.length; j += 1) {
													if (j === 0) {
															version += matches[j] + '.';
													} else {
															version += matches[j];
													}
											}
									} else {
											version = '0';
									}
									return {
											name: data[i].name,
											version: parseFloat(version)
									};
							}
					}
					return { name: 'unknown', version: 0 };
			}
	};
	function navigate(url, newWindow=false) {
		if (!newWindow) {
			window.location.href = url;
		} else {
			window.open(url);
		}
	}
	function appNavigate(app, $event) {
		var info = deviceInfo.init();
		//console.log(info);
		var storeLink = 'not set';
		switch(info.os.name) {
			case 'Macintosh':
			case 'iPhone':
			case 'iPad':
				storeLink = app.Apple;
				break;
			case 'Android':
				storeLink = app.Google;
				break;
			case 'Windows':
			case 'Windows Phone':
				storeLink = app.Microsoft;
				break;
			default:
				switch(info.browser.name) {
					case 'Chrome':
						storeLink = app.Google;
						break;
					case 'Safari':
						storeLink = app.Apple;
						break;
					case 'Internet Explorer':
						storeLink = app.Microsoft;
						break;
				}
		}
		//console.log(app);
		//console.log(storeLink);
		$event.preventDefault();
		$event.stopPropagation();
		setTimeout(function() {
			navigate(storeLink, app.OpenNewWindow);
		}, 1000);

		try {
			navigate(app.AppLink, app.OpenNewWindow);
		} catch(err) {
			navigate(storeLink, app.OpenNewWindow);
		}
	}
</script>
