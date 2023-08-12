import type { Err } from "@lib/services/Log";
import { IsErr, NewErr } from "@lib/services/Log";

import { aesDecrypt, aesEncrypt } from "./AES";
import { rsaDecrypt, rsaEncrypt } from "./RSA";

export type EncryptionType = string; // eslint-disable-line @typescript-eslint/no-type-alias

export const EncryptionTypeNone = "none" as EncryptionType;
export const EncryptionTypeAES128GCM = "aes128gcm" as EncryptionType;
export const EncryptionTypeRSA2048 = "rsa2048" as EncryptionType;

export interface EncryptedValue {
	/** Ciphertext is the encrypted part of a value. */
	ciphertext: string,

	/** Encryption is the encryption type of a value. */
	encryption: EncryptionType,
}

export class EncryptedValue {
	ciphertext: string;
	encryption: EncryptionType;

	constructor (ciphertext: string, encryption: EncryptionType) {
		this.ciphertext = ciphertext;
		this.encryption = encryption;
	}

	async decrypt (key: string): Promise<string | Err> {
		switch (this.encryption) {
		case EncryptionTypeNone:
			return this.ciphertext;
		case EncryptionTypeAES128GCM:
			return aesDecrypt(key, this.ciphertext);
		case EncryptionTypeRSA2048:
			return rsaDecrypt(key, this.ciphertext);
		}

		return NewErr("decrypt: unknown encryption", "Error decrypting value");
	}

	string (): string {
		return `${this.encryption}$${this.ciphertext}`;
	}
}

export function ParseEncryptedValue (s: string): EncryptedValue | Err {
	const r = s.split("$");

	let e = "";

	if (r.length === 2) {
		switch (r[0]) {
		case EncryptionTypeNone:
			e = EncryptionTypeNone;
			break;
		case EncryptionTypeAES128GCM:
			e = EncryptionTypeAES128GCM;
			break;
		case EncryptionTypeRSA2048:
			e = EncryptionTypeRSA2048;
			break;
		}

		if (e !== "") {
			return new EncryptedValue(r[1], e);
		}
	}

	return NewErr("ParseEncryptedValue: unknown encryption", "Error reading encrypted value");
}

export async function EncryptValue (e: EncryptionType, key: string, value: string): Promise<EncryptedValue | Err> {
	let c = "" as string | Err;

	switch (e) {
	case EncryptionTypeNone:
		c = value;
		break;
	case EncryptionTypeAES128GCM:
		c = await aesEncrypt(key, value);
		break;
	case EncryptionTypeRSA2048:
		c = await rsaEncrypt(key, value);
		break;
	}

	if (IsErr(c)) {
		return c;
	}

	if (c !== "") {
		return new EncryptedValue(c, e);
	}

	return NewErr("EncryptValue: unknown encryption", "Error encrypting value");
}
