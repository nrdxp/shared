import "./Table.css";

import { Clone } from "@lib/utilities/Clone";
import m from "mithril";

import { AppState } from "../states/App";
import { Icons } from "../types/Icons";
import { Animate, Animation } from "../utilities/Animate";
import { StringToID } from "../utilities/StringToID";
import type { ButtonAttrs } from "./Button";
import { Button } from "./Button";
import type { ButtonArrayAttrs } from "./ButtonArray";
import { TableData } from "./TableData";
import type { TableColumnAttrs, TableLayoutAttrs } from "./TableHeader";
import { TableHeader } from "./TableHeader";
import type { TableSearchAttrs } from "./TableSearch";
import { TableSearch } from "./TableSearch";
import type { TitleAttrs } from "./Title";
import { Title } from "./Title";
import { Toolbar } from "./Toolbar";

export interface TableAttrs extends TableLayoutAttrs {
	/** A list of action buttons to display at the top of the Table. */
	actions: ButtonAttrs[],

	/** Class to append to the table. */
	class?: string,
	editOnclick?(form: unknown): Promise<void> | void,

	/** A list of filter buttons to display at the top of the table. */
	filters: ButtonArrayAttrs[],
	getKey?(object?: any): string, // eslint-disable-line @typescript-eslint/no-explicit-any

	/** Is the Table loaded?  If not, display placeholder. */
	loaded: boolean,

	/** Disable showing new buttons if the Table is empty. */
	noNewButton?: boolean,

	/** TableSearch settings. */
	search?: TableSearchAttrs,

	/** Don't let users show/hide columns. */
	staticColumns?: boolean,

	/** Title for the Table. */
	title?: TitleAttrs,
}

let route = "";

const scroll: {
	[route: string]: number,
} = {};

export function Table (): m.Component<TableAttrs> {
	return {
		view: (vnode): m.Children => {
			return m("div.Table__container", {
				class: vnode.attrs.class,
				id: `table-container${StringToID(vnode.attrs.id)}`,
			}, [
				m(Title, vnode.attrs.title === undefined ?
					{} :
					vnode.attrs.title),
				m(Toolbar, {
					actions: vnode.attrs.actions,
					filters: [
						...vnode.attrs.filters,
						...vnode.attrs.staticColumns === true ?
							[] :
							[
								{
									name: AppState.preferences().translations.tableShowColumns,
									onclick: (e:string): void => {
										if (vnode.attrs.tableColumnsNameEnabled()[e] === undefined) {
											vnode.attrs.tableColumnsNameEnabled()[e] = "";
										} else {
											delete vnode.attrs.tableColumnsNameEnabled()[e];
										}
									},
									selected: (): string[] => {
										return Object.keys(vnode.attrs.tableColumnsNameEnabled());
									},
									value: vnode.attrs.tableColumns.reduce((arr, column) => {
										if (column.name !== "" && column.noToggle !== true) {
											arr.push({
												id: column.property,
												name: column.name,
											});
										}
										return arr;
									}, [] as {
										id: string, // eslint-disable-line jsdoc/require-jsdoc
										name: string, // eslint-disable-line jsdoc/require-jsdoc
									}[]),
								},
							],
					],
				}),
				vnode.attrs.search === undefined ?
					[] :
					m(TableSearch, vnode.attrs.search),
				[
					m("div.Table__div", m("table.Table", {
						id: `table${StringToID(vnode.attrs.id)}`,
					}, [
						m(TableHeader, {
							data: vnode.attrs.data,
							id: vnode.attrs.id,
							noFilters: vnode.attrs.noFilters,
							sort: vnode.attrs.sort,
							tableColumns: vnode.attrs.tableColumns,
							tableColumnsNameEnabled: vnode.attrs.tableColumnsNameEnabled,
						}),
						m("tbody", {
							oncreate: (tvnode): void => {
								route = m.route.get();

								if (scroll[route] !== undefined) {
									tvnode.dom.scrollTop = scroll[route];
								}
							},
							onremove: (tvnode): void => {
								scroll[route] = tvnode.dom.scrollTop;
							},
						}, vnode.attrs.loaded ?
							vnode.attrs.data === null ?
								[] :
								vnode.attrs.data
									.map((data: any, i: number): m.Children => { // eslint-disable-line @typescript-eslint/no-explicit-any
										const id = StringToID(vnode.attrs.getKey === undefined ?
											data.id :
											vnode.attrs.getKey(data), true);

										return m(`tr.${Animate.class(Animation.Fade)}`, {
											class: vnode.attrs.editOnclick !== undefined && AppState.isSessionOnline() ?
												"Table__edit" :
												undefined,
											id: `table${StringToID(vnode.attrs.id)}-row_${id}`,
											key: id,
											onclick: vnode.attrs.editOnclick !== undefined && AppState.isSessionOnline() ?
												async (): Promise<void> => {
													if (vnode.attrs.editOnclick !== undefined) {
														await vnode.attrs.editOnclick(vnode.attrs.data[i]);
													}
												} :
												undefined,
										}, [
											vnode.attrs.tableColumns.map((column: TableColumnAttrs): m.Children => {
												return Object.keys(vnode.attrs.tableColumnsNameEnabled())
													.includes(column.property) ?
													m(TableData, {
														checkboxOnclick: async () => {
															if (column.checkboxOnclick !== undefined) {
																return column.checkboxOnclick(Clone(data));
															}
														},
														collapsedNames: column.collapsedNames,
														currencyFormat: column.currencyFormat,
														data: data,
														formatter: column.formatter,
														id: `${id}-${column.property}`,
														linkExternal: column.linkExternal,
														linkFormatter: column.linkFormatter,
														linkRequireOnline: column.linkRequireOnline,
														name: column.name,
														onbeforeremove: Animate.onbeforeremove(Animation.Fade),
														onclick: column.onclick,
														ondragend: column.ondragend,
														permitted: () => {
															if (column.permitted !== undefined) {
																return column.permitted(vnode.attrs.data[i]);
															}

															return false;
														},
														positive: column.positive,
														property: column.property,
														render: column.render,
														rightAlign: column.rightAlign,
														type: column.typeCheck === undefined ?
															column.type :
															column.typeCheck(data),
													}) :
													[];
											}),
										]);
									}) :
							[
								...Array(5),
							].map(() => {
								return m(`tr.Table__loading.${Animate.class(Animation.Pulse)}`, vnode.attrs.tableColumns.map((column: TableColumnAttrs): m.Children => {
									return Object.keys(vnode.attrs.tableColumnsNameEnabled())
										.includes(column.property) ?
										m("td.") :
										[];
								},
								));
							}),
						),
						vnode.attrs.loaded && (vnode.attrs.data === null || vnode.attrs.data.length === 0) ?
							m("div.Table__none", [
								m("i.Table__none-icon", Icons.NotFound),
								m("p.Table__label", {
									id: "table-no-data-message",
								}, AppState.preferences().translations.tableNothingFound),
								vnode.attrs.noNewButton === true ?
									[] :
									AppState.data.layoutAppToolbarActionButtons.map((action) => {
										if (action.permitted && AppState.isSessionOnline() || !action.requireOnline) {
											return m(Button, {
												icon: Icons.Add,
												name: `${AppState.preferences().translations.actionNew} ${action.name}`,
												onclick: async (): Promise<void> => {
													return action.onclick!(); // eslint-disable-line @typescript-eslint/no-non-null-assertion
												},
												permitted: true,
												primary: true,
												requireOnline: true,
											});
										}

										return [];
									}),
							]) :
							[],
					])),
				],
			],
			);
		},
	};
}
