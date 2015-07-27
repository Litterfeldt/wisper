package main

import "strings"
import "regexp"

func Lex(cb string) (codeBlock string) {
	codeBlock = definePuts(cb)
	codeBlock = translateRequire(codeBlock)
	codeBlock = translateFuncDefinitions(codeBlock)
	codeBlock = translateModules(codeBlock)

	return
}

func definePuts(cb string) (codeBlock string) {
	codeBlock = strings.Replace(cb, "puts(", "fmt.Println(", -1)
	return
}

func translateRequire(cb string) (codeBlock string) {
	codeBlock = strings.Replace(cb, "require", "import", -1)
	return
}

func translateFuncDefinitions(cb string) (codeBlock string) {
	lines := strings.Split(cb, "\n")
	encounteredFunctionDefinition := false
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if encounteredFunctionDefinition {
			if line == "end" {
				line = "}"
				encounteredFunctionDefinition = false
			}
		}
		if strings.HasPrefix(line, "def") {
			encounteredFunctionDefinition = true
			line = rewriteFunctionDefinition(line)
		}
		lines[i] = line
	}
	codeBlock = strings.Join(lines, "\n")
	return
}

func rewriteFunctionDefinition(functionDefinition string) (newFd string) {
	//Function name
	re := regexp.MustCompile(`def (.*)\(`)
	fname := re.FindStringSubmatch(functionDefinition)[1]
	// Function arguments
	re = regexp.MustCompile(`\((.*)\)`)
	fargs := re.FindStringSubmatch(functionDefinition)[1]

	newFd = "func " + fname + "(" + fargs + ")" + " {"
	return
}

func translateModules(cb string) (codeBlock string) {
	lines := strings.Split(cb, "\n")
	if strings.HasPrefix(lines[0], "module") {
		encounteredModuleDefinition := false
		for i, line := range lines {
			line = strings.TrimSpace(line)
			if encounteredModuleDefinition {
				if line == "end" {
					line = ""
					encounteredModuleDefinition = false
				}
			}
			if strings.HasPrefix(line, "module") {
				encounteredModuleDefinition = true
				re := regexp.MustCompile(`module (.*)$`)
				moduleName := re.FindStringSubmatch(line)[1]
				moduleName = strings.ToLower(moduleName)
				line = "package " + moduleName
			}
			lines[i] = line
		}
	} else {
		lines = append([]string{"package main"}, lines...)
	}

	codeBlock = strings.Join(lines, "\n")
	return
}
