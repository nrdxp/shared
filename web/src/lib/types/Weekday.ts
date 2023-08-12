import { AppState } from "@lib/states/App";

export const Weekday = {
	getValues (formatWeek8601: boolean): string[] {
		return formatWeek8601 ?
			AppState.preferences().translations.formRecurrenceWeekdays :
			[
				AppState.preferences().translations.formRecurrenceWeekdays[6],
				...AppState.preferences().translations.formRecurrenceWeekdays.slice(0, 6),
			];
	},
	values: [
		"",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	],
};

export enum WeekdayEnum {
	None,
	Monday,
	Tuesday,
	Wednesday,
	Thursday,
	Friday,
	Saturday,
	Sunday,
}
