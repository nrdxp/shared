import { aesDecrypt, aesEncrypt } from "@lib/encryption/AES";

import { NewPBDKF2AES128Key } from "./PBDKF2";

test("PBDKF2", async () => {
	const key = await NewPBDKF2AES128Key("password", "12345") as string;
	expect(key)
		.toBe("Zg8nzLfVZ6/wDjDFgIXWew==");

	expect(await aesDecrypt(key, await aesEncrypt(key, "secret") as string))
		.toBe("secret");
});
