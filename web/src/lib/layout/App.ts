import "./App.css";

import { Drag } from "@lib/utilities/Drag";
import m from "mithril";
import type Stream from "mithril/stream";

import { Button } from "../components/Button";
import { Markdown } from "../components/Markdown";
import { Log } from "../services/Log";
import { AppState } from "../states/App";
import { DisplayEnum } from "../types/Display";
import { Animate, Animation } from "../utilities/Animate";
import { SetClass } from "../utilities/SetClass";
import { AppAlerts } from "./AppAlerts";
import { AppForm } from "./AppForm";
import { AppHeader } from "./AppHeader";
import { AppHelp } from "./AppHelp";
import type { AppLogoAttrs, AppLogoComponentAttrs } from "./AppLogo";
import type { AppMenuComponent } from "./AppMenu";
import { AppMenu } from "./AppMenu";
import { AppToolbar } from "./AppToolbar";
import type { AppToolbarActionAttrs } from "./AppToolbarAction";

export interface AppAttrs {
	/** A component to render a Contact Us block below appropriate items. */
	contactUs(): m.Component,

	/** Don't pad the width of children. */
	fullWidth?: boolean,

	/** Hide the header and menu. */
	hideHeader?: boolean,

	/** A component to render a logo within items. */
	logo(): m.Component<AppLogoComponentAttrs>,

	/** A list of Menu Components to render. */
	menuComponents: Stream<AppMenuComponent[]>,

	/** Whether to redirect page to signin. */
	public?: boolean,

	/** Handlers for App Search.*/
	searcher: AppLogoAttrs["searcher"],

	/** A list of toolbar actions to display across all pages. */
	toolbarActions(): AppToolbarActionAttrs,
}

export function App (): m.Component<AppAttrs> {
	return {
		oncreate: async (vnode): Promise<void> => {
			document.onkeydown = (e): void => {
				switch (e.key) {
				case "Escape":
					if (AppState.getLayoutAppForm().component !== null) {
						AppState.setLayoutAppForm();
					}

					if (AppState.isLayoutAppHelpOpen()) {
						AppState.toggleLayoutAppHelpOpen(false);
					}
					break;
				}

				if (!document.activeElement?.tagName.toLowerCase()
					.match(/input|select|textarea/)) {
					switch (e.key) {
					case "a":
						AppState.setComponentsDropdownMenu("app-toolbar-action", 0);
						break;
					case ",":
						AppState.toggleLayoutAppMenuOpen();
					}
				}
			};

			await AppState.oncreate();

			if (!AppState.isSessionAuthenticated() && vnode.attrs.public !== true) {
				m.route.set("/signin", {}, {
					replace: true,
				});
			}
		},
		view: (vnode): m.Children => {
			Log.debug("Redrawing");

			return m("main.App", {
				class: vnode.attrs.hideHeader === true ?
					"App--primary" :
					undefined,
				id: "app",
				key: m.route.get(),
				onclick: (e: m.Event<MouseEvent>) => {
					e.redraw = false;

					if (AppState.getComponentsDropdownMenu() !== "") {
						AppState.setComponentsDropdownMenu("", 0);
						m.redraw();
					}
				},
				ondragover: (e: m.Event<DragEvent>) => {
					Drag.setXY(e.clientX, e.clientY);
				},
				style: {
					"--border": AppState.preferences().darkMode ?
						"var(--border_dark)" :
						"var(--border_light)",
					"--border_invert": AppState.preferences().darkMode ?
						"var(--border_light)" :
						"var(--border_dark)",
					"--color_accent": `var(--color${AppState.preferences().styleAccent}${AppState.preferences().styleMode})`,
					"--color_accent-content": `var(--color${AppState.preferences().styleAccent}${AppState.preferences().styleMode}-content)`,
					"--color_accent-invert": `var(--color${AppState.preferences().styleAccent}${AppState.preferences().styleModeInvert})`,
					"--color_base-1": `var(--color_base${AppState.preferences().styleMode}-1)`,
					"--color_base-2": `var(--color_base${AppState.preferences().styleMode}-2)`,
					"--color_base-3": `var(--color_base${AppState.preferences().styleMode}-3)`,
					"--color_black": `var(--color_black${AppState.preferences().styleMode})`,
					"--color_black-content": `var(--color_black${AppState.preferences().styleMode}-content)`,
					"--color_blue": `var(--color_blue${AppState.preferences().styleMode})`,
					"--color_blue-content": `var(--color_blue${AppState.preferences().styleMode}-content)`,
					"--color_brown": `var(--color_brown${AppState.preferences().styleMode})`,
					"--color_brown-content": `var(--color_brown${AppState.preferences().styleMode}-content)`,
					"--color_content": `var(--color_content${AppState.preferences().styleMode})`,
					"--color_content-invert": `var(--color_content${AppState.preferences().styleModeInvert})`,
					"--color_gray": `var(--color_gray${AppState.preferences().styleMode})`,
					"--color_gray-content": `var(--color_gray${AppState.preferences().styleMode}-content)`,
					"--color_green": `var(--color_green${AppState.preferences().styleMode})`,
					"--color_green-content": `var(--color_green${AppState.preferences().styleMode}-content)`,
					"--color_indigo": `var(--color_indigo${AppState.preferences().styleMode})`,
					"--color_indigo-content": `var(--color_indigo${AppState.preferences().styleMode}-content)`,
					"--color_negative": `var(--color${AppState.preferences().styleNegative}${AppState.preferences().styleMode})`,
					"--color_orange": `var(--color_orange${AppState.preferences().styleMode})`,
					"--color_orange-content": `var(--color_orange${AppState.preferences().styleMode}-content)`,
					"--color_pink": `var(--color_pink${AppState.preferences().styleMode})`,
					"--color_pink-content": `var(--color_pink${AppState.preferences().styleMode}-content)`,
					"--color_positive": `var(--color${AppState.preferences().stylePositive}${AppState.preferences().styleMode})`,
					"--color_primary": `var(--color${AppState.preferences().stylePrimary}${AppState.preferences().styleMode})`,
					"--color_primary-content": `var(--color${AppState.preferences().stylePrimary}${AppState.preferences().styleMode}-content)`,
					"--color_primary-invert": `var(--color${AppState.preferences().stylePrimary}${AppState.preferences().styleModeInvert})`,
					"--color_purple": `var(--color_purple${AppState.preferences().styleMode})`,
					"--color_purple-content": `var(--color_purple${AppState.preferences().styleMode}-content)`,
					"--color_red": `var(--color_red${AppState.preferences().styleMode})`,
					"--color_red-content": `var(--color_red${AppState.preferences().styleMode}-content)`,
					"--color_secondary": `var(--color${AppState.preferences().styleSecondary}${AppState.preferences().styleMode})`,
					"--color_secondary-content": `var(--color${AppState.preferences().styleSecondary}${AppState.preferences().styleMode}-content)`,
					"--color_secondary-invert": `var(--color${AppState.preferences().styleSecondary}${AppState.preferences().styleModeInvert})`,
					"--color_teal": `var(--color_teal${AppState.preferences().styleMode})`,
					"--color_teal-content": `var(--color_teal${AppState.preferences().styleMode}-content)`,
					"--color_white": `var(--color_white${AppState.preferences().styleMode})`,
					"--color_white-content": `var(--color_white${AppState.preferences().styleMode}-content)`,
					"--color_yellow": `var(--color_yellow${AppState.preferences().styleMode})`,
					"--color_yellow-content": `var(--color_yellow${AppState.preferences().styleMode}-content)`,
					"--width_Alerts": vnode.attrs.hideHeader !== true && AppState.isLayoutAppMenuOpen() && AppState.getSessionDisplay() >= DisplayEnum.XLarge ?
						"calc(100% + var(--width_AppMenu))" :
						"100%",
				},
			}, [
				vnode.attrs.hideHeader === true || ! AppState.isSessionAuthenticated() ?
					[] :
					m(AppHeader, {
						logo: vnode.attrs.logo,
					}),
				m("section#app-section.App__section", [
					vnode.attrs.hideHeader === true || ! AppState.isSessionAuthenticated() ?
						[] :
						m(AppMenu, {
							components: vnode.attrs.menuComponents(),
							contactUs: vnode.attrs.contactUs,
							logo: {
								logo: vnode.attrs.logo,
								searcher: vnode.attrs.searcher,
							},
						}),
					m(AppHelp),
					m(AppForm),
					m(AppAlerts),
					m("div.App__children", {
						class: SetClass({
							"App__children--full": vnode.attrs.hideHeader === true,
							"App__children--full-width": vnode.attrs.fullWidth === true,
						}),
					}, [
						AppState.motd() !== undefined && !AppState.isSessionHideMOTD() && vnode.attrs.hideHeader !== true ?
							m("p.App__motd#motd", {
								onbeforeremove: Animate.onbeforeremove(Animation.Fade),
							}, [
								m(Markdown, {
									value: AppState.motd(),
								}),
								m(Button, {
									name: AppState.data.translations.actionDismiss,
									onclick: async (): Promise<void> => {
										return new Promise((resolve) => {
											AppState.setSessionHideMOTD();
											return resolve();
										});
									},
									permitted: true,
									requireOnline: false,
								}),
							],
							) :
							[],
						vnode.attrs.hideHeader === true ?
							[] :
							m(AppToolbar, {
								actions: vnode.attrs.toolbarActions,
							}),
						vnode.children,
					]),
				]),
			]);
		},
	};
}
