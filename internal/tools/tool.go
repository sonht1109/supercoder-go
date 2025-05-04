package tools

type Tool interface {
	Execute(arguments map[string]any) string
}

const CodeEditToolName = "file_edit"

const CodeEditToolDescription = `Edit/update a code file in the repository. You can create file if file does not exist yet. Provide the file path and the new content for the file. Arguments: {"filePath": "<file-path>", "content": "<new-content>"}`

const FileReadToolName = "file_read"

const FileReadToolDescription = `Read a file to understand its content. Use this tool to read a file and understand its content. Arguments: {"filePath": "<file-path-and-name>"}`

const ProjectStructureToolName = "project_structure"

const ProjectStructureToolDescription = `Analyze the project structure. Use this tool to analyze the project structure. Arguments: {}`

const CodeSearchToolName = "code_search"

const CodeSearchToolDescription = `Search for code across the project. The query parameter should be a regular expression. Arguments: {"query": "<search-query>"}`

const WebSearchToolName = "web_search"

const WebSearchToolDescription = `Perform web search using SearxNG. Use this when you need to find information that is not in the codebase or when you need to find a specific library or tool. Arguments: {"query": "<search-query>", "limit": <max-results>}. Default limit is 5.`

const URLFetchToolName = "url_fetch"

const URLFetchToolDescription = `Fetch the content of a URL. Use this when you need to fetch the content of a URL. Arguments: {"url": "<target-url>"}`
