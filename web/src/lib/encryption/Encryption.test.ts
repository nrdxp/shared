import { IsErr } from "@lib/services/Log";

import { NewAESKey } from "./AES";
import type { EncryptedValue } from "./Encryption";
import { EncryptionTypeAES128GCM, EncryptionTypeNone, EncryptionTypeRSA2048, EncryptValue } from "./Encryption";
import { NewRSAKey } from "./RSA";

test.each([
	[
		EncryptionTypeNone,
	],
	[
		EncryptionTypeAES128GCM,
	],
	[
		EncryptionTypeRSA2048,
	],
	[
		"a",
	],
])("EncryptedValue %s", async (name) => {
	let encryptKey = "";
	let decryptKey = "";

	switch (name) {
	case EncryptionTypeAES128GCM:
		encryptKey = await NewAESKey() as string;
		decryptKey = encryptKey;
		break;
	case EncryptionTypeRSA2048:
		({ publicKey: encryptKey, privateKey: decryptKey } = await NewRSAKey() as {
			privateKey: string,
			publicKey: string,
		});
	}

	const v = "testing123";

	const ev = await EncryptValue(name, encryptKey, v) as EncryptedValue;

	if (name !== "a") {
		expect(ev.encryption)
			.toBe(name);
	}

	if (name === "a") {
		expect(IsErr(ev))
			.toBeTruthy();
	} else {
		expect(await ev.decrypt(decryptKey))
			.toBe(v);
	}
});
