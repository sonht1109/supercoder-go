package tools

type Tool interface {
	Execute(arguments map[string]any) string
}

const CodeEditToolName = "file_edit"

const CodeEditToolDescription = `
  Edit a code file in the repository. Provide the file path and the new content for the file. Arguments: {"filePath": "<file-path>", "content": "<new-content>"}.
  Use the "file_read" tool to read file content if you are not sure about file content before editing.
  Example:
  <@TOOL>
  {"name": "file_edit", "arguments": {"filePath": "<file-path>", "content": "<new-content>"}, "id": <function-id-in-uuid-v4-format>}
  </@TOOL>`

const FileCreateToolName = "file_create"

const FileCreateToolDescription = `Create a new file in the repository and add content into it. NO need to check or read file before creating. Arguments: {"filePath": "<file-path>", "content": "<new-content>"}.
  Example:
  <@TOOL>
  {"name": "file_create", "arguments": {"filePath": "<file-path>", "content": "<new-content>"}, "id": <function-id-in-uuid-v4-format>}
  </@TOOL>
`

const FileReadToolName = "file_read"

const FileReadToolDescription = `Read a file to understand its content. Use this tool to read a file and understand its content. Arguments: {"filePath": "<file-path-and-name>"}.
  Example:
  <@TOOL>
  {"name": "file_read", "arguments": {"filePath": "<file-path-and-name>"}, "id": <function-id-in-uuid-v4-format>}
  </@TOOL>
`

const ProjectStructureToolName = "project_structure"

const ProjectStructureToolDescription = `Analyze the project structure. Use this tool to analyze the project structure. Arguments: {}.
  Example:
  <@TOOL>
  {"name": "project_structure", "arguments": {}, "id": <function-id-in-uuid-v4-format>}
  </@TOOL>
`

const CodeSearchToolName = "code_search"

const CodeSearchToolDescription = `Search for code across the project. The query parameter should be a regular expression. Arguments: {"query": "<search-query>"}.
  Example:
  <@TOOL>
  {"name": "code_search", "arguments": {"query": "<search-query>"}, "id": <function-id-in-uuid-v4-format>}
  </@TOOL>
`

const WebSearchToolName = "web_search"

const WebSearchToolDescription = `
  Perform web search using SearxNG. Use this when you need to find information that is not in the codebase or when you need to find a specific library or tool. Arguments: {"query": "<search-query>", "limit": <max-results>}. Default limit is 5.
  Example:
  <@TOOL>
  {"name": "web_search", "arguments": {"query": "<search-query>", "limit": <max-results>}, "id": <function-id-in-uuid-v4-format>}
  </@TOOL>
`

const URLFetchToolName = "url_fetch"

const URLFetchToolDescription = `Fetch the content of a URL. Use this when you need to fetch the content of a URL. Arguments: {"url": "<target-url>"}.
  Example:
  <@TOOL>
  {"name": "url_fetch", "arguments": {"url": "<target-url>"}, "id": <function-id-in-uuid-v4-format>}
  </@TOOL>
`
