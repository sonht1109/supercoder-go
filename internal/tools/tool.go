package tools

type Tool interface {
	Execute(arguments string) string
}

const CodeEditToolName = "code_edit"

const CodeEditToolDescription = `Edit/update a code file in the repository. Provide the file path and the new content for the file. Arguments: {"filepath": "<file-path>", "content": "<new-content>"}`

const FileReadToolName = "file_read"

const FileReadToolDescription = `Read a file to understand its content. Use this tool to read a file and understand its content. Arguments: {\"filePath\": \"<file-path-and-name>\"}`
