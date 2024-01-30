# mustcalls

mustcalls checks functions that do not conform to specified rules for function calls.

## Installation

```console
go install github.com/nametake/mustcalls/cmd/mustcalls@latest
```

## Usage

```console
go vet -vettool=`which mustcalls` -mustargs.config=$(pwd)/config.yaml .
```

## Rule Definition

Rules are configured in a YAML file.

The overall structure of the rules that can be configured is shown below. Parameters without description are all optional.

```yaml
---
# Multiple rules can be configured. Each rule is independent.
rules:
  # These are the rules for function calls. Multiple configurations are possible within each rule.
    # If not following all argument rules, an error will occur.
  - funcs:
        # 'name' is the name of the function that must be called.
        # This is a required field.
      - name: mustCalls
    # These are the patterns for the functions targeted by the rule.
    # Multiple patterns can be specified as a list.
    # Patterns support regular expressions.
    # Each pattern is an AND condition, and the list of patterns is an OR condition.
    file_patterns: # File patterns.
      - usecase/tenant_.*.go
    ignore_file_patterns: # Patterns to ignore files.
      - .*_gen.go
    func_patterns: # Function name patterns.
      - Create.*
      - Upsert.*
    ignore_func_patterns: # Patterns to ignore function names.
      - ^New.*
    recv_patterns: # Receiver patterns.
      - ^Tenant.*Usecase$
    ignore_recv_patterns: # Patterns to ignore receivers.
      - ^Debug.*Usecase$
```
