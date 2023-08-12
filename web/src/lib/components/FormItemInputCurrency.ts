import m from "mithril";

import type { CurrencyEnum } from "../types/Currency";
import { Currency } from "../types/Currency";
import { FormItem } from "./FormItem";

export interface FormItemInputCurrencyAttrs {
	/** Format to display Currency with. */
	format: CurrencyEnum,

	/** Name of the FormItem. */
	name: string,
	oninput(amount: number): void,

	/** Force only positive numbers. */
	onlyPositive?: boolean,

	/** Is the FormItem required? */
	required?: boolean,

	/** Start the input as a negative number. */
	startNegative?: boolean,

	/** Help tooltip. */
	tooltip: string,

	/** Current value, will be displayed as a Currency. */
	value: number,
}

export function FormItemInputCurrency (): m.Component<FormItemInputCurrencyAttrs> {
	let amountString = "";
	let zeroString = "";

	return {
		oninit: (vnode): void => {
			amountString = Currency.toString(vnode.attrs.value, vnode.attrs.format);
			zeroString = Currency.toString(0, vnode.attrs.format)
				.trim();

			if (vnode.attrs.startNegative === true && vnode.attrs.value === 0) {
				amountString = `-${amountString.trim()}`;
			}
		},
		view: (vnode): m.Children => {
			return m(FormItem, {
				input: {
					oninput: (e: string): void => {
						e = e.trim(); // eslint-disable-line no-param-reassign

						if (vnode.attrs.onlyPositive === true && e.includes("-")) {
							e = e.replace("-", ""); // eslint-disable-line no-param-reassign
						}

						if (e === `${zeroString}-`) {
							amountString = `-${zeroString}`;

							return;
						}

						amountString = "";

						if (e.startsWith("-") && e.endsWith("-")) {
							e = e.replace(/-/g, ""); // eslint-disable-line no-param-reassign
						}

						vnode.attrs.oninput(Currency.toNumber(e));
					},
					required: vnode.attrs.required === true,
					type: "text",
					value: amountString === "" ?
						Currency.toString(vnode.attrs.value, vnode.attrs.format) :
						amountString,
				},
				name: vnode.attrs.name,
				tooltip: vnode.attrs.tooltip,
			});
		},
	};
}
