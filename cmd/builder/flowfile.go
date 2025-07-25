package main

import (
	"github.com/jahvon/flow/types/executable"
)

// RootExecFlowFile generates the root exec examples FlowFile
func RootExecFlowFile(opts ...Option) *executable.FlowFile {
	d := &executable.FlowFile{
		Tags: sharedExecTags(),
		Executables: []*executable.Executable{
			NamelessExec(opts...),
		},
	}
	if len(opts) > 0 {
		vals := NewOptionValues(opts...)
		d.SetContext(vals.WorkspaceName, vals.WorkspacePath, vals.FlowFilePath)
	}
	return d
}

// ExecFlowFile generates the main exec examples FlowFile
func ExecFlowFile(opts ...Option) *executable.FlowFile {
	d := &executable.FlowFile{
		Namespace:   "exec",
		Description: "This is a flow executable that demonstrates how to use exec executable types.",
		Tags:        []string{"exec"},
		FromFile:    []string{"assets/generated.sh"},
		Executables: []*executable.Executable{
			SimpleExec(opts...),
			SimpleFileExec(opts...),
			ExecWithPauses(opts...),
			ExecWithExit(opts...),
			ExecWithTmpDir(opts...),
			ExecWithArgs(opts...),
			ExecWithParams(opts...),
			ExecWithLogMode(opts...),
			ExecWithTimeout(opts...),
			ExecWithInput(opts...),
			{
				Verb: "run",
				Name: "with-vim",
				// TODO: verify that processes like vim can be started mid exec
				Visibility: func() *executable.ExecutableVisibility { v := executable.ExecutableVisibility("hidden"); return &v }(),
				Exec: &executable.ExecExecutableType{
					Cmd: "echo \"Opening vim...\"\nvim\necho \"Vim closed.\"",
				},
			},
		},
	}
	if len(opts) > 0 {
		vals := NewOptionValues(opts...)
		d.SetContext(vals.WorkspaceName, vals.WorkspacePath, vals.FlowFilePath)
	}
	return d
}

// ParallelFlowFile generates the parallel examples FlowFile
func ParallelFlowFile(opts ...Option) *executable.FlowFile {
	d := &executable.FlowFile{
		Namespace:   "parallel",
		Description: parallelBaseDesc,
		Tags:        []string{"parallel"},
		Executables: []*executable.Executable{
			ParallelExecByRefConfig(opts...),
			ParallelExecWithExit(opts...),
			ParallelExecWithMaxThreads(opts...),
		},
	}
	if len(opts) > 0 {
		vals := NewOptionValues(opts...)
		d.SetContext(vals.WorkspaceName, vals.WorkspacePath, vals.FlowFilePath)
	}
	return d
}

// SerialFlowFile generates the serial examples FlowFile
func SerialFlowFile(opts ...Option) *executable.FlowFile {
	d := &executable.FlowFile{
		Namespace:   "serial",
		Description: serialBaseDesc,
		Tags:        []string{"serial"},
		Executables: []*executable.Executable{
			SerialExecByRefConfig(opts...),
			SerialExecWithExit(opts...),
		},
	}
	if len(opts) > 0 {
		vals := NewOptionValues(opts...)
		d.SetContext(vals.WorkspaceName, vals.WorkspacePath, vals.FlowFilePath)
	}
	return d
}

// RequestFlowFile generates the request examples FlowFile
func RequestFlowFile(opts ...Option) *executable.FlowFile {
	d := &executable.FlowFile{
		Namespace:   "request",
		Description: requestBaseDesc,
		Tags:        []string{"request"},
		Executables: []*executable.Executable{
			RequestExec(opts...),
			RequestExecWithValidatedStatus(opts...),
			RequestExecWithTimeout(opts...),
			RequestExecWithBody(opts...),
			RequestExecWithTransform(opts...),
		},
	}
	if len(opts) > 0 {
		vals := NewOptionValues(opts...)
		d.SetContext(vals.WorkspaceName, vals.WorkspacePath, vals.FlowFilePath)
	}
	return d
}

// RenderFlowFile generates the render examples FlowFile
func RenderFlowFile(opts ...Option) *executable.FlowFile {
	d := &executable.FlowFile{
		Namespace:   "render",
		Description: "This is a flow executable that demonstrates how to use render executable types.",
		Tags:        []string{"render"},
		Executables: []*executable.Executable{
			{
				Verb: "view",
				Name: "markdown",
				Render: &executable.RenderExecutableType{
					TemplateFile:     "assets/template.md",
					TemplateDataFile: "assets/template-data.yaml",
					Params: executable.ParameterList{
						{Prompt: "What is your name?", EnvKey: "NAME"},
						{Text: "Hi", EnvKey: "GREETING"},
					},
				},
			},
		},
	}
	if len(opts) > 0 {
		vals := NewOptionValues(opts...)
		d.SetContext(vals.WorkspaceName, vals.WorkspacePath, vals.FlowFilePath)
	}
	return d
}

// LaunchFlowFile generates the launch examples FlowFile
func LaunchFlowFile(opts ...Option) *executable.FlowFile {
	d := &executable.FlowFile{
		Namespace:   "launch",
		Description: "This is a flow executable that demonstrates how to use launch executable types.",
		Tags:        []string{"launch"},
		Executables: []*executable.Executable{
			LaunchGitHubExample(opts...),
			LaunchWorkspaceExample(opts...),
			LaunchMacSettingsExample(opts...),
		},
	}
	if len(opts) > 0 {
		vals := NewOptionValues(opts...)
		d.SetContext(vals.WorkspaceName, vals.WorkspacePath, vals.FlowFilePath)
	}
	return d
}
