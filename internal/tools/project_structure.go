package tools

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/sonht1109/supercoder-go/internal/utils"
)

type ProjectStructureTool struct {
	GitignorePatterns []string
}

type ProjectStructureToolArguments struct{}

func NewProjectStructureTool() *ProjectStructureTool {
	return &ProjectStructureTool{
		GitignorePatterns: []string{
			"node_modules", "build", "dist", "out", "target", ".idea", ".vscode", ".git",
		},
	}
}

func (pst *ProjectStructureTool) Execute(arguments map[string]any) string {
	fmt.Println(utils.Green("🔎 Reading project structure..."))
	pst.loadGitignore()
	result := pst.buildProjectTree(".", 0)
	return result
}

func (pst *ProjectStructureTool) loadGitignore() {
	file, _ := os.Open(".gitignore")
	defer file.Close()

	if file != nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" && !strings.HasPrefix(line, "#") {
				pst.GitignorePatterns = append(pst.GitignorePatterns, line)
			}
		}
	}
}

func (pst *ProjectStructureTool) isIgnored(path string, isDir bool) bool {
	relPath := strings.TrimPrefix(strings.ReplaceAll(path, "\\", "/"), "./")

	for _, pattern := range pst.GitignorePatterns {
		isDirPattern := strings.HasSuffix(pattern, "/")
		cleanPattern := strings.TrimSuffix(pattern, "/")

		if isDirPattern && !isDir {
			continue
		}

		regex := ""
		if strings.Contains(cleanPattern, "/") {
			// Convert path pattern
			regex = "^" + regexp.QuoteMeta(cleanPattern)
			regex = strings.ReplaceAll(regex, "\\*\\*", ".*")
			regex = strings.ReplaceAll(regex, "\\*", "[^/]*")
		} else {
			// Simple name match
			regex = "^" + strings.ReplaceAll(regexp.QuoteMeta(cleanPattern), "\\*", ".*") + "$"
		}

		matched, _ := regexp.MatchString(regex, relPath)
		if matched {
			return true
		}
	}

	return false
}

func (pst *ProjectStructureTool) buildProjectTree(root string, depth int) string {
	var builder strings.Builder
	prefix := strings.Repeat("│   ", depth)

	if depth == 0 {
		builder.WriteString(".\n")
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		return builder.String()
	}

	var dirs, files []os.DirEntry
	for _, entry := range entries {
		fullPath := filepath.Join(root, entry.Name())
		if pst.isIgnored(fullPath, entry.IsDir()) {
			continue
		}
		if entry.IsDir() {
			dirs = append(dirs, entry)
		} else {
			files = append(files, entry)
		}
	}

	// Sort by name
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name() < dirs[j].Name()
	})
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, d := range dirs {
		builder.WriteString(fmt.Sprintf("%s├── %s/\n", prefix, d.Name()))
		builder.WriteString(pst.buildProjectTree(filepath.Join(root, d.Name()), depth+1))
	}

	for _, f := range files {
		builder.WriteString(fmt.Sprintf("%s├── %s\n", prefix, f.Name()))
	}

	return builder.String()
}
