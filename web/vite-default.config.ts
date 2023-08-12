/* eslint-disable camelcase */

import path from "path";
import type { UserConfigExport } from "vite";
import { defineConfig, loadEnv } from "vite";
import { imagetools } from "vite-imagetools";
import type { VitePWAOptions } from "vite-plugin-pwa";
import { VitePWA } from "vite-plugin-pwa";

export default (name: string, port: number, pwa: Partial<VitePWAOptions>): UserConfigExport => {
	return defineConfig(({ mode }) => {
		const env = {
			...process.env,
			...loadEnv(mode, process.cwd()),
		};

		return {
			build: {
				emptyOutDir: true,
				outDir: `../../dist/${name}`,
				sourcemap: env.NODE_ENV === "development",
			},
			define: {
				"process.env": {
					BUILD_VERSION_HOMECHART: env.BUILD_VERSION_HOMECHART,
					NODE_ENV: env.NODE_ENV,
					PUPPETEER_URL_HOMECHART: env.PUPPETEER_URL_HOMECHART,
				},
			},
			plugins: [
				imagetools(),
				pwa === undefined ?
					[] :
					VitePWA(pwa),
			],
			publicDir: "static",
			resolve: {
				alias: {
					"./Worker?worker": path.resolve(__dirname, "./src/homechart/worker/Worker.ts"), // Needed for service worker to resolve the web worker correctly
					"@lib": path.resolve(__dirname, "./src/lib"),
					"events": "events", // Needed for xml2js
					"stream": "stream-browserify", // Needed for xml2js
				},
			},
			root: `src/${name}`,
			server: {
				host: true,
				port: port,
			},
		};
	});
};
