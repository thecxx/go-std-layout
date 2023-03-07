// Copyright 2023 Kami
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/thecxx/go-std-layout/tools/pkg/internal"
)

type command interface {
	Install(ctx context.Context, gp *GoProject, name, parent string) (err error)
	Remove(ctx context.Context, gp *GoProject, name, parent string) (err error)
}

type commandImpl struct {
}

// Install implements commandiface
func (c *commandImpl) Install(ctx context.Context, gp *GoProject, name string, parent string) (err error) {
	if parent != "" {
		if err = c.generateMainCommand(ctx, gp, parent); err != nil {
			return
		}
		if err = c.generateSubCommand(ctx, gp, name, parent); err != nil {
			return
		}
	} else {
		err = c.generateMainCommand(ctx, gp, name)
	}
	return
}

// Remove implements command
func (c *commandImpl) Remove(ctx context.Context, gp *GoProject, name string, parent string) (err error) {
	if parent != "" {
		err = c.removeSubCommand(ctx, gp, name, parent)
	} else {
		err = c.removeMainCommand(ctx, gp, name)
	}
	return
}

func (c *commandImpl) generateMainCommand(ctx context.Context, gp *GoProject, name string) (err error) {
	var (
		cd = fmt.Sprintf("%s/cmd/%s", gp.Workspace, name)
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, name, name)
		hd = fmt.Sprintf("%s/pkg/console/%s/handler", gp.Workspace, name)
		hf = fmt.Sprintf("%s/pkg/console/%s/handler/%s.go", gp.Workspace, name, name)
	)
	hb, err := internal.IsFile(hf)
	if err != nil {
		return err
	}
	if !hb {
		err = os.MkdirAll(hd, 0755)
		if err != nil {
			return err
		}
		mainhandler := c.insertLicenseHeader(c.generateConsoleMainHandler(ctx, gp.Module, name), gp.License.Header)
		err = os.WriteFile(hf, []byte(mainhandler), 0644)
		if err != nil {
			return err
		}
	}
	cb, err := internal.IsFile(cf)
	if err != nil {
		return err
	}
	if !cb {
		err = os.MkdirAll(cd, 0755)
		if err != nil {
			return err
		}
		maincommand := c.insertLicenseHeader(c.generateMainCommandHandler(ctx, gp.Module, name), gp.License.Header)
		err = os.WriteFile(cf, []byte(maincommand), 0644)
		if err != nil {
			return err
		}
	}
	return
}

func (c *commandImpl) generateSubCommand(ctx context.Context, gp *GoProject, name, parent string) (err error) {
	var (
		cd = fmt.Sprintf("%s/cmd/%s", gp.Workspace, parent)
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, parent, name)
		hd = fmt.Sprintf("%s/pkg/console/%s/handler", gp.Workspace, parent)
		hf = fmt.Sprintf("%s/pkg/console/%s/handler/%s.go", gp.Workspace, parent, name)
	)
	hb, err := internal.IsFile(hf)
	if err != nil {
		return err
	}
	if !hb {
		err = os.MkdirAll(hd, 0755)
		if err != nil {
			return err
		}
		subhandler := c.insertLicenseHeader(c.generateConsoleSubHandler(ctx, gp.Module, name, parent), gp.License.Header)
		err = os.WriteFile(hf, []byte(subhandler), 0644)
		if err != nil {
			return err
		}
	}
	cb, err := internal.IsFile(cf)
	if err != nil {
		return err
	}
	if !cb {
		err = os.MkdirAll(cd, 0755)
		if err != nil {
			return err
		}
		subcommand := c.insertLicenseHeader(c.generateSubCommandHandler(ctx, gp.Module, name, parent), gp.License.Header)
		err = os.WriteFile(cf, []byte(subcommand), 0644)
		if err != nil {
			return err
		}
		if err = c.setupSubCommand(ctx, gp, name, parent); err != nil {
			return err
		}
	}
	return
}

func (c *commandImpl) setupSubCommand(ctx context.Context, gp *GoProject, name, parent string) (err error) {
	var (
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, parent, parent)
	)
	buf, err := os.ReadFile(cf)
	if err != nil {
		return err
	}

	src := regexp.MustCompile(`[\t ]+// sub command placeholder\n`).
		ReplaceAll(buf, []byte(fmt.Sprintf("    cmds = append(cmds, %sc)\n    // sub command placeholder\n", name)))

	return os.WriteFile(cf, []byte(src), 0644)
}

func (c *commandImpl) deleteSubCommand(ctx context.Context, gp *GoProject, name, parent string) (err error) {
	var (
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, parent, parent)
	)
	buf, err := os.ReadFile(cf)
	if err != nil {
		return err
	}

	src := regexp.MustCompile(fmt.Sprintf(`[\t ]+cmds = append\(cmds, %sc\)\n`, name)).ReplaceAll(buf, nil)

	return os.WriteFile(cf, []byte(src), 0644)
}

func (c *commandImpl) generateMainCommandHandler(ctx context.Context, mod, name string) (tpl string) {
	var (
		lower   = strings.ToLower(name)
		ucfirst = strings.ToUpper(string(lower[0])) + lower[1:]
	)
	tpl += "package main\n"
	tpl += "\n"
	tpl += "import (\n"
	tpl += "    \"github.com/spf13/cobra\"\n"
	tpl += fmt.Sprintf("    \"%s/pkg/console/%s/handler\"\n", mod, name)
	tpl += ")\n"
	tpl += "\n"
	tpl += "var (\n"
	tpl += fmt.Sprintf("    %sc = &cobra.Command{}\n", lower)
	tpl += "    gflags = &handler.GlobalFlags{}\n"
	tpl += ")\n"
	tpl += "\n"
	tpl += "func init() {\n"
	tpl += "    // var (\n"
	tpl += fmt.Sprintf("    //    flags = &handler.%sFlags{GlobalFlags: gflags}\n", ucfirst)
	tpl += "    // )\n"
	tpl += fmt.Sprintf("    %sc.Use = \"%s\"\n", lower, lower)
	tpl += fmt.Sprintf("    %sc.Short = \"A short description\"\n", lower)
	tpl += fmt.Sprintf("    %sc.Long = \"A long description\"\n", lower)
	tpl += fmt.Sprintf("    %sc.Version = \"1.0.0\"\n", lower)
	tpl += fmt.Sprintf("    %sc.SilenceUsage = true\n", lower)
	tpl += fmt.Sprintf("    %sc.CompletionOptions.HiddenDefaultCmd = true\n", lower)
	tpl += "    // Events\n"
	tpl += fmt.Sprintf("    %sc.RunE = func(cmd *cobra.Command, args []string) error {\n", lower)
	tpl += "        return cmd.Help()\n"
	tpl += fmt.Sprintf("        // return handler.On%sHandler(cmd.Context(), flags, args)\n", ucfirst)
	tpl += "    }\n"
	tpl += "    // Flags\n"
	tpl += fmt.Sprintf("    // if f := %sc.Flags(); f != nil {\n", lower)
	tpl += "    //     f.StringVarP(&flags.Test, \"test\", \"t\", \"\", \"a test flag\")\n"
	tpl += "    // }\n"
	tpl += fmt.Sprintf("    // if pf := %sc.PersistentFlags(); pf != nil {\n", lower)
	tpl += "    //     pf.StringVarP(&gflags.Test, \"test\", \"t\", \"\", \"a test flag\")\n"
	tpl += "    // }\n"
	tpl += "}\n"
	tpl += "\n"
	tpl += "func main() {\n"
	tpl += "    var (\n"
	tpl += "        cmds []*cobra.Command\n"
	tpl += "    )\n"
	tpl += "\n"
	tpl += "    // Register sub commands\n"
	tpl += "    // sub command placeholder\n"
	tpl += "\n"
	tpl += fmt.Sprintf("    %sc.AddCommand(cmds...)\n", lower)
	tpl += "    defer func() {\n"
	tpl += fmt.Sprintf("        %sc.RemoveCommand(cmds...)\n", lower)
	tpl += "    }()\n"
	tpl += "\n"
	tpl += fmt.Sprintf("    %sc.Execute()\n", lower)
	tpl += "}\n"
	tpl += ""
	return
}

func (c *commandImpl) generateSubCommandHandler(ctx context.Context, mod, name, parent string) (tpl string) {
	var (
		lower   = strings.ToLower(name)
		ucfirst = strings.ToUpper(string(lower[0])) + lower[1:]
	)
	tpl += "package main\n"
	tpl += "\n"
	tpl += "import (\n"
	tpl += "    \"github.com/spf13/cobra\"\n"
	tpl += fmt.Sprintf("    \"%s/pkg/console/%s/handler\"\n", mod, parent)
	tpl += ")\n"
	tpl += "\n"
	tpl += "var (\n"
	tpl += fmt.Sprintf("    %sc = &cobra.Command{}\n", lower)
	tpl += ")\n"
	tpl += "\n"
	tpl += "func init() {\n"
	tpl += "    var (\n"
	tpl += fmt.Sprintf("        flags = &handler.%sFlags{GlobalFlags: gflags}\n", ucfirst)
	tpl += "    )\n"
	tpl += fmt.Sprintf("    %sc.Use = \"%s\"\n", lower, lower)
	tpl += fmt.Sprintf("    %sc.Short = \"A short description\"\n", lower)
	tpl += fmt.Sprintf("    %sc.Long = \"A long description\"\n", lower)
	tpl += "    // Events\n"
	tpl += fmt.Sprintf("    %sc.RunE = func(cmd *cobra.Command, args []string) error {\n", lower)
	tpl += fmt.Sprintf("        return handler.On%sHandler(cmd.Context(), flags, args)\n", ucfirst)
	tpl += "    }\n"
	tpl += "    // Flags\n"
	tpl += fmt.Sprintf("    // if f := %sc.Flags(); f != nil {\n", lower)
	tpl += "    //     f.StringVarP(&flags.Test, \"test\", \"t\", \"\", \"a test flag\")\n"
	tpl += "    // }\n"
	tpl += "}\n"
	tpl += ""
	return
}

func (c *commandImpl) generateConsoleMainHandler(ctx context.Context, mod, name string) (tpl string) {
	var (
		lower   = strings.ToLower(name)
		ucfirst = strings.ToUpper(string(lower[0])) + lower[1:]
	)
	tpl += "package handler\n"
	tpl += "\n"
	tpl += "import (\n"
	tpl += "    \"context\"\n"
	tpl += ")\n"
	tpl += "\n"
	tpl += "type GlobalFlags struct {\n"
	tpl += "    // Test string\n"
	tpl += "}\n"
	tpl += "\n"
	tpl += fmt.Sprintf("type %sFlags struct {\n", ucfirst)
	tpl += "    *GlobalFlags\n"
	tpl += "    // Test string\n"
	tpl += "}\n"
	tpl += "\n"
	tpl += fmt.Sprintf("func On%sHandler(ctx context.Context, flags *%sFlags, args []string) (err error) {\n", ucfirst, ucfirst)
	tpl += "    return\n"
	tpl += "}\n"
	tpl += ""
	return
}

func (c *commandImpl) generateConsoleSubHandler(ctx context.Context, mod, name, parent string) (tpl string) {
	var (
		lower   = strings.ToLower(name)
		ucfirst = strings.ToUpper(string(lower[0])) + lower[1:]
	)
	tpl += "package handler\n"
	tpl += "\n"
	tpl += "import (\n"
	tpl += "    \"context\"\n"
	tpl += ")\n"
	tpl += "\n"
	tpl += fmt.Sprintf("type %sFlags struct {\n", ucfirst)
	tpl += "    *GlobalFlags\n"
	tpl += "    // Test string\n"
	tpl += "}\n"
	tpl += "\n"
	tpl += fmt.Sprintf("func On%sHandler(ctx context.Context, flags *%sFlags, args []string) (err error) {\n", ucfirst, ucfirst)
	tpl += "    return\n"
	tpl += "}\n"
	tpl += ""
	return
}

func (c *commandImpl) removeMainCommand(ctx context.Context, gp *GoProject, name string) (err error) {
	var (
		cd = fmt.Sprintf("%s/cmd/%s", gp.Workspace, name)
		hd = fmt.Sprintf("%s/pkg/console/%s", gp.Workspace, name)
	)
	hb, err := internal.IsDir(hd)
	if err != nil {
		return err
	}
	if hb {
		if err = os.RemoveAll(hd); err != nil {
			return
		}
	}
	cb, err := internal.IsDir(cd)
	if err != nil {
		return err
	}
	if cb {
		if err = os.RemoveAll(cd); err != nil {
			return
		}
	}
	return
}

func (c *commandImpl) removeSubCommand(ctx context.Context, gp *GoProject, name, parent string) (err error) {
	var (
		cf = fmt.Sprintf("%s/cmd/%s/%s.go", gp.Workspace, parent, name)
		hf = fmt.Sprintf("%s/pkg/console/%s/handler/%s.go", gp.Workspace, parent, name)
	)
	hb, err := internal.IsFile(hf)
	if err != nil {
		return err
	}
	if hb {
		if err = os.Remove(hf); err != nil {
			return
		}
	}
	cb, err := internal.IsFile(cf)
	if err != nil {
		return err
	}
	if cb {
		if err = os.Remove(cf); err != nil {
			return
		}
		if err = c.deleteSubCommand(ctx, gp, name, parent); err != nil {
			return
		}
	}
	return
}

func (c *commandImpl) insertLicenseHeader(source, license string) (tpl string) {
	license = strings.Trim(license, "\n\r\t ")
	if license != "" {
		tpl += license
		tpl += "\n\n"
	}
	tpl += source
	return
}

var (
	Command command = &commandImpl{}
)
