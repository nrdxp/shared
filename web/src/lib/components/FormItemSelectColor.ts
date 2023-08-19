import { AppState } from "@lib/states/App";
import m from "mithril";

import type { ColorEnum } from "../types/Color";
import { FormItem } from "./FormItem";

export interface FormItemSelectColorAttrs {
	/** Is the FormItem disabled? */
	disabled?: boolean,

	/** Override default name. */
	name?: string,

	oninput(e: ColorEnum): void,

	/** Value of the FormItem. */
	value: ColorEnum,
}

export function FormItemSelectColor (): m.Component<FormItemSelectColorAttrs> {
	return {
		view: (vnode): m.Children => {
			return m(FormItem, {
				name: vnode.attrs.name === undefined ?
					AppState.data.translations.formItemSelectColorName :
					vnode.attrs.name,
				select: {
					disabled: vnode.attrs.disabled,
					oninput: (e: string): void => {
						vnode.attrs.oninput(parseInt(e, 10));
					},
					options: AppState.data.translations.formItemSelectColorValues,
					value: vnode.attrs.value,
				},
				tooltip: AppState.data.translations.formItemSelectColorTooltip,
			});
		},
	};
}
