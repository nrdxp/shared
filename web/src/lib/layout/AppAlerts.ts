import "./AppAlerts.css";

import m from "mithril";

import { AppState } from "../states/App";
import { Animate, Animation } from "../utilities/Animate";
import { StringToID } from "../utilities/StringToID";

export interface AppAlert {
	/** Optional list of buttons to render with Alert. */
	actions?: {
		/** Button name. */
		name: string,
		onclick (): Promise<void>,
	}[],

	/** Message for the alert.  Will be deduplicated. */
	message: string,

	/** Persist the alert instead of removing it eventually. */
	persist?: boolean,
}

export function AppAlerts (): m.Component {
	return {
		view: (): m.Children => {
			return m("div.AppAlerts", AppState.getLayoutAppAlerts()
				.map((alert) => {
					return m(`p.AppAlerts__alert.${Animate.class(Animation.FromTop)}`, {
						onbeforeremove: Animate.onbeforeremove(Animation.FromTop),
					}, [
						m("span", alert.message),
						alert.actions === undefined ?
							[] :
							m("div.AppAlerts__actions", alert.actions.map((action) => {
								return m("span", {
									id: `button${StringToID(alert.message)}${StringToID(action.name)}`,
									onclick: async () => {
										if (action.onclick === undefined) {
											AppState.clearLayoutAppAlert(alert.message);
										} else {
											return action.onclick();
										}
									},
								}, action.name);
							})),
					]);
				}));
		},
	};
}
