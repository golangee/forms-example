package builder

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	wasmFilename       = "app.wasm"
	goRootJsBridge     = "misc/wasm/wasm_exec.js"
	wasmBridgeFilename = "wasm_exec.js"
	indexHtml          = "index.html"
)

type BuildErr struct {
	delegate error
}

func (b BuildErr) Error() string {
	return b.delegate.Error()
}

func (b BuildErr) Unwrap() error {
	return b.delegate
}

// GoBuildWasm builds an idiomatic wasm go module. The wasm main entry point must be defined at cmd/wasm. The
// output file is forwarded.
func GoBuildWasm(dir string, outFile string) error {
	modName, err := GoModName(dir)
	if err != nil {
		return err
	}

	err = GoBuild(Options{
		GOOS:       "js",
		GOARCH:     "wasm",
		WorkingDir: dir,
		Output:     outFile,
		Packages:   []string{modName + "/cmd/wasm"},
		LDFLAGS: LDFLAGS{

		},
	})

	if err != nil {
		return err
	}

	return nil
}

// Options represent the various build options for the go build command.
type Options struct {
	GOOS       string
	GOARCH     string
	WorkingDir string
	Output     string
	Packages   []string
	Env        []string
	LDFLAGS    LDFLAGS
}

// LDFLAGS represent the go linker flags.
type LDFLAGS struct {
	// X is to inject variables at compilation/linking time.
	X []string
}

// String returns the linker flags.
func (f *LDFLAGS) String() string {
	sb := &strings.Builder{}
	for _, x := range f.X {
		sb.WriteString("-X ")
		sb.WriteString(x)
		sb.WriteString(" ")
	}

	return sb.String()
}

// GoBuild just issues the go build command.
func GoBuild(opts Options) error {
	args := []string{"build"}
	ldflags := opts.LDFLAGS.String()
	if ldflags != "" {
		args = append(args, "-ldflags", "\""+ldflags+"\"")
	}

	if opts.Output != "" {
		args = append(args, "-o", opts.Output)
	}

	for _, p := range opts.Packages {
		args = append(args, p)
	}

	cmd := exec.Command("go", args...)
	cmd.Dir = opts.WorkingDir
	cmd.Env = opts.Env
	if len(cmd.Env) == 0 {
		cmd.Env = append(cmd.Env, os.Environ()...)
	}

	if opts.GOOS != "" {
		cmd.Env = append(cmd.Env, "GOOS="+opts.GOOS)
	}

	if opts.GOARCH != "" {
		cmd.Env = append(cmd.Env, "GOARCH="+opts.GOARCH)
	}

	res, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, string(res))
	}

	return nil
}

// GoModName returns the name of the go module in the given directory.
func GoModName(dir string) (string, error) {
	cmd := exec.Command("go", "list", "-m")
	cmd.Dir = dir
	cmd.Env = os.Environ()
	res, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(res))
	}

	for _, s := range strings.Split(string(res), "\n") {
		s = strings.TrimSpace(s)
		if s != "" {
			return s, nil
		}
	}

	return "", fmt.Errorf("no module name found: %s", string(res))
}

// GoEnv requests the given parameter name.
func GoEnv(name string) (string, error) {
	cmd := exec.Command("go", "env", name)
	cmd.Env = os.Environ()
	res, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(res))
	}

	for _, s := range strings.Split(string(res), "\n") {
		s = strings.TrimSpace(s)
		if s != "" {
			return s, nil
		}
	}

	return "", nil
}

// CopyFile copies a file from src to dst
func CopyFile(dst, src string) error {
	df, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to open dst file: %w", err)
	}
	defer df.Close()

	sf, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to open src file: %w", err)
	}
	defer sf.Close()

	if _, err := io.Copy(df, sf); err != nil {
		return fmt.Errorf("unable to copy file bytes: %w", err)
	}

	return nil
}

// BuildProject builds an entire golangee wasm project from src to dst.
func BuildProject(srcDir, dstDir string) error {
	wasmHash, err := HashFileTree(srcDir)
	if err != nil {
		return fmt.Errorf("unable to calculate hash version: %w", err)
	}

	goRoot, err := GoEnv("GOROOT")
	if err != nil || goRoot == "" {
		return fmt.Errorf("unable to determine GOROOT: %w", err)
	}

	if err := CopyFile(filepath.Join(dstDir, wasmBridgeFilename), filepath.Join(goRoot, goRootJsBridge)); err != nil {
		return fmt.Errorf("unable to provide wasm-js-bridge: %w", err)
	}

	bridgeHash, err := HashFile(filepath.Join(goRoot, goRootJsBridge))
	if err != nil {
		return fmt.Errorf("unable to hash bridge js: %w", err)
	}

	idxDat := IndexData{
		WasmVersion:       hex.EncodeToString(wasmHash),
		WasmBridgeVersion: hex.EncodeToString(bridgeHash),
		HotReload:         true,
	}

	buildErr := GoBuildWasm(srcDir, filepath.Join(dstDir, wasmFilename))

	if buildErr != nil {
		idxDat.Body = strings.Join(strings.Split(buildErr.Error(), "\n"), "<br/>")
		idxDat.LoadWasm = false
	} else {
		idxDat.LoadWasm = true
	}

	if err := BuildIndex(filepath.Join(dstDir, indexHtml), idxDat); err != nil {
		return fmt.Errorf("unable to create index html: %w", err)
	}

	if buildErr != nil {
		return BuildErr{buildErr}
	}

	return nil
}
