import { aesDecrypt, aesEncrypt, NewAESKey } from "./AES";

test("AES", async () => {
	const key = await NewAESKey() as string;

	expect(key.length)
		.toBe(Math.ceil(16/3) * 4);

	const input = "testing";

	let output = await aesEncrypt(key, input) as string;
	expect(output)
		.toHaveLength(Math.ceil((input.length + 12 + 4 + 12)/3) * 4); // IV + counter + append IV


	output = await aesDecrypt(key, output) as string;
	expect(output)
		.toBe(input);

	// Test Go outputs
	const goKey = "9/dd0iJw0TWcg9szIWefsA==";
	const goOutput = "tyOOaSfO6K3bx6lEVikbX7fAnIok3bu1Y9HXNPUZf3kLo3I=";

	output = await aesEncrypt(goKey, input) as string;
	expect(output)
		.toHaveLength(Math.ceil((input.length + 12 + 4 + 12)/3) * 4); // IV + counter + append IV

	output = await aesDecrypt(goKey, output) as string;
	expect(output)
		.toBe(input);

	output = await aesDecrypt(goKey, goOutput) as string;
	expect(output)
		.toBe(input);
});
