import { defineConfig } from "vitest/config";

export default defineConfig({
	resolve: {
		alias: {
			"./Worker?worker": "src/homechart/worker/WorkerMock.ts",
			"@lib": "src/lib",
		},
	},
	test: {
		globalSetup: [
			"src/lib/testing/globalSetup.ts",
		],
		globals: true,
		setupFiles: [
			"src/lib/testing/dom.ts",
			"src/lib/testing/mocks.ts",
			"src/lib/testing/helpers.ts",
		],
		watch: false,
	},
});
