workflow "Test" {
  on = "push"
  resolves = ["Run Tests"]
}

action "Run Tests" {
  uses = "./.github/actions/tests"
}
