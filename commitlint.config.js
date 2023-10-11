module.exports = {
  extends: [
    '@commitlint/config-conventional',
    '@commitlint/config-nx-scopes'
  ],
  rules: {
    "scope-empty": [2, 'never'],
    "subject-case": [0, 'always'],
  }
}
