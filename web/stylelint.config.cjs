module.exports = {
	extends: "stylelint-config-recommended",
	plugins: [
		"stylelint-no-unsupported-browser-features",
	],
	rules: {
		"indentation": "tab",
		"plugin/no-unsupported-browser-features": [
			true,
			{
				ignore: [
					"css-appearance",
					"css-resize",
					"css-sticky",
					"intrinsic-width",
					"multicolumn",
				],
				severity: "error",
			},
		],
		"selector-type-no-unknown": [
			true,
			{
				ignore: [
					"custom-elements",
				],
			},
		],
		"string-quotes": "double",
	},
};
