package commit

var defaultConfig = `{
	"message": {
		"items": [
			{
				"name": "type",
				"desc": "Select the type of change that you're committing:",
				"form": "select",
				"options": [
					{ "name": "✨feat", "desc": "✨feat: A new feature" },
					{ "name": "🐛fix", "desc": "🐛fix: A bug fix" },
					{ "name": "📚docs", "desc": "📚docs: Documentation only changes" },
					{
					  "name": "style",
					  "desc":
						"style: Changes that do not affect the meaning of the code\n            (white-space, formatting, missing semi-colons, etc)"
					},
					{
					  "name": "🔨refactor",
					  "desc": "🔨refactor: A code change that neither fixes a bug nor adds a feature"
					},
					{
					  "name": "🐎perf",
					  "desc": "🐎perf: A code change that improves performance"
					},
					{ "name": "🚨test", "desc": "🚨test: Adding missing tests" },
					{
					  "name": "♿chore",
					  "desc":
						"♿chore: Changes to the build process or auxiliary tools\n            and libraries such as documentation generation"
					},
					{ "name": "⏪revert", "desc": "⏪revert: Revert to a commit" },
					{ "name": "🚧WIP", "desc": "🚧WIP: Work in progress" }
				],
				"required": true
			},
			{
				"name": "scope",
				"desc": "Scope. Could be anything specifying place of the commit change (users, db, poll):",
				"form": "input"
			},
			{
				"name": "subject",
				"desc": "Subject. Concise description of the changes. Imperative, lower case and no final dot:",
				"form": "input",
				"required": true
			},
			{
				"name": "body",
				"desc": "Body. Motivation for the change and contrast this with previous behavior:",
				"form": "multiline"
			},
			{
				"name": "footer",
				"desc": "Footer. Information about Breaking Changes and reference issues that this commit closes:",
				"form": "multiline"
			}
		],
		"template": "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"
	}
}
`
