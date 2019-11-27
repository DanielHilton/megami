import svelte from 'rollup-plugin-svelte';
import commonjs from 'rollup-plugin-commonjs';
import resolve from 'rollup-plugin-node-resolve';
import json from 'rollup-plugin-json';

export default [
	{
		input: ['src/index.js'],
		output: {
			dir: 'dist',
			format: 'cjs',
			sourcemap: true
		},
		plugins: [
			resolve(),
			svelte({
				include: 'src/components/**/*.svelte',
				css: css => {
					css.write('static/svelte.css')
				},
				nestedTransitions: true,
				skipIntroByDefault: true
			}),
			commonjs(),
			json()
		],
		experimentalCodeSplitting: true,
		experimentalDynamicImport: true,
		external: [
			'electron',
			'child_process',
			'fs',
			'path',
			'url',
			'module',
			'os'
		]
	}
];