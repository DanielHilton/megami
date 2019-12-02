import App from './components/App.svelte';
import svelte from 'svelte/compiler';

const app = new App({
    target: document.body,
    props: {
        version: svelte.version
    }
});

export default app;