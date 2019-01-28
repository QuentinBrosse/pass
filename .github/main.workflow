workflow "Checks" {
  on = "push"
  resolves = ["Test", "Lint"]
}

action "Test" {
  uses = "./.github/actions/go"
  args = "test"
}

action "Lint" {
  uses = "./.github/actions/go"
  args = "lint"
}
