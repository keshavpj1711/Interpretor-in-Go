# Setting up GO Workspace

This was a bit difficult to do as i was really confused on what was going on. I am stil a bit confused but here's what i have understood.\
To get your go program running first thing you need is:
- a go.mod file, and
- a go.work file 

## `go.mod` 
To understand what are go.mod files let us imagine a recipe for baking a cake

- `go.mod` is like a recipe book for your go program
- As you make your cake(program) you add ingredients(modules) to the recipe list for your cake\

So basically `go.mod` serves as the place where you list your required module path

### Location

`go.mod` file resides in the directory where the module is present

## `go.work`

`go.work` serves as a higher-level-config file as compared to go.mod. It's main purpose is to define and organise relationship between various **go modules**. It gives us more control over how our program runs.\
To work with different modules we use directives like `use`, `replace` and `exclude`.

- As the name suggests `use` tells us which modules to be used
- For other directives GOOGLE

### Location

`go.work` resides in the root directory of your workspace