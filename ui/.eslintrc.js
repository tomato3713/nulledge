module.exports = {
	root: true,
	extends: [
		"next/core-web-vitals",
		"plugin:import/recommended",
		"airbnb-base",
		"prettier",
	],
	rules: {
		"import/prefer-default-export": "off",
	},
};
