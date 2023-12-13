import { Color, ColorEnum } from "./Color";

describe("Color", () => {
	test("contentColor", () => {
		expect(Color.contentColor(ColorEnum.orange.light))
			.toBe(Color.content.black)

		expect(Color.contentColor(ColorEnum.orange.dark))
			.toBe(Color.content.white)
		expect(Color.contentColor(ColorEnum.blue.dark))
			.toBe(Color.content.white)
	})
})
