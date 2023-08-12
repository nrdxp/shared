export const Color = {
	getValue: (value: number): string => {
		if (value === undefined || value > Color.values.length - 1) {
			return Color.values[0];
		}

		return Color.values[value];
	},
	toString: (color: ColorEnum, darkMode?: boolean, content?: boolean): string => {
		return `var(--color_${Color.values[color].toLowerCase()}${darkMode === undefined ?
			"" :
			darkMode ?
				"-dark" :
				"-light"}${content === true ?
			"-content" :
			""})`;
	},
	values: [
		"Default",
		"Red",
		"Pink",
		"Orange",
		"Yellow",
		"Green",
		"Teal",
		"Blue",
		"Indigo",
		"Purple",
		"Brown",
		"Black",
		"Gray",
		"White",
	],
};

export enum ColorEnum {
	Default,
	Red,
	Pink,
	Orange,
	Yellow,
	Green,
	Teal,
	Blue,
	Indigo,
	Purple,
	Brown,
	Black,
	Gray,
	White,
}
