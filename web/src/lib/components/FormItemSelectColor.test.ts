import { Color } from "../types/Color";
import type { FormItemSelectColorAttrs } from "./FormItemSelectColor";
import { FormItemSelectColor } from "./FormItemSelectColor";

test("FormItemSelectColor", async () => {
	let value = 0;
	const attrs: FormItemSelectColorAttrs = {
		name: "test",
		oninput: (e): void => {
			value = e;
		},
		value: value,
	};
	testing.mount(FormItemSelectColor, attrs);
	const select = testing.find("#form-item-select-test");
	testing.findAll("#form-item-select-test > option", Color.values.length);
	testing.input(select, `${Color.values.indexOf("Blue")}`);
	expect(value)
		.toBe(Color.values.indexOf("Blue"));
});

