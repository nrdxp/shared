import { StringCapitalize } from "@lib/utilities/StringCapitalize";

/* eslint-disable sort-keys*/
export const ColorEnum: {
	[key: string]: {
		dark: string,
		light: string,
	},
} = {
	red: {
		dark: "#ff5252",
		light: "#d50000",
	},
	pink: {
		dark: "#ff4081",
		light: "#f50057",
	},
	orange: {
		dark: "#ffab40",
		light: "#ff6d00",
	},
	yellow: {
		dark: "#ffd740",
		light: "#ffab00",
	},
	green: {
		dark: "#69f0ae",
		light: "#00c853",
	},
	teal: {
		dark: "#64ffda",
		light: "#00bfa5",
	},
	blue: {
		dark: "#448aff",
		light: "#2962ff",
	},
	indigo: {
		dark: "#536dfe",
		light: "#6366f1",
	},
	purple: {
		dark: "#7c4dff",
		light: "#6200ea",
	},
	brown: {
		dark: "#795548",
		light: "#5d4037",
	},
	black: {
		dark: "#c7c7c7",
		light: "#101010",
	},
	gray: {
		dark: "#dfe6e9",
		light: "#b2bec3",
	},
	white: {
		dark: "#ffffff",
		light: "#636363",
	},
};
/* eslint-enable sort-keys*/

export const Color = {
	content: {
		black: "#1f2937",
		white: "#ebecf0",
	},
	contentColor: (color: string): string => {
		if (!color.startsWith("#")) {
			return Color.content.black;
		}

		const rgb = parseInt(color.substring(1), 16); // convert rrggbb to decimal
		const r = rgb >> 16 & 0xff; // extract red
		const g = rgb >>  8 & 0xff; // extract green
		const b = rgb >>  0 & 0xff; // extract blue

		const luma = 0.2126 * r + 0.7152 * g + 0.0722 * b; // per ITU-R BT.709

		if (luma < 40) {
			return Color.content.black;
		}

		return Color.content.white;
	},
	toHex: (color: string, darkMode?: boolean): string => {
		if (Object.hasOwn(ColorEnum, color)) {
			return ColorEnum[color][darkMode === true ?
				"dark" :
				"light"
			];
		}

		return color;
	},
	values: [
		"Default",
		...Object.keys(ColorEnum)
			.map((key) => {
				StringCapitalize(key);
			}),
	],
};
