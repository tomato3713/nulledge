module.exports = {
  root: true,
  extends: [
    "next/core-web-vitals",
    "plugin:import/recommended",
    "airbnb-base",
    "prettier",
  ],
  parserOptions: {
    project: "./tsconfig.json",
  },
  rules: {
    "import/prefer-default-export": "off",
  },
};
