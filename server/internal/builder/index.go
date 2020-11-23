package builder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type IndexData struct {
	WasmVersion       string
	WasmBridgeVersion string
	Body string
}

// BuildIndex writes an HTML index page for WASM into the dstFile.
func BuildIndex(dstFile string, indexData IndexData) error {

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

    <script src="wasm_exec.js?v={{.WasmBridgeVersion}}"></script>
    <script>
        const isSafari = /^((?!chrome|android).)*safari/i.test(navigator.userAgent);

        const go = new Go();

        if (isSafari) {
            fetch('app.wasm?v={{.WasmVersion}}').then(response =>
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
{{.Body}}
</div>
</body>
</html>
`

	tpl, err := template.New("index.html").Parse(html)
	if err != nil {
		return fmt.Errorf("unable to parse html template: %w", err)
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, indexData); err != nil {
		return fmt.Errorf("unable to apply index template: %w", err)
	}

	return ioutil.WriteFile(dstFile, buf.Bytes(), os.ModePerm)
}
