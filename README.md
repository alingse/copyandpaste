# copyandpaste
Do not do copy and paste things.

## Post Actions

```bash
gh search repos util --limit 100 --language=go --json url --jq '.[]|.url' | xargs -I {} gh workflow run .github/workflows/check-any.yaml -F repo_url={}
```