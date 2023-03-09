import App from "./App.svelte";
import "./style.scss";

const app = new App({
  target: document.getElementById("app"),
});

export const version = "0.2.0";
export default app;
