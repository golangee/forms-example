package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golangee/forms-example/www/nestor"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const assetPrefix = "/assets/gen/"

func buildIndex(src, webDst string) error {
	fragments, err := nestor.ParseDir(src)
	if err != nil {
		return fmt.Errorf("failed to parse dir: %w", err)
	}

	if err := normalizeAndCopyFiles("", fragments, webDst); err != nil {
		return fmt.Errorf("failed to copy attachments: %w", err)
	}

	buf, err := json.Marshal(fragments)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	baseEncoded := base64.StdEncoding.EncodeToString(buf)
	idxFile := filepath.Join(src, "internal", "index", "gen.go")
	genFile := &bytes.Buffer{}
	genFile.WriteString("//generated by forms-example generator\n\n")
	genFile.WriteString("package index\n")
	genFile.WriteString("const tutorials =\"" + baseEncoded + "\"\n")
	if err := ioutil.WriteFile(idxFile, genFile.Bytes(), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func normalizeAndCopyFiles(ctxPath string, frags []*nestor.Fragment, dst string) error {
	for _, frag := range frags {
		buf := &bytes.Buffer{}
		md := goldmark.New(
			goldmark.WithExtensions(extension.GFM),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithXHTML(),
			),
		)

		err := md.Convert([]byte(frag.Body), buf)
		if err != nil {
			return fmt.Errorf("unable to convert body to markdown: %w", err)
		}

		frag.Body = buf.String()

		ctxPath := ctxPath + "/" + frag.ID()
		dstDir := filepath.Join(dst, ctxPath)
		if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
			return err
		}

		for _, attachment := range frag.Teaser {
			dstFile := filepath.Join(dstDir, attachment.Raw)
			err := rewrite(attachment, dstFile)
			if err != nil && err != os.ErrNotExist {
				return fmt.Errorf("unable to rewrite teaser: %w", err)
			}

			attachment.File = ctxPath + "/" + attachment.Raw
			if err == nil {
				attachment.File = assetPrefix + attachment.File
			}
		}

		for _, attachment := range frag.Attachments {
			if strings.HasPrefix(attachment.Raw, "/") {
				attachment.File = attachment.Raw
				continue
			}

			dstFile := filepath.Join(dstDir, attachment.Raw)
			err := rewrite(attachment, dstFile)
			if err != nil && err != os.ErrNotExist {
				return fmt.Errorf("unable to rewrite attachment: %w", err)
			}

			attachment.File = ctxPath + "/" + attachment.Raw

			if err == nil {
				attachment.File = assetPrefix + attachment.File
			}
		}

		if err := normalizeAndCopyFiles(ctxPath, frag.Fragments, dst); err != nil {
			return err
		}
	}

	return nil
}

func rewrite(a *nestor.Attachment, fname string) error {
	if _, err := os.Stat(a.File); err != nil {
		return os.ErrNotExist
	}

	src, err := os.OpenFile(a.File, os.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("unable to open source file")
	}

	defer src.Close()

	dst, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to open dest file: %s", fname)
	}

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("unable to copy: %w")
	}

	return nil
}
