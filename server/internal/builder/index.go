package builder

import (
	"io/ioutil"
	"os"
)

// BuildIndex writes an HTML index page for WASM into the dstFile.
func BuildIndex(dstFile string) error {
	const html = `<html >
<head>
    <meta http-equiv="Pragma" content="no-cache">
    <meta charset="utf-8"/>
    <!-- <meta name='viewport' content='width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0'/> -->
    <meta name="viewport" content="user-scalable=no, initial-scale=1, maximum-scale=1, minimum-scale=1, width=device-width, height=device-height, target-densitydpi=device-dpi" />
    <meta name="apple-mobile-web-app-capable" content="yes" />
    <meta name="apple-mobile-web-app-status-bar-style" content="black" />
    <link rel="stylesheet" type="text/css" href="/material/material-components-web.min.css">
    <link rel="stylesheet" type="text/css" href="/material/wtk.css">

    <script src="wasm_exec.js"></script>
    <script>
        const isSafari = /^((?!chrome|android).)*safari/i.test(navigator.userAgent);

        const go = new Go();

        if (isSafari) {
            fetch('app.wasm').then(response =>
                response.arrayBuffer()
            ).then(bytes =>
                WebAssembly.instantiate(bytes, go.importObject)
            ).then(result =>
                go.run(result.instance)
            );
        } else {
            WebAssembly.instantiateStreaming(fetch("app.wasm"), go.importObject).then((result) => {
                go.run(result.instance);
            });
        }
    </script>
     <!-- <script type="text/javascript" src="/material/material-components-web.min.js"></script> -->
     <script type="text/javascript" src="/material/material-components-web.js"></script>
</head>
<body>
<div id="content" class="content">
    <div style="text-align: center;">
        <div class="lds-ring">
            <div></div>
            <div></div>
            <div></div>
            <div></div>
        </div>
        <br/>
        Bitte warten...
    </div>

</div>
</body>
</html>
`
	return ioutil.WriteFile(dstFile, []byte(html), os.ModePerm)
}
