## Rulebook - Simple Per-Project Linting
Make project-specific text rules using regular expressions.

### Example

```
$ echo "Make & use rules to avoid simple mistakes.\na = ! foobar; puts 'bar'\ndebugger;\n:not{}" | rulebook validate                                                                                                                                     |
Make & use rules to avoid simple mistakes.
a = ! foobar; puts 'bar'
- Rulebook Violation ->  {regex: !\s\w, message: Ruby: Space after !}
debugger;
- Rulebook Violation ->  {regex: debugger, message: Remove debuggers before committing}
:not{}
- Rulebook Violation ->  {regex: :not{}, message: IE8 Incompatible}
---------------------
3 Rulebook Violations
regex: !\s\w                message: Ruby: Space after !
regex: debugger             message: Remove debuggers before committing
regex: :not{}               message: IE8 Incompatible
```

## To Use
1. Make a Rulebook directory on Github.
2. Add .yaml files in it with rules.  These should be styled as such:
```
rules:
  - regex: ‘!\s\w’
    warning: “Ruby: Space after !”
  - group:
      warning: not in ie8
      regex:

        ##CSS3
        - :@namespace
        - :@keyframes
        - :@-ms-viewport
```

3. Create a ```.rulebook``` file in a project you would like to check.  This should contain the name of the rulebook you created.
```
github.com/oagr/rulebook1
```
5. On your machine, cd into the project with the .rulebook file.  Run ```rulebook book clone``` to download it to the ```.rulebooks``` directory on your machine.
6. Run ```rulebook diff``` to run the rulebook rules against your current ```git diff```, or ```echo foo | rulebook validate``` to validate any arbirary text.

## Pull Request Commenting
Rulebook can directly comment on Github Pull requests.  To do this, simply run ```rulebook comment https://github.com/org/project/pull/pr_number```.
