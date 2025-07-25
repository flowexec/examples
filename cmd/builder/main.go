package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jahvon/flow/types/executable"
)

const (
	schemaComment = "yaml-language-server: $schema=https://flowexec.io/schemas/flowfile_schema.json"
	editComment   = "DO NOT EDIT - this file was generated by the example builder tool"
)

func main() {
	var (
		outputDir     = flag.String("output", "./", "Output directory for generated examples")
		workspaceName = flag.String("workspace", "flow", "Workspace name for examples")
		namespaceName = flag.String("namespace", "examples", "Namespace for examples")
	)
	flag.Parse()

	if err := os.MkdirAll(*outputDir, 0655); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	opts := []Option{
		WithWorkspaceName(*workspaceName),
		WithNamespaceName(*namespaceName),
		WithWorkspacePath(*outputDir),
		WithFlowFilePath(*outputDir),
	}

	generateExamples(*outputDir, opts...)
	fmt.Printf("Successfully generated example files in %s\n", *outputDir)
}

func generateExamples(outputDir string, opts ...Option) {
	var failure bool

	flowFiles := map[string]*executable.FlowFile{
		"nameless.flow.yaml": RootExecFlowFile(opts...),
		"exec.flow.yaml":     ExecFlowFile(opts...),
		"parallel.flow.yaml": ParallelFlowFile(opts...),
		"serial.flow.yaml":   SerialFlowFile(opts...),
		"request.flow.yaml":  RequestFlowFile(opts...),
		"render.flow.yaml":   RenderFlowFile(opts...),
		"launch.flow.yaml":   LaunchFlowFile(opts...),
	}

	flowfileComments := []string{schemaComment, editComment}
	for filename, flowFile := range flowFiles {
		path := filepath.Join(outputDir, filename)
		content, err := flowFile.YAML()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error generating YAML for flow file %s: %v\n", filename, err)
			failure = true
			continue
		}
		if err := writeFileWithComments(path, content, flowfileComments); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error writing flow file %s: %v\n", path, err)
			failure = true
		}
	}

	assetFiles := map[string]string{
		"generated.sh":       GeneratedShellScript(),
		"simple-script.sh":   SimpleShellScript(),
		"hello.sh":           HelloShellScript(),
		"env-script.sh":      EnvShellScript(),
		"template.md":        TemplateMarkdown(),
		"template-data.yaml": TemplateDataYAML(),
		"message.txt":        MessageText(),
	}

	rootFiles := map[string]string{
		"exec-template.flow.tmpl": ExecTemplateFlow(),
	}

	assetsDir := filepath.Join(outputDir, "assets")
	if err := os.MkdirAll(assetsDir, 0755); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error creating supporting directory: %v\n", err)
		failure = true
	}

	for filename, content := range assetFiles {
		path := filepath.Join(assetsDir, filename)
		if err := writeFile(path, content); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", path, err)
			failure = true
		}
	}

	for filename, content := range rootFiles {
		path := filepath.Join(outputDir, filename)
		if err := writeFile(path, content); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", path, err)
			failure = true
		}
	}

	if failure {
		os.Exit(1)
	}
}

func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func writeFileWithComments(path, content string, comments []string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(path)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", path, err)
			}
		} else {
			return err
		}
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, comment := range comments {
		_, err = writer.WriteString(fmt.Sprintf("# %s\n", comment))
		if err != nil {
			return err
		}
	}
	if _, err = writer.WriteString(content); err != nil {
		return err
	}
	if err = writer.Flush(); err != nil {
		return err
	}
	return nil
}
