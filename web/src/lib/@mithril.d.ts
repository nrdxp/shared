import type m from "mithril";

declare module "mithril" {
	type Event<T> = T & { // eslint-disable-line @typescript-eslint/no-type-alias
		redraw: boolean,
		target: HTMLElement,
	};
	export interface Input {
		data: string,
		inputType: string,
		key: string,
		preventDefault(): void,
		target: HTMLInputElement,
	}
	interface VnodeDOMHTML<Attrs = {}, State extends m.Lifecycle<Attrs, State> = {}> extends m.Vnode<Attrs, State> {
		dom: HTMLElement,
		domSize?: number,
	}
}
